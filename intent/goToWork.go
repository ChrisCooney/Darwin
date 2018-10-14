package intent

import (
	"Darwin/applications"
	"Darwin/output"
	"fmt"
	"strings"
)

var WorkApplications = []string{"Spotify", "IntelliJ IDEA", "Visual Studio Code 2", "Slack", "Microsoft Teams", "Microsoft Outlook", "Skype for Business"}

type GoToWorkIntent struct{}

func (intent GoToWorkIntent) Name() string {
	return "Go to work"
}

func (intent GoToWorkIntent) Usage() string {
	return "Let's go to work!"
}

func (intent GoToWorkIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), "lets go to work") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "let's go to work") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "go to work") {
		return 50
	}

	if strings.Contains(strings.ToLower(input), " work ") {
		return 25
	}

	if strings.Contains(strings.ToLower(input), " work") {
		return 10
	}

	return 0
}

func (intent GoToWorkIntent) Response(input string) string {
	output.PrintOutput("Opening Work Applications")
	for _, app := range WorkApplications {
		applications.OpenApplication(app)
		output.PrintOutput(fmt.Sprintf("%s opened", app))
	}

	return "Have fun"
}

func (intent GoToWorkIntent) PostSpeech(input string) {
	// No op
}
