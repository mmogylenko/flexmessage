package flexmessage

import (
	"strings"
)

// FlexMessage type used for Msg/Err representation
type FlexMessage struct {
	Messages      []string `json:"messages,omitempty"`
	Errors        []string `json:"errors,omitempty"`
	ErrorExitCode int      `json:"-"`
}

// NewFlexMessage returns a new FlexMessage with a default values
func NewFlexMessage() *FlexMessage {
	return &FlexMessage{
		Messages:      nil,
		Errors:        nil,
		ErrorExitCode: 1,
	}
}

// Empty func checks if there are any notifications
// both messages and errors
func (f *FlexMessage) Empty() bool {
	if len(f.Errors) > 0 || len(f.Messages) > 0 {
		return false
	}
	return true
}

// NoErrors func checks if there are error entries
func (f *FlexMessage) NoErrors() bool {
	return len(f.Errors) == 0
}

// NoMessages func checks if there are message entries
func (f *FlexMessage) NoMessages() bool {
	return len(f.Messages) == 0
}

// Error func creates a new Error entry
func (f *FlexMessage) Error(err string) []string {
	f.Errors = append(f.Errors, err)
	return f.Errors
}

// Message func creates a new Message entry
func (f *FlexMessage) Message(msg string) []string {
	f.Messages = append(f.Messages, msg)
	return f.Messages
}

// Reset func both Errors and Messages
func (f *FlexMessage) Reset() {
	var zeroF = &FlexMessage{}
	*f = *zeroF
}

// Compact func makes FlexMessage tiny :-)
func (f *FlexMessage) Compact() *map[string]interface{} {
	message := make(map[string]interface{})

	if !f.NoErrors() {
		if len(f.Errors) == 1 {
			message["error"] = strings.Join(f.Errors, ",")
		} else {
			message["errors"] = f.Errors
		}
	}

	if !f.NoMessages() {
		if len(f.Messages) == 1 {
			message["message"] = strings.Join(f.Messages, ",")
		} else {
			message["messages"] = f.Messages
		}
	}

	return &message
}

// ExitCode not 0 when there are errors
func (f *FlexMessage) ExitCode() int {
	if f.NoErrors() {
		return 0
	}
	return f.ErrorExitCode
}
