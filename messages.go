package clui

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
	"github.com/olekukonko/tablewriter"
)

// Normal returns a UIMessage that prints a normal message
func (u *UI) Normal() *Message {
	return &Message{
		ui:           u,
		msgType:      normal,
		interactions: []interaction{},
		end:          -1,
	}
}

// Exclamation returns a UIMessage that prints an exclamation message
func (u *UI) Exclamation() *Message {
	return &Message{
		ui:           u,
		msgType:      exclamation,
		interactions: []interaction{},
		end:          -1,
	}
}

// Note returns a UIMessage that prints a note message
func (u *UI) Note() *Message {
	return &Message{
		ui:           u,
		msgType:      note,
		interactions: []interaction{},
		end:          -1,
	}
}

// Success returns a UIMessage that prints a success message
func (u *UI) Success() *Message {
	return &Message{
		ui:           u,
		msgType:      success,
		interactions: []interaction{},
		end:          -1,
	}
}

// Problem returns a Message that prints a message that describes a problem
func (u *UI) Problem() *Message {
	return &Message{
		ui:           u,
		msgType:      problem,
		interactions: []interaction{},
		end:          -1,
	}
}

// Msgf prints a formatted message on the CLI
func (u *Message) Msgf(message string, a ...interface{}) {
	u.Msg(fmt.Sprintf(message, a...))
}

// Do is syntactic sugar for Msg("")
func (u *Message) Do() {
	u.Msg("")
}

// Msg prints a message on the CLI, resolving emoji as it goes
func (u *Message) Msg(message string) {
	message = emoji.Sprint(message)

	// Print a newline before starting output, if not compact.
	if message != "" && !u.compact {
		u.ui.println()
	}

	if !u.noNewline {
		message += "\n"
	}

	switch u.msgType {
	case normal:
	case exclamation:
		message = emoji.Sprintf(":warning: %s", message)
		message = color.YellowString(message)
	case note:
		message = emoji.Sprintf(":ship:%s", message)
		message = color.BlueString(message)
	case success:
		message = emoji.Sprintf(":heavy_check_mark: %s", message)
		message = color.GreenString(message)
	case progress:
		message = emoji.Sprintf(":three-thirty: %s", message)
	case problem:
		message = emoji.Sprintf(":cross_mark: %s", message)
		message = color.RedString(message)
	}

	u.ui.printf("%s", message)

	for _, interaction := range u.interactions {
		switch interaction.variant {
		case ask:
			u.ui.printf("> ")
			switch interaction.valueType {
			case tBool:
				*(interaction.value.(*bool)) = u.readBool(interaction.name)
			case tInt:
				*(interaction.value.(*int64)) = u.readInt(interaction.name)
			case tString:
				*(interaction.value.(*string)) = u.readString(interaction.name)
			}
		case show:
			switch interaction.valueType {
			case tBool:
				u.ui.printf("%s: %s\n", emoji.Sprint(interaction.name), color.MagentaString("%t", interaction.value))
			case tInt:
				u.ui.printf("%s: %s\n", emoji.Sprint(interaction.name), color.CyanString("%d", interaction.value))
			case tString:
				u.ui.printf("%s: %s\n", emoji.Sprint(interaction.name), color.GreenString("%s", interaction.value))
			case tErr:
				u.ui.printf("%s\n", color.RedString("%+v", interaction.value))
			}
		}
	}

	for idx, headers := range u.tableHeaders {
		table := tablewriter.NewWriter(u.ui.output)
		table.SetHeader(headers)
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")

		if idx < len(u.tableData) {
			table.AppendBulk(u.tableData[idx])
		}

		table.Render()
	}

	if u.end > -1 {
		os.Exit(u.end)
	}
}

// NoNewline disables the printing of a newline after a message output
func (u *Message) NoNewline() *Message {
	u.noNewline = true
	return u
}

// Compact disables the printing of a newline before starting output
func (u *Message) Compact() *Message {
	u.compact = true
	return u
}

// WithEnd ends the entire process after printing the message.
func (u *Message) WithEnd(code int) *Message {
	u.end = code
	return u
}

// WithBoolValue adds a bool value to be printed in the message
func (u *Message) WithBoolValue(name string, value bool) *Message {
	u.interactions = append(u.interactions, interaction{
		name:      name,
		variant:   show,
		valueType: tBool,
		value:     value,
	})
	return u
}

// WithStringValue adds a string value to be printed in the message
func (u *Message) WithStringValue(name string, value string) *Message {
	u.interactions = append(u.interactions, interaction{
		name:      name,
		variant:   show,
		valueType: tString,
		value:     value,
	})
	return u
}

// WithErr adds an error value to be printed in the message
func (u *Message) WithErr(err error) *Message {
	u.interactions = append(u.interactions, interaction{
		variant:   show,
		valueType: tErr,
		value:     err,
	})
	return u
}

// WithIntValue adds an int value to be printed in the message
func (u *Message) WithIntValue(name string, value int64) *Message {
	u.interactions = append(u.interactions, interaction{
		name:      name,
		variant:   show,
		valueType: tInt,
		value:     value,
	})
	return u
}
