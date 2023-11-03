package pubsub

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SnsTestSuite struct {
	suite.Suite
}

func TestSnsTestSuite(t *testing.T) {
	suite.Run(t, new(SnsTestSuite))
}

func (suite *SnsTestSuite) TestSendMessge() {

	publisher := publisherForTest(false)
	testMessage := messageForTest(100)
	topic := "test"

	suite.Nil(publisher.Send(topic, testMessage))
	suite.Len(publisher.(*SnsPublisher).snsClient.(*snsMock).messages, 1)

	testMessage2 := messageForTest(300000)
	suite.NotNil(publisher.Send(topic, testMessage2))
	suite.Len(publisher.(*SnsPublisher).snsClient.(*snsMock).messages, 1)
}

func (suite *SnsTestSuite) TestSendKargeMessge() {

	publisher := publisherForTest(false)
	publisher.(*SnsPublisher).persistence = newPersistenceMock(false)
	testMessage := messageForTest(300000)
	topic := "test"

	suite.Nil(publisher.Send(topic, testMessage))
	suite.Len(publisher.(*SnsPublisher).snsClient.(*snsMock).messages, 1)
	suite.Len(publisher.(*SnsPublisher).persistence.(*persistenceMock).messages, 1)
}

func (suite *SnsTestSuite) TestSendMessgeFailed() {

	publisher := publisherForTest(true)
	testMessage := messageForTest(100)
	topic := "test"

	suite.NotNil(publisher.Send(topic, testMessage))
	suite.Len(publisher.(*SnsPublisher).snsClient.(*snsMock).messages, 0)
}

func (suite *SnsTestSuite) TestSendKargeMessgeFailed() {

	publisher := publisherForTest(false)
	publisher.(*SnsPublisher).persistence = newPersistenceMock(true)
	testMessage := messageForTest(300000)
	topic := "test"

	suite.NotNil(publisher.Send(topic, testMessage))
	suite.Len(publisher.(*SnsPublisher).snsClient.(*snsMock).messages, 0)
	suite.Len(publisher.(*SnsPublisher).persistence.(*persistenceMock).messages, 0)
}

func publisherForTest(shouldReturnError bool) Publisher {

	publisher := NewSnsPublisher(nil)
	publisher.(*SnsPublisher).snsClient = newSnsMock(shouldReturnError)
	return publisher
}

func messageForTest(size int) proto.Message {
	return &ExampleMessage{Value: strings.Repeat("x", size), Timestamp: timestamppb.New(time.Now())}
}
