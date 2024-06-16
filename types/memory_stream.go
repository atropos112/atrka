// Package types contains all the types that are used in the application.
package types

// InMemoryStreamChannel provides a channel for in-memory stream that is used by the InMemoryProducer and InMemoryConsumer.
var InMemoryStreamChannel = make(chan Envelope, 1000)
