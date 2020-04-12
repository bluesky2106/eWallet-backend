package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// DataMsg : struct
type DataMsg struct {
	Key  Key         `json:"key"`
	Data interface{} `json:"data"`
}

// RabbitMQ : struct
type RabbitMQ struct {
	conf       *Config
	channel    *amqp.Channel
	connection *amqp.Connection

	Queues          []QueueName
	DeliveryChannel chan amqp.Delivery
	ErrorChannel    chan *amqp.Error
	CloseChannel    chan bool
}

// Init : rabbitmq config
//
// Be sure close connect and channel
func Init(conf *Config, queues []QueueName) *RabbitMQ {
	rbmq := new(RabbitMQ)
	rbmq.conf = conf
	rbmq.Queues = queues

	rbmq.Connect()

	return rbmq
}

// Connect to rabbitmq server
func (rbmq *RabbitMQ) Connect() {
	connURL := fmt.Sprintf("amqp://%s:%s@%s/", rbmq.conf.User, rbmq.conf.Password, rbmq.conf.URL)

	for {
		// log.Printf("Connecting to rabbitmq on %s\n", connURL)
		conn, err := amqp.Dial(connURL)
		if err == nil {
			rbmq.DeliveryChannel = make(chan amqp.Delivery)
			rbmq.ErrorChannel = make(chan *amqp.Error)
			rbmq.CloseChannel = make(chan bool)

			rbmq.connection = conn
			rbmq.connection.NotifyClose(rbmq.ErrorChannel)

			log.Println("Connection established!")

			if err = rbmq.openChannel(); err != nil {
				logError("Opening channel failed", err)
				rbmq.CloseAll()
				sleepDueToError(err)
				continue
			}

			if err = rbmq.declareQueue(); err != nil {
				logError("Queue declaration failed", err)
				rbmq.CloseAll()
				sleepDueToError(err)
				continue
			}

			return
		}

		sleepDueToError(err)
	}
}

func (rbmq *RabbitMQ) openChannel() error {
	channel, err := rbmq.connection.Channel()
	if err != nil {
		return err
	}
	rbmq.channel = channel
	return nil
}

func (rbmq *RabbitMQ) declareQueue() error {
	for _, queueName := range rbmq.Queues {
		_, err := rbmq.channel.QueueDeclare(
			string(queueName), // name
			true,              // durable
			false,             // delete when unused
			false,             // exclusive
			false,             // no-wait
			nil,               // arguments
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// PushMessage : queue name, data
func (rbmq *RabbitMQ) PushMessage(queue QueueName, data DataMsg) (err error) {
	bytes, _ := json.Marshal(data)
	qName := string(queue)

	err = rbmq.channel.Publish(
		"",    // exchange
		qName, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         bytes,
			DeliveryMode: amqp.Persistent,
		})
	return err
}

// ConsumeMessage : channel consume messages for each queue
func (rbmq *RabbitMQ) ConsumeMessage() error {
	for _, queueName := range rbmq.Queues {
		msgs, err := rbmq.channel.Consume(
			string(queueName), // queue
			"",                // messageConsumer
			false,             // auto-ack
			false,             // exclusive
			false,             // no-local
			false,             // no-wait
			nil,               // args
		)
		if err != nil {
			logError("Consuming message from queue "+string(queueName)+" failed", err)
			return err
		}
		go func() {
			for msg := range msgs {
				rbmq.DeliveryChannel <- msg
			}
		}()
	}

	return nil
}

// CloseAll : Close connection and channel
func (rbmq *RabbitMQ) CloseAll() error {
	if rbmq.channel != nil {
		err := rbmq.channel.Close()
		if err != nil {
			logError("Close channel error", err)
			return err
		}
		rbmq.channel = nil
	}

	if rbmq.connection != nil {
		err := rbmq.connection.Close()
		if err != nil {
			logError("Close connection error", err)
			return err
		}
		rbmq.connection = nil
	}

	rbmq.CloseChannel <- true

	return nil
}
