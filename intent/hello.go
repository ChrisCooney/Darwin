package intent

import (
	"strings"
)

type HelloIntent struct{}

func (intent HelloIntent) Name() string {
	return "Hello Intent"
}

func (intent HelloIntent) Usage() string {
	return "Hello"
}

func (intent HelloIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), " hello ") {
		return 50
	}

	if strings.Contains(strings.ToLower(input), "hello") {
		return 10
	}

	return 0
}

func (intent HelloIntent) Response(input string) string {
	return "Well hello there."
}

func (intent HelloIntent) PostSpeech(input string) {
	// No op
}
