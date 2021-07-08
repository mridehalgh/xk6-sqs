package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/mitchellh/mapstructure"
	"go.k6.io/k6/js/modules"
	"os"
)

func init() {
	modules.Register("k6/x/sqs", new(Sqs))
}

type Sqs struct{}

func (*Sqs) NewClient() *sqs.Client {
	cfg, err := getAwsConfig()

	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sqs.NewFromConfig(cfg)
	return client
}

func getAwsConfig() (aws.Config, error) {
	awsEndpoint := os.Getenv("AWS_ENDPOINT")

	if awsEndpoint != "" {
		customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: "eu-west-1",
			}, nil
		})

		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolver(customResolver))
		return cfg, err
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	return cfg, err
}

func (s *Sqs) Send(sqsClient *sqs.Client, messageInput interface{}) {
	var sqsMessageInput sqs.SendMessageInput
	_ = mapstructure.Decode(messageInput, &sqsMessageInput)
	_, err := sqsClient.SendMessage(context.TODO(), &sqsMessageInput)
	if err != nil {
		panic("unable to send  message, " + err.Error())
	}
}
