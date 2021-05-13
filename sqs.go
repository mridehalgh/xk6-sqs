package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"go.k6.io/k6/js/modules"
	"github.com/mitchellh/mapstructure"
)

func init() {
	modules.Register("k6/x/sqs", new(Sqs))
}

type Sqs struct{}

func (*Sqs) NewClient() *sqs.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sqs.NewFromConfig(cfg)
	return client
}

func (s *Sqs) Send(sqsClient *sqs.Client, messageInput interface{}) {
	var sqsMessageInput sqs.SendMessageInput
	_ = mapstructure.Decode(messageInput, &sqsMessageInput)
	_, err := sqsClient.SendMessage(context.TODO(), &sqsMessageInput)
	if err != nil {
		panic("unable to send  message, " + err.Error())
	}
}
