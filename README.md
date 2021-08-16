# clui

[![Go Reference](https://pkg.go.dev/badge/github.com/aserto-dev/clui.svg)](https://pkg.go.dev/github.com/aserto-dev/clui)
[![Go Report Card](https://goreportcard.com/badge/github.com/aserto-dev/clui)](https://goreportcard.com/report/github.com/aserto-dev/clui)

Command Line User Interface library for building user interfaces that interact with humans.
You can find a more comprehensive example [here](https://github.com/aserto-dev/clui/tree/main/example).

- [clui](#clui)
  - [Usage](#usage)
    - [Printing messages](#printing-messages)
    - [Asking for input](#asking-for-input)
    - [Spinner progress](#spinner-progress)
    - [Tables](#tables)


![demo](./demo.gif)

## Usage

### Printing messages

```go
ui := clui.NewUI()
ui.Normal().Msg("The quick brown fox jumps over the lazy dog!")
```

### Asking for input

```go
ui := clui.NewUI()

var stringResult string
ui.Normal().
  WithAskString("What's your name?", &stringResult).
  Do()
```


### Spinner progress

```go
ui := clui.NewUI()
p := ui.Progress("Doing something in the background")
p.Start()
defer p.Stop()
time.Sleep(5 * time.Second)
```

### Tables

```go
ui := clui.NewUI()
ui.Normal().WithTable("#", "Name", "Age", "Link").
  WithTableRow("1", "Stephen J. Fry", "20", "https://aserto.com").Do()
```
