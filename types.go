package pubsub

import (
	"github.com/tommzn/go-log"
	"google.golang.org/protobuf/proto"
)

// MESSAGE_TYPE_ATTRIBUTE is the message attribute name where message type is placed.
const MESSAGE_TYPE_ATTRIBUTE string = "MESSAGE_TYPE"

const (

	// MESSAGE_TYPE_SNS is used if a message is send via AWS SNS directly.
	MESSAGE_TYPE_SNS string = "SNS"

	// MESSAGE_TYPE_S3 is used if a message is to large for SNS and therefore be uploaded to ans S4 bucket.
	MESSAGE_TYPE_S3 string = "S3"
)

// ObjectId is used to identify messages passed to a persistence layer.
type ObjectId string

// SnsPublisher is a client to send messages to AWS SNS. In case a message exceeds SNS message size imit a persitence layer
// can be used to store such a large message and distribute its id via SNS, only.
type SnsPublisher struct {
	snsClient   snsClient
	persistence Persistence
}

// LogConsumer simple log received message.
type LogConsumer struct {
	logger  log.Logger
	message proto.Message
}

// LambdaMessageSubscriber handles message received from AWS sNS
// and forwards this messages, depending on topic to registered consumers.
type LambdaMessageSubscriber struct {
	logger log.Logger

	// Consumers, list of message consumer per topic
	consumers map[string][]Consumer
}
