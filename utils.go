package pubsub

import (
	"encoding/base64"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// NewMessageId generates a new UUID V4 message id.
func newMessageId() string {
	return uuid.NewV4().String()
}

// MarshalMessage uses protobuf to marshal given message.
func marshalMessage(message proto.Message) (string, error) {

	protoContent, err := proto.Marshal(message)
	return base64.StdEncoding.EncodeToString(protoContent), err
}

// UnmarshalMessage run proto unnarshaller to convert given message data into passed message.
func UnmarshalMessage(messageData string, message proto.Message) error {

	stringData := strings.TrimSuffix(strings.TrimPrefix(messageData, "\""), "\"")
	protoData, err := base64.StdEncoding.DecodeString(stringData)
	if err != nil {
		return err
	}
	return proto.Unmarshal(protoData, message)
}
