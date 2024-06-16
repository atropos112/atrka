// Package producers is a package that contains all the producers that can be used in the application.
package producers

import (
	"github.com/atropos112/atrka/types"
	"go.uber.org/zap"
)

// LoggingProducer is a producer that logs the message (most likely for testing purposes only)
type LoggingProducer struct {
	Logger *zap.Logger
}

// Send is a method that sends a message to the logger
func (lp LoggingProducer) Send(envelope types.Envelope) error {
	lp.Logger.Sugar().Infow("", "Message", envelope.Message)
	return nil
}

// Init initializes the producer (no-op in this case)
func (lp LoggingProducer) Init() error {
	return nil
}
