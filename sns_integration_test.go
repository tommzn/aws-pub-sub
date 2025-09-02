package pubsub

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SnsIntegrationTestSuite struct {
	suite.Suite
	awsConfig aws.Config
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

	// Load AWS config with custom region
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(awsRegion))
	if err != nil {
		suite.T().Fatalf("Failed to load AWS config: %v", err)
	}
	suite.awsConfig = cfg

	snsTopic, isSet := os.LookupEnv("SNS_TEST_TOPIC")
	if !isSet {
		suite.T().Error("Missing SNS topic ARN for integration test, SNS_TEST_TOPIC")
	}
	suite.topicArn = snsTopic
}

func (suite *SnsIntegrationTestSuite) skipCI() {
	if _, isSet := os.LookupEnv("CI"); isSet {
		suite.T().SkipNow()
	}
}

func (suite *SnsIntegrationTestSuite) TestSendMessage() {
	publisher := &SnsPublisher{
		snsClient: sns.NewFromConfig(suite.awsConfig),
	}
	testMessage := &ExampleMessage{
		Value:     "AWS SNS Test Message",
		Timestamp: timestamppb.New(time.Now()),
	}

	suite.Nil(publisher.Send(suite.topicArn, testMessage))
}
