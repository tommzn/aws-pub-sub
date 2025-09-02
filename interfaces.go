package pubsub

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"google.golang.org/protobuf/proto"
)

// Publisher sends messages.
type Publisher interface {

	// Send will publish a message.
	Send(topicArn string, message proto.Message) error
}

// Consumer gets messages for processing.
type Consumer interface {

	// Process receives messages for further processing.
	Process(messageData string) error
}

// LambdaHandler invoked by messages forwarder from an AWS SNS topic.
type LambdaHandler interface {

	// Receive is an AWS lambda event handler to consume SNS messages.
	Receive(ctx context.Context, snsEvent events.SNSEvent)

	// Subscribe is used to register message consumer for specific topics.
	Subsribe(topicArn string, consumer Consumer)
}

// Persistence is used to store large messages.
type Persistence interface {

	// Upload a message to a persistence location. Returns an object id if passed message has been persisted successfully.
	Upload(topic string, message proto.Message) (*ObjectId, error)

	// Download can be used to retrieve a message by given object id.
	Download(id *ObjectId) (proto.Message, error)

	// Type returns type of used persistence layer.
	Type() string
}

// snsClient is an internal interface for an AWS SNS client.
type snsClient interface {

	// Publish sends a message to an SNS topic.
	Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}
