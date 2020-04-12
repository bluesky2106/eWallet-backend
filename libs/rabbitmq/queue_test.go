package rabbitmq

import (
	"log"
	"testing"
	"time"

	gwConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/stretchr/testify/assert"
	_ "go.uber.org/automaxprocs"
)

var (
	gwConf       *gwConfig.Config
	rabbitmqConf *Config
	queues       = []QueueName{QueueEmailService}
	rabbitMQ     *RabbitMQ
)

func init() {
	gwConf = &gwConfig.Config{
		RabbitMQ: gwConfig.RabbitMQ{
			Host:     "localhost",
			Port:     "5672",
			Username: "rabbitmq",
			Password: "rabbitmq",
		},
		Env: gwConfig.Debug,
	}
}

func TestParseConfig(t *testing.T) {
	assert := assert.New(t)

	// load config
	rabbitmqConf = ParseConfig(gwConf)
	assert.NotNil(rabbitmqConf)
	assert.Equal("rabbitmq", rabbitmqConf.User, "rabbitmq user mismatched")
	assert.Equal("rabbitmq", rabbitmqConf.Password, "rabbitmq pwd mismatched")
	assert.Equal("localhost:5672", rabbitmqConf.URL, "rabbitmq url mismatched")
}

func TestInitRabbitMQ(t *testing.T) {
	assert := assert.New(t)

	rabbitMQ = Init(rabbitmqConf, queues)
	// init rabbitmq queues
	assert.NotNil(rabbitMQ)
}

func TestPushMessage(t *testing.T) {
	assert := assert.New(t)

	err := rabbitMQ.PushMessage(QueueEmailService, DataMsg{Key: "key1", Data: "{id:1}"})
	assert.Nil(err)

	err = rabbitMQ.PushMessage(QueueEmailService, DataMsg{Key: "key2", Data: "{id:2}"})
	assert.Nil(err)
}

func TestConsumeMessages(t *testing.T) {
	assert := assert.New(t)
	err := rabbitMQ.ConsumeMessage()
	assert.Nil(err)
	go func() {
		for {
			select {
			case msg := <-rabbitMQ.DeliveryChannel:
				// log.Println(msg)
				log.Println(string(msg.Body))
				assert.NotNil(msg)
				msg.Ack(false)
			case <-rabbitMQ.CloseChannel:
				return
			}
		}
	}()

	time.Sleep(time.Duration(2) * time.Second)
	err = rabbitMQ.CloseAll()
	assert.Nil(err)
}
