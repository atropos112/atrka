// Package producers is a package that contains all the producers that can be used in the application.
package producers

import (
	"fmt"

	"github.com/atropos112/atrka/types"
)

// InMemoryProducer is a producer that sends messages to an in-memory stream
type InMemoryProducer struct{}

// Send is a method that sends a message to the in-memory stream
func (imp InMemoryProducer) Send(envelope types.Envelope) error {
	fmt.Println("Sending message: ", envelope)
	types.InMemoryStreamChannel <- envelope

	return nil
}

// Init initializes the producer a no-op
func (imp InMemoryProducer) Init() error {
	return nil
}
