package intent

import (
	"Darwin/applications"
	"fmt"
	"strings"
)

type OpenAppIntent struct{}

func (intent OpenAppIntent) Name() string {
	return "Open Application Intent"
}

func (intent OpenAppIntent) Usage() string {
	return "open IntelliJ Idea"
}

func (intent OpenAppIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), "open ") {
		return 100
	}

	return 0
}

func (intent OpenAppIntent) Response(input string) string {
	appName := grabAppName(input)
	applications.OpenApplication(appName)
	return fmt.Sprintf("%s opened", appName)
}

func (intent OpenAppIntent) PostSpeech(input string) {
	// No op
}

func grabAppName(input string) string {
	return strings.Replace(input, "open ", "", 1)
}
