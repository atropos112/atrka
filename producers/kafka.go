// Package producers is a package that contains all the producers that can be used in the application.
package producers

import (
	"context"
	"errors"
	"fmt"

	"github.com/atropos112/atrka/types"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
)

// KafkaStreamProducer is a stream producer that sends messages to a Kafka stream
type KafkaStreamProducer struct {
	Topic      string
	Brokers    []string
	Registries []string
}

// Init initializes the Kafka stream producer
func (k *KafkaStreamProducer) Init(ctx *context.Context) error {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(k.Brokers...),
		kgo.DefaultProduceTopic(k.Topic),
	)
	if err != nil {
		return err
	}

	// Using the struct itself as the key
	*ctx = context.WithValue(*ctx, k, cl)
	return nil
}

// Send sends a message to the Kafka stream
func (k *KafkaStreamProducer) Send(ctx context.Context, message types.Message) error {
	cl, ok := ctx.Value(k).(*kgo.Client)
	if !ok {
		return errors.New("invalid context")
	}

	var serde sr.Serde
	var prErr error

	cl.Produce(
		context.Background(),
		&kgo.Record{
			Value: serde.MustEncode(message),
		},
		func(r *kgo.Record, err error) {
			prErr = err
			fmt.Printf("Produced simple record, value bytes: %x\n", r.Value)
		},
	)
	if prErr != nil {
		return prErr
	}

	return nil
}
