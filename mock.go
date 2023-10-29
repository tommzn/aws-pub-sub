package pubsub

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/sns"
	"google.golang.org/protobuf/proto"
)

// SnsMock for testing.
type snsMock struct {
	shouldReturnError bool
	messages          map[string][]*sns.PublishInput
}

// NewSnsMock returns a SNS client mock. You can specifiy whether each call to Send should return an error.
func newSnsMock(shouldReturnError bool) *snsMock {
	return &snsMock{
		shouldReturnError: shouldReturnError,
		messages:          make(map[string][]*sns.PublishInput),
	}
}

// Publish conforms to interface method and will persist given message locally for testing.
func (mock *snsMock) Publish(publishInput *sns.PublishInput) (*sns.PublishOutput, error) {

	if mock.shouldReturnError {
		return nil, errors.New("Failed!")
	}
	if _, ok := mock.messages[*publishInput.TopicArn]; !ok {
		mock.messages[*publishInput.TopicArn] = []*sns.PublishInput{}
	}
	mock.messages[*publishInput.TopicArn] = append(mock.messages[*publishInput.TopicArn], publishInput)
	return &sns.PublishOutput{}, nil
}

// PersistenceMock for testing.
type persistenceMock struct {
	shouldReturnError bool
	messages          map[ObjectId]proto.Message
}

// NewPersistenceMock returns a new mock. Use shouldReturnError to determins if each call to a method should return with an error.
func newPersistenceMock(shouldReturnError bool) *persistenceMock {

	return &persistenceMock{
		shouldReturnError: shouldReturnError,
		messages:          make(map[ObjectId]proto.Message),
	}
}

// Upload confirms to interface method and will store given message locally.
func (mock *persistenceMock) Upload(topic string, message proto.Message) (*ObjectId, error) {

	if mock.shouldReturnError {
		return nil, errors.New("Failed!")
	}

	id := ObjectId(newMessageId())
	mock.messages[id] = message
	return &id, nil
}

// Download will lookup if there's a message with given id in local storage.
func (mock *persistenceMock) Download(id *ObjectId) (proto.Message, error) {

	if mock.shouldReturnError {
		return nil, errors.New("Failed!")
	}

	if message, ok := mock.messages[*id]; ok {
		return message, nil
	}
	return nil, errors.New("Not found!")
}

// Type will return default value "MOCK".
func (mock *persistenceMock) Type() string {
	return "MOCK"
}
