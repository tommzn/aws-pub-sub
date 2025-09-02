package pubsub

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	confighlp "github.com/tommzn/go-config"
	"google.golang.org/protobuf/proto"
)

// NewSnsPublisher creates a new publisher with default settings.
func NewSnsPublisher(conf confighlp.Config) Publisher {
	awsCfg, _ := newAWSConfig(conf)
	return &SnsPublisher{
		snsClient: sns.NewFromConfig(awsCfg),
	}
}

// Send publishes passed message to given SNS topic.
// If the message exceeds the SNS message size limit it uses a persistence layer
// to store the message and publishes its id instead.
func (publisher *SnsPublisher) Send(topicArn string, message proto.Message) error {

	messageData, err := marshalMessage(message)
	if err != nil {
		return err
	}

	attributes := make(map[string]types.MessageAttributeValue)
	attributes[MESSAGE_TYPE_ATTRIBUTE] = types.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String(MESSAGE_TYPE_SNS),
	}

	if isSnsMessageSizeExceeded(messageData) {

		if publisher.persistence == nil {
			return errors.New("unable to publish large message without persistence layer")
		}

		objectId, err := publisher.persistence.Upload(topicArn, message)
		if err != nil {
			return err
		}

		messageData = string(*objectId)
		attributes[MESSAGE_TYPE_ATTRIBUTE] = types.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(publisher.persistence.Type()),
		}
	}

	_, err = publisher.snsClient.Publish(context.TODO(), &sns.PublishInput{
		Message:           aws.String(messageData),
		TopicArn:          aws.String(topicArn),
		MessageAttributes: attributes,
	})
	return err
}

// isSnsMessageSizeExceeded determines if AWS SNS message size (256KB) is reached.
func isSnsMessageSizeExceeded(messageData string) bool {
	return len([]byte(messageData)) > 256*1024
}
