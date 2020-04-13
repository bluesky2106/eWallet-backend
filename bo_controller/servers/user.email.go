package servers

import (
	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/libs/rabbitmq"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
)

func (u *UserSrv) sendEmail(emailInfo pb.EmailInfo) error {
	data := rabbitmq.DataMsg{
		Key:  rabbitmq.KeySendEmail,
		Data: emailInfo,
	}

	err := u.rbmq.PushMessage(rabbitmq.QueueEmailService, data)
	if err != nil {
		return errs.New(errs.ECSystemError, err.Error(), "u.rbmq.PushMessage")
	}

	return nil
}
