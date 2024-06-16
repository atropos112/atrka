package types

// Consumer  is an interface for a streamer that consumes messages.
type Consumer interface {
	Consume(envelope Envelope) error
	Init() error
}
