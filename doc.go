/*
Package flexmessage is a simple way to deal with Notifications like Errors or Messages.
Personally found it pretty useful:
	package main
	import (
		"encoding/json"
		"github.com/mmogylenko/flexmessage"

		"fmt"
	)
	func Foo() error {
		return fmt.Errorf("Foo error")
	}
	func main() {
		var notifications flexmessage.FlexMessage
		err := Foo()
		if err != nil {
			notifications.Error(err.Error())
		}
		if !notifications.Empty() {
			c, _ := json.MarshalIndent(notifications.Compact(), "", "  ")
			fmt.Println(string(c))
		}
	}
Output:
	{
		"error": "Foo error"
	}
For a full guide visit https://github.com/mmogylenko/flexmessage
*/
package flexmessage
