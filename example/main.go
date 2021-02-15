package main

import (
	"time"

	"github.com/aserto-dev/clui"
)

func main() {
	ui := clui.NewUI()
	ui.Normal().Msg("The quick brown fox jumps over the lazy dog!")

	ui.Note().Msg("The quick brown fox jumps over the lazy dog!")

	ui.Exclamation().Msg("The quick brown fox jumps over the lazy dog!")

	ui.Problem().Msg("The quick brown fox jumps over the lazy dog!")

	ui.Success().Msg("The quick brown fox jumps over the lazy dog!")

	var (
		boolResult   bool
		stringResult string
		intResult    int64
	)
	ui.Normal().
		WithAskBool("Do you want to work today?", &boolResult).
		WithAskString("What's your name?", &stringResult).
		WithAskInt("How old are you?", &intResult).
		Do()

	ui.Normal().
		WithStringValue("Your name", stringResult).
		WithBoolValue("Work is on", boolResult).
		WithIntValue("Your age is", intResult).
		Msg("Hello! These are the answers to the questionaire.")

	ui.Normal().WithTable("#", "Name", "Age", "Link").
		WithTableRow("1", "Stephen J. Fry", "20", "https://aserto.com").Do()

	p := ui.Progress("Doing something in the background")
	p.Start()
	defer p.Stop()
	time.Sleep(5 * time.Second)
}
