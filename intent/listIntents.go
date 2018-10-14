package intent

import (
	"bytes"
	"fmt"
	"strings"
)

type ListIntentionsIntent struct{}

func (intent ListIntentionsIntent) Name() string {
	return "List Intentions Intent"
}

func (intent ListIntentionsIntent) Usage() string {
	return "What can you do?"
}

func (intent ListIntentionsIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), "list intentions") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "list intents") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "what can you do") {
		return 100
	}

	return 0
}

func (intent ListIntentionsIntent) Response(input string) string {
	var buffer bytes.Buffer

	buffer.WriteString("I am programmed with the following intents: \n")

	for _, intent := range IntentsList {
		buffer.WriteString(fmt.Sprintf("-> %s - Example: %s \n", intent.Name(), intent.Usage()))
	}

	return strings.TrimRight(buffer.String(), " \n")
}

func (intent ListIntentionsIntent) PostSpeech(input string) {
	// No op
}
