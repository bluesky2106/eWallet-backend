package servers

import (
	"encoding/json"
	"reflect"

	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/libs/rabbitmq"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// RouteMessage : amqp.Delivery
func (e *EmailSrv) RouteMessage(m *amqp.Delivery) {
	var (
		body *rabbitmq.DataMsg
		data pb.EmailInfo
	)

	err := json.Unmarshal(m.Body, &body)
	if err != nil {
		zap.L().Error("body data invalid", zap.Error(err))
		m.Ack(false)
		return
	}

	zap.L().Info("reiceived message",
		zap.String("exchange", m.Exchange),
		zap.String("routingKey", m.RoutingKey),
		zap.Any("key", body.Key),
		zap.Any("data", body.Data),
	)

	switch body.Key {
	case rabbitmq.KeySendEmail:
		data = reflect.ValueOf(body.Data).Interface().(pb.EmailInfo)
		e.SendEmail(&data)
		m.Ack(false)
	}
}

// SendEmail : email request
func (e *EmailSrv) SendEmail(data *pb.EmailInfo) error {
	if !e.isValidEmailInfo(data) {
		return errs.New(errs.ECInvalidMessage)
	}

	return e.sendgrid.SendEmail(data)
}
