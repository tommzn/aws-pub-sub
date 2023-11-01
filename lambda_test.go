package pubsub

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LambdaTestSuite struct {
	suite.Suite
}

func TestLambdaTestSuite(t *testing.T) {
	suite.Run(t, new(LambdaTestSuite))
}

func (suite *LambdaTestSuite) TestLogdMessge() {

	topicArn := "arn:aws:sns:eu-west-0:1234567890:test"
	testMessage := &ExampleMessage{Value: "AWS SNS Test Message", Timestamp: timestamppb.New(time.Now())}
	snsEvent := snsEventForTest(topicArn, testMessage)

	consumer := newConsumerMock()
	lambdaSubscriber := lambdaMessageSubscriberForTest()
	lambdaSubscriber.Subsribe(topicArn, consumer)
	lambdaSubscriber.Receive(context.Background(), snsEvent)

	suite.Len(consumer.messages, 1)
}

func lambdaMessageSubscriberForTest() *LambdaMessageSubscriber {
	return NewLambdaMessageSubscriber(loggerForTest())
}

func snsEventForTest(TopicArn string, message proto.Message) events.SNSEvent {

	messageData, _ := marshalMessage(message)
	return events.SNSEvent{
		Records: []events.SNSEventRecord{
			events.SNSEventRecord{
				EventVersion:         "1",
				EventSubscriptionArn: "arn:aws:sns:eu-west-0:1234567890:xxx",
				EventSource:          "test",
				SNS: events.SNSEntity{
					TopicArn: TopicArn,
					Message:  messageData,
				},
			},
		},
	}
}
