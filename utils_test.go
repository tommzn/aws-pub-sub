package pubsub

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UtilsTestSuite struct {
	suite.Suite
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}

func (suite *UtilsTestSuite) TestGenrateId() {

	uuidattern := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	suite.True(uuidattern.MatchString(newMessageId()))
}

func (suite *UtilsTestSuite) TestMessageMarshaling() {

	testMessage := &ExampleMessage{Value: "Val01", Timestamp: timestamppb.New(time.Now())}

	messageDate, err := marshalMessage(testMessage)
	suite.Nil(err)

	testMessage2 := &ExampleMessage{}
	suite.Nil(UnmarshalMessage(messageDate, testMessage2))

	suite.Equal(testMessage.Value, testMessage2.Value)
	suite.Equal(testMessage.Timestamp.AsTime().Format(time.RFC3339), testMessage2.Timestamp.AsTime().Format(time.RFC3339))
}
