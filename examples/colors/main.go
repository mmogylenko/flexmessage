package main

import (
	"fmt"

	"github.com/mmogylenko/flexmessage"
)

// Foo function returns an error
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

	if !notifications.Empty() {
		options := &flexmessage.PrintOptions{}
		commentColoringOptions := &flexmessage.ColoringSchema{StringColor: "green", RawStrings: true}

		fmt.Println(
			flexmessage.Colorize("// Compact = false, Colors = false, Indent = 0", commentColoringOptions),
		)
		notifications.Print(options)

		fmt.Println(
			flexmessage.Colorize("// Compact = true, Colors = false, Indent = 0", commentColoringOptions),
		)
		options.Compact = true
		notifications.Print(options)

		fmt.Println(
			flexmessage.Colorize("// Compact = true, Colors = true, Indent = 0 ", commentColoringOptions),
		)

		options.Colors = true
		notifications.Print(options)

		fmt.Println(
			flexmessage.Colorize("// Compact = true, Colors = true, Indent = 4 ", commentColoringOptions),
		)

		options.Indent = 4
		notifications.Print(options)
	}
}
