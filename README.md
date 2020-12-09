# FlexMessage

![Go](https://github.com/mmogylenko/flexmessage/workflows/Go/badge.svg) ![Gosec](https://github.com/mmogylenko/flexmessage/workflows/Gosec/badge.svg) [![GitHub Super-Linter](https://github.com/mmogylenko/flexmessage/workflows/Lint%20Code%20Base/badge.svg)](https://github.com/marketplace/actions/super-linter) [![GitHub tag](https://img.shields.io/github/tag/mmogylenko/flexmessage.svg)](https://github.com/mmogylenko/flexmessage/tags/)

`FlexMessage` - Notifications in a nice way :-)

![gopher](https://github.com/egonelbre/gophers/blob/master/sketch/fairy-tale/messenger-red-letter.png?raw=true)

## Motivation

>I don't know. But why not?

The idea was followed by [golang declaration syntax](https://blog.golang.org/declaration-syntax). When the `HandleFunc` function from `net/http` is not explicitly returning something, why I should use `return` statement? So I went from this
```go
func StatusHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method != "GET" {
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(`{"message": "not implemented"}`))
        return
    }
    w.Write([]byte(`{"message": "Ok"}`))
}
```
to this
```go
func StatusHandler(w http.ResponseWriter, r *http.Request) {
    var notify flexmessage.FlexMessage

    w.Header().Set("Content-Type", "application/json")
    if r.Method != "GET" {
        w.WriteHeader(http.StatusNotImplemented)
        notify.Error(r.Method + " method is not implemented")
    }
    if notify.NoErrors() {
        notify.Message("Ok")
    }
    json.NewEncoder(w).Encode(notify.Compact())
}
```
More complicated handler can reveal the benefits of `FlexMessage` package.

*Key Features*

- It *MAY WORK* or may not
- Simplicity is a priority
- Colors!!! I love colors!


## Installation

```sh
go get -u github.com/mmogylenko/flexmessage
```

## Usage


```go
package main

import (
    "fmt"
    "github.com/mmogylenko/flexmessage"
)

var notify flexmessage.FlexMessage

func main() {
    if 10 > 9 {
        notify.Message("Surprise!")
    }
    if !notify.NoMessages() {
        fmt.Println(notify.Messages)
    }
}
```

## Customization


- `Compact()` - Returns single Message/Error when `len(Errors) == 1 ||  len(Messages) == 1`

```go
package main

import (
    "encoding/json"
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

    // Check if need to notify
    if !notifications.Empty() {
        fmt.Println("Output without Compact()")
        s, _ := json.MarshalIndent(notifications, "", "  ")
        fmt.Println(string(s))
        fmt.Println("Output with Compact()")
        c, _ := json.MarshalIndent(notifications.Compact(), "", "  ")
        fmt.Println(string(c))
    }
}
```

Result:
```json
Output without Compact()
{
  "messages": [
    "Very important message"
  ],
  "errors": [
    "Foo error"
  ]
}
Output with Compact()
{
  "error": "Foo error",
  "message": "Very important message"
}
```

- `Print()` - Prints messages type `PrintOptions` used for print options:

```go
// PrintOptions messages print options
type PrintOptions struct {
	Colors  bool
	Compact bool
	Indent  int
}
```
[Example](examples) of `Print()`:
![flexmesssage_colors_example](https://user-images.githubusercontent.com/7536624/101554300-3184c200-396b-11eb-8256-cc60e2ff3b1b.png)


## Examples

Check [examples](examples) directory for a use-cases of `flexmessage` usage


## Licensing

This project is licensed under the Apache V2 License. See [LICENSE](LICENSE) for more information.
