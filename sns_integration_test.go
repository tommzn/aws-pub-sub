package pubsub

import (
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SnsIntegrationTestSuite struct {
	suite.Suite
	awsConfig *aws.Config
	topicArn  string
}

func TestSnsIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(SnsIntegrationTestSuite))
}

func (suite *SnsIntegrationTestSuite) SetupSuite() {

	suite.skipCI()

	awsRegion, isSet := os.LookupEnv("SNS_TEST_REGION")
	if !isSet {
		suite.T().Error("Missing AWS region for integration test, SNS_TEST_REGION")
	}
	awsConfig := aws.NewConfig()
	awsConfig.WithRegion(awsRegion)
	suite.awsConfig = awsConfig

	snsTopic, isSet := os.LookupEnv("SNS_TEST_TOPIC")
	if !isSet {
		suite.T().Error("Missing AWS region for integration test, SNS_TEST_TOPIC")
	}
	suite.topicArn = snsTopic
}

func (suite *SnsIntegrationTestSuite) skipCI() {
	if _, isSet := os.LookupEnv("CI"); isSet {
		suite.T().SkipNow()
	}
}

func (suite *SnsIntegrationTestSuite) TestSendMessge() {

	publisher := &SnsPublisher{
		snsClient: sns.New(newAwsSession(suite.awsConfig)),
	}
	testMessage := &ExampleMessage{Value: "AWS SNS Test Message", Timestamp: timestamppb.New(time.Now())}

	suite.Nil(publisher.Send(suite.topicArn, testMessage))
}
