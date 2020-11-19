package flexmessage

import (
	"strings"
)

// FlexMessage type used for Msg/Err representation
type FlexMessage struct {
	Messages []string `json:"messages,omitempty"`
	Errors   []string `json:"errors,omitempty"`
}

// Empty method is check if there are any notifications
// both messages and errors
func (f *FlexMessage) Empty() bool {
	if len(f.Errors) > 0 || len(f.Messages) > 0 {
		return false
	}
	return true
}

// NoErrors method checks if there are error entries
func (f *FlexMessage) NoErrors() bool {
	if len(f.Errors) > 0 {
		return false
	}
	return true
}

// NoMessages method checks if there are message entries
func (f *FlexMessage) NoMessages() bool {
	if len(f.Messages) > 0 {
		return false
	}
	return true
}

// Error method creates a new Error entry
func (f *FlexMessage) Error(err string) []string {
	f.Errors = append(f.Errors, err)
	return f.Errors
}

// Message method creates a new Message entry
func (f *FlexMessage) Message(msg string) []string {
	f.Messages = append(f.Messages, msg)
	return f.Messages
}

// Compact method makes FlexMessage tiny :-)
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
