package main

import (
	"encoding/json"

	"github.com/mmogylenko/flexmessage"

	"fmt"
)

// Foo function returns error
func Foo() error {
	return fmt.Errorf("Foo error")
}

func main() {

	var notifications flexmessage.FlexMessage

	if 10 > 0 {
		// We add our 1st Message
		notifications.Message("Very important message")
	}

	err := Foo()

	if err != nil {
		// We add our 1st Error
		notifications.Error(err.Error())
	}

	// Check if need to notify
	if !notifications.Empty() {
		c, _ := json.MarshalIndent(notifications.Compact(), "", "  ")
		fmt.Println(string(c))
	}

}
