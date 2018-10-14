package intent

import (
	"Darwin/applications"
	"Darwin/output"
	"fmt"
	"strings"
)

type HomeTimeIntent struct{}

func (intent HomeTimeIntent) Name() string {
	return "Home time"
}

func (intent HomeTimeIntent) Usage() string {
	return "Home time"
}

func (intent HomeTimeIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), "home time") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "home") {
		return 50
	}

	return 0
}

func (intent HomeTimeIntent) Response(input string) string {
	output.PrintOutput("Closing work applications")
	for _, app := range WorkApplications {
		applications.CloseApplication(app)
		output.PrintOutput(fmt.Sprintf("%s closed", app))
	}

	return "Work applications closed"
}

func (intent HomeTimeIntent) PostSpeech(input string) {
	// No op
}
