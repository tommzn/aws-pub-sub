package pubsub

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	config "github.com/tommzn/go-config"
)

// newAWSConfig try to find AWS region in passed config or in environment variable AWS_REGION
// and returns a new AWS config.
func newAWSConfig(conf config.Config) *aws.Config {

	awsConfig := aws.NewConfig()

	if conf != nil {
		configKeys := []string{"aws.region", "aws.region"}
		for _, configKey := range configKeys {
			if awsRegion := conf.Get(configKey, nil); awsRegion != nil {
				return awsConfig.WithRegion(*awsRegion)
			}
		}

	}

	if awsRegion, ok := os.LookupEnv("AWS_REGION"); ok {
		return awsConfig.WithRegion(awsRegion)
	}

	return awsConfig
}

// NewAwsSession create a new AWS session with given config.
func newAwsSession(conf *aws.Config) *session.Session {
	return session.Must(session.NewSession(conf))
}
