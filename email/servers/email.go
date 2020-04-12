package servers

import (
	"github.com/bluesky2106/eWallet-backend/email/config"
	"github.com/bluesky2106/eWallet-backend/libs/rabbitmq"
	"github.com/bluesky2106/eWallet-backend/libs/sendgrid"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
)

// EmailSrv : email server
type EmailSrv struct {
	conf     *config.Config
	sendgrid *sendgrid.MailManager
	rbmq     *rabbitmq.RabbitMQ

	pb.UnimplementedEmailSvcServer
}

// NewEmailSrv : new email server
func NewEmailSrv(conf *config.Config, sendgrid *sendgrid.MailManager, rbmq *rabbitmq.RabbitMQ) *EmailSrv {
	return &EmailSrv{
		conf:     conf,
		sendgrid: sendgrid,
		rbmq:     rbmq,
	}
}
