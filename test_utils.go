package pubsub

import (
	"github.com/tommzn/go-log"
	"google.golang.org/protobuf/proto"
)

type protoMessage proto.Message

func loggerForTest() log.Logger {
	return log.NewLogger(log.Debug, nil, nil)
}
