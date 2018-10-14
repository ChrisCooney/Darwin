package intent

import (
	"os"
	"strings"
)

type SleepIntent struct{}

func (intent SleepIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), "go to sleep") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "quit") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "exit") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "sleep") {
		return 50
	}

	if strings.Contains(strings.ToLower(input), "thanks") {
		return 25
	}

	return 0
}

func (intent SleepIntent) Response(input string) string {
	if strings.Contains(strings.ToLower(input), "thanks") {
		return "You're welcome."
	}
	return "Good night"
}

func (intent SleepIntent) Name() string {
	return "Sleep Intent"
}

func (intent SleepIntent) Usage() string {
	return "Go to sleep"
}

func (intent SleepIntent) PostSpeech(input string) {
	os.Exit(0)
}
