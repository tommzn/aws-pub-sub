package pubsub

import (
	"github.com/tommzn/go-log"
)

// NewLogConsumer returns a new log consumer which logs received messaged to passed logger.
func NewLogConsumer(logger log.Logger) *LogConsumer {
	return &LogConsumer{
		logger: logger,
	}
}

// Process log given message to used logger.
func (consumer *LogConsumer) Process(messageData string) error {

	defer consumer.logger.Flush()

	consumer.logger.Status("Message Received.")
	if consumer.message != nil {

		err := UnmarshalMessage(messageData, consumer.message)
		if err != nil {
			return err
		}
		consumer.logger.Statusf("Message %T: %s", consumer.message, consumer.message)

	}
	return nil
}
