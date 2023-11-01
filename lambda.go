package pubsub

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tommzn/go-log"
)

func NewLambdaMessageSubscriber(logger log.Logger) *LambdaMessageSubscriber {
	return &LambdaMessageSubscriber{
		logger:    logger,
		consumers: make(map[string][]Consumer),
	}
}

// Receive is a AWS lambda event handler to consumer SNS messages.
func (subsriber *LambdaMessageSubscriber) Receive(ctx context.Context, snsEvent events.SNSEvent) {

	defer subsriber.logger.Flush()
	subsriber.logger.Debugf("Receive %d Event(s).", len(snsEvent.Records))

	for _, record := range snsEvent.Records {

		topicArn := record.SNS.TopicArn
		subsriber.logger.Debug("Process message from topic: ", topicArn)
		if consumers, ok := subsriber.consumers[topicArn]; ok {

			for _, consumer := range consumers {
				if err := consumer.Process(record.SNS.Message); err != nil {
					subsriber.logger.Info("Message processing failed, reason: ", err)
				}

			}
		} else {
			subsriber.logger.Info("No consumer for topic: ", topicArn)
		}
	}
}

// Subsribe is used to register message consumer for specific topics.
func (subsriber *LambdaMessageSubscriber) Subsribe(topicArn string, consumer Consumer) {

	if _, ok := subsriber.consumers[topicArn]; !ok {
		subsriber.consumers[topicArn] = []Consumer{}
	}
	subsriber.consumers[topicArn] = append(subsriber.consumers[topicArn], consumer)
}
