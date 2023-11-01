package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	secrets "github.com/tommzn/go-secrets"

	pubsub "github.com/tommzn/aws-pub-sub"
)

func main() {

	subscriber, err := bootstrap()
	if err != nil {
		panic(err)
	}
	lambda.Start(subscriber.Receive)
}

// bootstrap loads config and creates a new scheduled collector with a exchangerate datasource.
func bootstrap() (pubsub.LambdaHandler, error) {

	conf := loadConfig()
	secretsManager := newSecretsManager()
	logger := newLogger(conf, secretsManager)

	subscriber := pubsub.NewLambdaMessageSubscriber(logger)
	return subscriber, nil
}

// loadConfig from config file.
func loadConfig() config.Config {

	configSource, err := config.NewS3ConfigSourceFromEnv()
	if err != nil {
		panic(err)
	}

	conf, err := configSource.Load()
	if err != nil {
		panic(err)
	}
	return conf
}

// newSecretsManager retruns a new secrets manager from passed config.
func newSecretsManager() secrets.SecretsManager {
	return secrets.NewSecretsManager()
}

// newLogger creates a new logger from  passed config.
func newLogger(conf config.Config, secretsMenager secrets.SecretsManager) log.Logger {
	logger := log.NewLoggerFromConfig(conf, secretsMenager)
	logger = log.WithNameSpace(logger, "sns-subsriber")
	return log.WithK8sContext(logger)
}
