package pubsub

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	confighlp "github.com/tommzn/go-config"
)

// newAWSConfig tries to find AWS region in passed config or in environment variable AWS_REGION
// and returns a new AWS config.
func newAWSConfig(conf confighlp.Config) (aws.Config, error) {

	if conf != nil {
		configKeys := []string{"aws.region", "aws.s3.region"}
		for _, configKey := range configKeys {
			if awsRegion := conf.Get(configKey, nil); awsRegion != nil {
				return config.LoadDefaultConfig(context.TODO(),
					config.WithRegion(*awsRegion),
				)
			}
		}
	}

	if awsRegion, ok := os.LookupEnv("AWS_REGION"); ok {
		return config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(awsRegion),
		)
	}

	// fallback: default AWS resolution chain (shared config, env vars, IAM role, etc.)
	return config.LoadDefaultConfig(context.TODO())
}

// newAWSConfigWithRegion explicitly loads AWS config with given region (helper).
func newAWSConfigWithRegion(region string) (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
}
