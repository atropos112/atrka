package types

// Producer is an interface for a streamer that produces messages.
type Producer interface {
	Send(envelope Envelope) error
	Init() error
}
