package clui

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
)

const (
	show valueVariant = iota
	ask
)

// UI contains functionality for dealing with the user
// on the CLI
type UI struct {
}

// Message represents a piece of information we want displayed to the user
type Message struct {
	ui           *UI // For access to requested verbosity.
	msgType      msgType
	end          int
	compact      bool
	noNewline    bool
	interactions []interaction
	tableHeaders [][]string
	tableData    [][][]string
}

type interaction struct {
	variant   valueVariant
	valueType valueType
	name      string
	value     interface{}
}

// NewUI creates a new UI
func NewUI() *UI {
	return &UI{}
}
