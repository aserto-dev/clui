package clui

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

// WithAskBool waits for the user's input for a boolean value
func (u *Message) WithAskBool(name string, result *bool) *Message {
	return u.WithAskBoolMap(name, result, map[string]bool{
		"true":  true,
		"false": false,
	})
}

// WithAskBool waits for the user's input for a boolean value
func (u *Message) WithAskBoolMap(name string, result *bool, answerMap map[string]bool) *Message {
	u.interactions = append(u.interactions, interaction{
		name:      name,
		variant:   ask,
		valueType: tBool,
		value:     result,
		boolMap:   answerMap,
	})
	return u
}

// WithAskString waits for the user's input for a string value
func (u *Message) WithAskString(name string, result *string) *Message {
	u.interactions = append(u.interactions, interaction{
		name:      name,
		variant:   ask,
		valueType: tString,
		value:     result,
	})
	return u
}

// WithAskInt waits for the user's input for an int value
func (u *Message) WithAskInt(name string, result *int64) *Message {
	u.interactions = append(u.interactions, interaction{
		name:      name,
		variant:   ask,
		valueType: tInt,
		value:     result,
	})
	return u
}

func (u *Message) readBool(message string, boolMap map[string]bool) bool {
	if !strings.HasSuffix(message, "?") && !strings.HasSuffix(message, ":") {
		message = message + ":"
	}

	if len(boolMap) == 0 {
		boolMap = map[string]bool{
			"true":  true,
			"false": false,
		}
	}

	scanner := bufio.NewScanner(u.ui.input)
	for {
		u.ui.printf("[%s] %s ", color.MagentaString("bool"), emoji.Sprint(message))
		scanner.Scan()
		text := scanner.Text()

		value, ok := boolMap[strings.ToLower(text)]

		if !ok {
			u.ui.Problem().WithStringValue("  input", text).Msg("Invalid value.")
			continue
		}

		return value
	}
}

func (u *Message) readString(message string) string {
	if !strings.HasSuffix(message, "?") && !strings.HasSuffix(message, ":") {
		message = message + ":"
	}

	u.ui.printf("[%s] %s ", color.GreenString("text"), emoji.Sprint(message))

	scanner := bufio.NewScanner(u.ui.input)
	scanner.Scan()
	value := scanner.Text()
	return value
}

func (u *Message) readInt(message string) int64 {
	if !strings.HasSuffix(message, "?") && !strings.HasSuffix(message, ":") {
		message = message + ":"
	}

	var err error
	var result int64

	scanner := bufio.NewScanner(u.ui.input)
	for {
		u.ui.printf("[%s] %s ", color.CyanString("integer"), emoji.Sprint(message))
		scanner.Scan()
		text := scanner.Text()

		result, err = strconv.ParseInt(text, 10, 64)

		if err != nil {
			u.ui.Problem().WithStringValue("  input", text).Msg("Value is not an integer.")
			continue
		}

		return result
	}
}
