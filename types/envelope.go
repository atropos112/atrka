package types

// Envelope  is an interface describing a package sent over stream.
type Envelope struct {
	Plugin          string
	Timestamp       int64
	EnvelopeVersion string
	MessageVersion  string
	Message         Message
}

// MakeDummyEnvelope is a dummy envelope for testing.
func MakeDummyEnvelope(msg Message) Envelope {
	return Envelope{Plugin: "", Timestamp: 0, EnvelopeVersion: "", MessageVersion: "", Message: msg}
}

// RegisterSchema(logger *zap.SugaredLogger, topic string, registries []string, message Message) error
// Message() Message
// MessageToEnvelope(message Message) Envelope
