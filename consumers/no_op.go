// Package consumers contains all the consumers that can be used in the application.
package consumers

import "github.com/atropos112/atrka/types"

// NoOpConsumer is a consumer that no-ops the consumption of messages.
type NoOpConsumer struct{}

// Consume no-ops the consumer.
func (noc NoOpConsumer) Consume(_ types.Envelope) error {
	return nil
}

// Init no-ops the consumer.
func (noc NoOpConsumer) Init() error {
	return nil
}
