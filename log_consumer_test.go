package pubsub

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LogConsumerTestSuite struct {
	suite.Suite
}

func TestLogConsumerTestSuite(t *testing.T) {
	suite.Run(t, new(LogConsumerTestSuite))
}

func (suite *LogConsumerTestSuite) TestLogdMessge() {

	consumer := NewLogConsumer(loggerForTest())
	consumer.message = &ExampleMessage{}
	testMessage := &ExampleMessage{Value: "AWS SNS Test Message", Timestamp: timestamppb.New(time.Now())}
	messageData, err := marshalMessage(testMessage)
	suite.Nil(err)

	suite.Nil(consumer.Process(messageData))
}
