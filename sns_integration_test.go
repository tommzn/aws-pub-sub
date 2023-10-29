package pubsub

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/suite"
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

}
