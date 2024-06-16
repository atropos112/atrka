package types

import (
	"context"
	"encoding/json"

	"github.com/hamba/avro/v2"
	"github.com/twmb/franz-go/pkg/sr"
	"go.uber.org/zap"
)

// AvroEnvelope represents an Avro schema envelope used to describe the schema of a message.
type AvroEnvelope struct {
	Type      string
	Name      string
	Namespace string
	Fields    []AvroField
}

// AvroField represents a field in an Avro schema.
type AvroField struct {
	Name string
	Type string
}

// RegisterSchema registers the schema with the schema registry (such as Kafka Schema Registry).
func (a *AvroEnvelope) RegisterSchema(
	logger *zap.SugaredLogger,
	topic string,
	registries []string,
	message Message,
) error {
	// 1. Create a JSON schema
	byteJSONSchema, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	schemaText := string(byteJSONSchema)

	// 2. Parse the JSON schema to an Avro schema
	avroSchema, err := avro.Parse(schemaText)
	if err != nil {
		return err
	}

	// 3. Register the Avro schema with the schema registry

	rcl, err := sr.NewClient(sr.URLs(registries...))
	if err != nil {
		return err
	}

	ss, err := rcl.CreateSchema(context.Background(), topic+"-value", sr.Schema{
		Schema: schemaText,
		Type:   sr.TypeAvro,
	})
	if err != nil {
		return err
	}

	logger.Infow(
		"created schema",
		"schema", schemaText,
		"topic", topic+"-value",
		"subject", ss.Subject,
		"version", ss.Version,
		"id", ss.ID,
		"type", "avro", // Can also do json or protobuf
	)

	var serde sr.Serde
	serde.Register(
		ss.ID,
		message,
		sr.EncodeFn(func(v any) ([]byte, error) {
			return avro.Marshal(avroSchema, v)
		}),
		sr.DecodeFn(func(b []byte, v any) error {
			return avro.Unmarshal(avroSchema, b, v)
		}),
	)

	return nil
}
