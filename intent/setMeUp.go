package intent

import (
	"Darwin/applications"
	"Darwin/output"
	"fmt"
	"strings"
)

var BrewCaskInstallations = []string{"virtualbox", "minikube"}

type SetMeUpIntent struct{}

func (intent SetMeUpIntent) Name() string {
	return "Clean Setup Intent"
}

func (intent SetMeUpIntent) Usage() string {
	return "Lets get set up"
}

func (intent SetMeUpIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), "lets get set up") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "set me up") {
		return 100
	}

	return 0
}

func (intent SetMeUpIntent) Response(input string) string {
	output.PrintOutput("Commencing initialisation of new development machine.")

	output.PrintOutput("Checking if Homebrew exists...")

	if !applications.IsHomebrewInstalled() {
		output.PrintError("Please install Homebrew before running this command")
		return "Installation cancelled. Homebrew not present"
	}

	output.PrintOutput("Homebrew exists. Proceeding with installation.")

	output.PrintCountSummary("Installation Summary", installCommandLineApps(), installCaskApps())
	return ""
}

func installCommandLineApps() output.CountSummary {
	successful := 0

	apps := applications.CommandLineApps()

	for _, app := range apps {
		applications.InstallBrewApp(app.BrewPackageName)

		if applications.IsHealthy(app.HealthCheckCommand) {
			successful++
			output.PrintOutput(fmt.Sprintf("Installed %s", app.BrewPackageName))
		} else {
			output.PrintError(fmt.Sprintf("Error installing %s.", app))
		}
	}
	return output.CountSummary{Count: successful, Total: len(apps), Message: "command line applications installed."}
}

func installCaskApps() output.CountSummary {
	successful := 0
	for _, app := range BrewCaskInstallations {
		err := applications.InstallBrewCaskApp(app)

		// Can't invoke healthcheck from command line for GUI apps so rely on brew.
		if err != nil {
			output.PrintError(fmt.Sprintf("Error installing %s. %s", app, err.Error()))
		} else {
			successful++
			output.PrintOutput(fmt.Sprintf("Installed %s", app))
		}
	}
	return output.CountSummary{Count: successful, Total: len(BrewCaskInstallations), Message: "GUI applications installed."}
}

func (intent SetMeUpIntent) PostSpeech(input string) {
	// No op
}
