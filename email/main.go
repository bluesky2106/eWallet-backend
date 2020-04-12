package main

import (
	"fmt"
	"net"

	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/email/config"
	"github.com/bluesky2106/eWallet-backend/email/servers"
	"github.com/bluesky2106/eWallet-backend/libs/rabbitmq"
	"github.com/bluesky2106/eWallet-backend/libs/sendgrid"
	"github.com/bluesky2106/eWallet-backend/log"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	logger *zap.Logger
)

func main() {
	// 1. Get global config
	conf := commonConfig.ParseConfig("config.json", "../config")

	// 2. Init logger
	logger = log.InitLogger(conf.Env)

	// 3. Extract email config
	emailConf := config.ParseConfig(conf)
	emailConf.Print()

	// 4. Init rbmq, sendgrid and email server
	var (
		sendgrid = sendgrid.NewMailManager(emailConf.Sendgrid)
		rbmq     = rabbitmq.Init(emailConf.RabbitMQ, []rabbitmq.QueueName{rabbitmq.QueueEmailService})
		emailSrv = servers.NewEmailSrv(emailConf, sendgrid, rbmq)
	)
	defer func() {
		if rbmq != nil {
			err := rbmq.CloseAll()
			if err != nil {
				logger.Error("rabbitmq close", zap.Error(err))
			}
		}
	}()
	consumingQueue(rbmq, emailSrv)

	// 5. Run grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", emailConf.Host, emailConf.Port))
	if err != nil {
		logger.Error("failed to listen", zap.Error(err))
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEmailSvcServer(grpcServer, emailSrv)
	grpcServer.Serve(lis)
}

func consumingQueue(rbmq *rabbitmq.RabbitMQ, emailSrv *servers.EmailSrv) {
	err := rbmq.ConsumeMessage()
	if err != nil {
		logger.Error("consume message error: ", zap.Error(err))
	}

	go func() {
		for {
			select {
			case msg := <-rbmq.DeliveryChannel:
				emailSrv.RouteMessage(&msg)
			case err := <-rbmq.ErrorChannel:
				logger.Error("Reconnecting after connection error: ", zap.Error(err))
				rbmq.Connect()
				rbmq.ConsumeMessage()
			case <-rbmq.CloseChannel:
				return
			}
		}
	}()
}
