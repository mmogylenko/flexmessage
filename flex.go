package flexmessage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// FlexMessage type used for Msg/Err representation
type FlexMessage struct {
	Messages []string `json:"messages,omitempty"`
	Errors   []string `json:"errors,omitempty"`
}

// PrintOptions messages print options
type PrintOptions struct {
	Colors  bool
	Compact bool
	Indent  int
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

// Flush func clears both Errors and Messages
func (f *FlexMessage) Flush() {
	var zeroF = &FlexMessage{}
	*f = *zeroF
}

// Compact func makes FlexMessage tiny :-)
func (f *FlexMessage) Compact() map[string]interface{} {
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

	return message
}
func (f *FlexMessage) toMap() map[string]interface{} {
	return map[string]interface{}{
		"Messages": f.Messages,
		"Errors":   f.Errors,
	}
}

// Print func prints messages
func (f *FlexMessage) Print(options *PrintOptions) {
	var s []byte
	var compact map[string]interface{}

	if !f.NoMessages() {
		if options.Compact {
			compact = f.Compact()
		} else {
			compact = f.toMap()
		}

		if options.Indent > 0 {
			s, _ = json.MarshalIndent(compact, "", strings.Repeat(" ", options.Indent))
		} else {
			s, _ = json.Marshal(compact)
		}

		if options.Colors {
			out := Colorize(compact, &ColoringSchema{Indent: options.Indent})
			fmt.Println(string(out))
		} else {

			fmt.Println(string(s))
		}

	}

}

// Colorize just add some colors...
func Colorize(obj interface{}, schema *ColoringSchema) string {
	var b = bytes.Buffer{}
	cs := schema.New()

	cs.colorizeValue(obj, &b, initialDepth)

	return string(b.Bytes())
}
