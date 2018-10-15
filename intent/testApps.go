package intent

import (
	"Darwin/applications"
	"Darwin/output"
	"fmt"
	"strings"
)

type TestApplicationsIntent struct{}

func (intent TestApplicationsIntent) Name() string {
	return "Test Applications Intent"
}

func (intent TestApplicationsIntent) Usage() string {
	return "Are my apps okay?"
}

func (intent TestApplicationsIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), "are my work applications okay?") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "test work applications") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "applications okay") {
		return 50
	}

	if strings.Contains(strings.ToLower(input), "apps okay") {
		return 50
	}

	return 0
}

func (intent TestApplicationsIntent) Response(input string) string {
	responseChannel := make(chan applications.ApplicationHealth)
	testAppsInParallel(responseChannel)
	output.PrintCountSummary("Application Healthcheck", reportAppHealth(responseChannel))
	return "Health check completed"
}

func testAppsInParallel(responseChannel chan applications.ApplicationHealth) {
	for _, app := range applications.CommandLineApps() {
		go applications.IsHealthyAsync(app.BrewPackageName, app.HealthCheckCommand, responseChannel)
	}
}

func reportAppHealth(responseChannel chan applications.ApplicationHealth) output.CountSummary {

	healthyCount := 0

	apps := applications.CommandLineApps()

	for x := 0; x < len(apps); x++ {
		appHealth := <-responseChannel

		if appHealth.Health == true {
			output.PrintOutput(fmt.Sprintf("%s is healthy.", appHealth.AppName))
			healthyCount++
		} else {
			output.PrintError(fmt.Sprintf("%s is not healthy", appHealth.AppName))
		}
	}

	return output.CountSummary{Count: healthyCount, Total: len(apps), Message: "command line applications healthy."}
}

func (intent TestApplicationsIntent) PostSpeech(input string) {
	// No op
}
