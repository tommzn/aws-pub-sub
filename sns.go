package pubsub

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/sns"
	"google.golang.org/protobuf/proto"
)

// NewSnsPublisher creates a new publisher with default settings.
func NewSnsPublisher() Publisher {
	return &SnsPublisher{
		snsClient: sns.New(newAwsSession(newAWSConfig(nil))),
	}
}

// Send publishes passed message to given SNS topic.
// Is case given message exceeds SNS message size limit it used a perstience layer to store this message
// and will publish its id, only.
func (publisher *SnsPublisher) Send(topicArn string, message proto.Message) error {

	messageData, err := marshalMessage(message)
	if err != nil {
		return err
	}

	attributes := make(map[string]*sns.MessageAttributeValue)
	attributes[MESSAGE_TYPE_ATTRIBUTE] = &sns.MessageAttributeValue{DataType: awsString("String"), StringValue: awsString(MESSAGE_TYPE_SNS)}
	if isSnsMessageSizeExceeded(messageData) {

		if publisher.persistence == nil {
			return errors.New("Unable to publish large message without persistence layer.")
		}

		objectId, err := publisher.persistence.Upload(topicArn, message)
		if err != nil {
			return err
		}
		messageData = string(*objectId)
		attributes[MESSAGE_TYPE_ATTRIBUTE] = &sns.MessageAttributeValue{DataType: awsString("String"), StringValue: awsString(publisher.persistence.Type())}
	}

	_, err1 := publisher.snsClient.Publish(&sns.PublishInput{
		Message:           awsString(messageData),
		TopicArn:          awsString(topicArn),
		MessageAttributes: attributes,
	})
	return err1
}

// IsSnsMessageSizeExceeded determines if AWS SNS message size (256KB) is reached.
func isSnsMessageSizeExceeded(messageData string) bool {
	return len([]byte(messageData)) > 256*1024
}
