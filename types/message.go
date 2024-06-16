package types

// Message is a generic interface for messages.
type Message interface{}

// CronMessage is just a timestamp (nothing else to put there)
type CronMessage struct {
	Timestamp int64
}

// JSONMessage is a message that contains a json object (as a map[string]interface{}), not ideal, should be avoided if possible.
type JSONMessage struct {
	JSON map[string]interface{}
}
