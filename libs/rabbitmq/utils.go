package rabbitmq

import (
	"log"
	"time"
)

// QueueName : queue name
type QueueName string

// Rabbitmq: define all of queues name
const (
	// QueueEmailService : email queue
	QueueEmailService QueueName = "QueueEmailService"
)

// Key identifies actions
type Key string

const (
	// KeySendEmail : send email
	KeySendEmail Key = "KeySendEmail"
)

func logError(message string, err error) {
	if err != nil {
		log.Printf("%s: %s\n", message, err)
	}
}

func sleepDueToError(err error) {
	logError("Connection to rabbitmq failed. Retrying in 10 seconds... ", err)
	time.Sleep(10 * time.Second)
}
