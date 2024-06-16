// Package consumers contains all the consumers that can be used in the application.
package consumers

import (
	"fmt"
	"time"

	"github.com/atropos112/atrka/types"
	"golang.org/x/sync/errgroup"
)

// InMemoryConsumer is a consumer that consumes messages from an in-memory stream.
type InMemoryConsumer struct{}

// Consume is consuming a message, reacting to it somehow.
func (imc InMemoryConsumer) Consume(envelope types.Envelope) error {
	fmt.Println("Consuming message: ", envelope)
	time.Sleep(3 * time.Second)
	return nil
}

// Init initializes the consumer, starting the consumption of messages.
func (imc InMemoryConsumer) Init() error {
	eg := errgroup.Group{}
	for envelope := range types.InMemoryStreamChannel {
		fn := func() error {
			return imc.Consume(envelope)
		}

		eg.Go(fn)
	}

	return eg.Wait()
}
