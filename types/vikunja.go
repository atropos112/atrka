// Package types is a collection of types used in the application both for producer and consumer
package types

//go:generate stringer -type=VikunjaActionType

// VikunjaActionType is an enum to represent the type of action that happened in Vikunja
type VikunjaActionType int

// Define constants for the enum
const (
	TaskCreated VikunjaActionType = iota + 1 // Start from 1 instead of 0
)

// VikunjaActionMessage is message describing an action that happened in Vikunja
type VikunjaActionMessage struct {
	Type   VikunjaActionType
	Object string
}
