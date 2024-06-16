package types

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// Trigger is a generic struct to express the concept of a trigger
type Trigger[T Producer] struct {
	Handler  func(ctx context.Context, ch chan<- Message) error // Init is a function that initializes the trigger this might be connecting to a Kafka topic, starting a timer, etc. and could be long running
	Producer T
}

// NewTrigger creates a new trigger
func NewTrigger[T Producer](handler func(ctx context.Context, ch chan<- Message) error, producer T) Trigger[T] {
	return Trigger[T]{Handler: handler, Producer: producer}
}

// Start starts the trigger
func (t *Trigger[T]) Start(ctx context.Context) error {
	if err := t.Producer.Init(); err != nil {
		return err
	}

	// Create a message channel, note without buffer size, this will block the producer
	messageChan := make(chan Message, 1000)

	eg := errgroup.Group{}

	// Creating a coroutine to run the handler and the message channel processing
	eg.Go(func() error { return t.Handler(ctx, messageChan) })
	eg.Go(func() error { return t.ProcessMessageChannel(messageChan) })

	return eg.Wait()
}

// ProcessMessageChannel processes the message channel
func (t *Trigger[T]) ProcessMessageChannel(ch chan Message) error {
	for msg := range ch {
		envelope := MakeDummyEnvelope(msg)
		if err := t.Producer.Send(envelope); err != nil {
			return err
		}
	}

	return nil
}
