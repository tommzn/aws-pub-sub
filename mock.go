package pubsub

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"google.golang.org/protobuf/proto"
)

// snsMock for testing.
type snsMock struct {
	shouldReturnError bool
	messages          map[string][]*sns.PublishInput
}

// newSnsMock returns an SNS client mock. You can specify whether each call to Publish should return an error.
func newSnsMock(shouldReturnError bool) *snsMock {
	return &snsMock{
		shouldReturnError: shouldReturnError,
		messages:          make(map[string][]*sns.PublishInput),
	}
}

// Publish conforms to snsClient interface method and will persist given message locally for testing.
func (mock *snsMock) Publish(ctx context.Context, publishInput *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {

	if mock.shouldReturnError {
		return nil, errors.New("failed")
	}
	if _, ok := mock.messages[*publishInput.TopicArn]; !ok {
		mock.messages[*publishInput.TopicArn] = []*sns.PublishInput{}
	}
	mock.messages[*publishInput.TopicArn] = append(mock.messages[*publishInput.TopicArn], publishInput)
	return &sns.PublishOutput{}, nil
}

// persistenceMock for testing.
type persistenceMock struct {
	shouldReturnError bool
	messages          map[ObjectId]proto.Message
}

// newPersistenceMock returns a new mock. Use shouldReturnError to determine if each call should return an error.
func newPersistenceMock(shouldReturnError bool) *persistenceMock {
	return &persistenceMock{
		shouldReturnError: shouldReturnError,
		messages:          make(map[ObjectId]proto.Message),
	}
}

// Upload conforms to Persistence interface and will store given message locally.
func (mock *persistenceMock) Upload(topic string, message proto.Message) (*ObjectId, error) {

	if mock.shouldReturnError {
		return nil, errors.New("failed")
	}

	id := ObjectId(newMessageId())
	mock.messages[id] = message
	return &id, nil
}

// Download will lookup if there's a message with given id in local storage.
func (mock *persistenceMock) Download(id *ObjectId) (proto.Message, error) {

	if mock.shouldReturnError {
		return nil, errors.New("failed")
	}

	if message, ok := mock.messages[*id]; ok {
		return message, nil
	}
	return nil, errors.New("not found")
}

// Type will return default value "MOCK".
func (mock *persistenceMock) Type() string {
	return "MOCK"
}

// consumerMock is used for test message subscription.
type consumerMock struct {
	messages []string
}

// newConsumerMock creates a new consumer mock for testing.
func newConsumerMock() *consumerMock {
	return &consumerMock{
		messages: []string{},
	}
}

// Process will add passed message data to internal list.
func (mock *consumerMock) Process(messageData string) error {
	mock.messages = append(mock.messages, messageData)
	return nil
}
