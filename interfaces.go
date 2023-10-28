package pubsub

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"google.golang.org/protobuf/proto"
)

// Publisher sends messages.
type Publisher interface {

	// Send will publish a message.
	Send(topic string, message proto.Message) error
}

// Consumer gets messages for processing.
type Consumer interface {

	// Process receives messages for furthe processing.
	Process(message proto.Message) error
}

// AwsLambdaHandler invoked by messges forwarder from a AWS SNS topic.
type AwsLambdaHandler interface {

	// Receive is a AWS lambda event handler to consumer SNS messages.
	Receive(ctx context.Context, snsEvent events.SNSEvent)

	// Subsribe is used to register message consumer for specific topics.
	Subsribe(topic string, consumer Consumer)
}

// Persistence is used to store large message.
type Persistence interface {

	// Upload a message to a persitence location. Returns an object id if passed message has been persisted successful.
	Upload(topic string, message proto.Message) (*ObjectId, error)

	// Download can be used to retrieve a message by given object id.
	Download(id *ObjectId) (proto.Message, error)
}
