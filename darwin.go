package main

import (
	"Darwin/cognition"
	"Darwin/consciousness"
	"Darwin/output"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	command = kingpin.Arg("command", "🔎 Command for Darwin to execute").String()
)

func main() {
	kingpin.Parse()

	if strings.ToLower(*command) == "listen" || *command == "" {
		beginListeningMode()
	} else {
		executeCommandDirectly()
	}
}

func beginListeningMode() {
	output.PrintOutput("I'm all ears")
	consciousness.PayAttention()
}

func executeCommandDirectly() {
	intent, err := cognition.ResolveIntent(*command)

	if intent == nil {
		consciousness.HandleMissingIntent(intent)
	} else {
		consciousness.Respond(err, *command, intent, intent.Response(*command))
	}
}
