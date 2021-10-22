package clui

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/fatih/color"
)

type msgType int
type valueType int
type valueVariant int

const (
	normal msgType = iota
	exclamation
	problem
	note
	success
	progress
)

const (
	tBool valueType = iota
	tString
	tInt
	tErr
	tPassword
)

const (
	show valueVariant = iota
	ask
)

// UI contains functionality for dealing with the user
// on the CLI
type UI struct {
	output io.Writer
	input  io.Reader
}

// Message represents a piece of information we want displayed to the user
type Message struct {
	ui           *UI // For access to requested verbosity.
	msgType      msgType
	end          int
	compact      bool
	noNewline    bool
	stacks       bool
	interactions []interaction
	tableHeaders [][]string
	tableData    [][][]string
}

type interaction struct {
	variant   valueVariant
	valueType valueType
	name      string
	value     interface{}
	boolMap   map[string]bool
	stdin     bool
}

// NewUI creates a new UI
func NewUI() *UI {
	if runtime.GOOS == "windows" {
		return NewUIWithOutput(color.Output)
	} else {
		return NewUIWithOutput(os.Stdout)
	}
}

// NewUI creates a new UI with a specific output
func NewUIWithOutput(output io.Writer) *UI {
	return NewUIWithOutputAndInput(output, os.Stdin)
}

// NewUI creates a new UI with a specific input and output
func NewUIWithOutputAndInput(output io.Writer, input io.Reader) *UI {
	return &UI{
		output: output,
		input:  input,
	}
}

// Input returns the io.Reader used to read user input
func (u *UI) Input() io.Reader {
	return u.input
}

// Output returns the io.Write used to print
func (u *UI) Output() io.Writer {
	return u.output
}

func (u *UI) printf(format string, args ...interface{}) {
	u.output.Write([]byte(fmt.Sprintf(format, args...)))
}

func (u *UI) println(args ...interface{}) {
	u.output.Write([]byte(fmt.Sprintln(args...)))
}
