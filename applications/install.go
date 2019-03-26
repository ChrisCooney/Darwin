package applications

import (
	"fmt"
	"os/exec"
	"strings"
)

type CommandLineApp struct {
	BrewPackageName    string
	HealthCheckCommand string
}

func InstallBrewApp(app string) error {
	// Check if need to install first
	return exec.Command("brew", "install", app).Run()
}

func InstallBrewCaskApp(app string) error {
	return exec.Command("brew", "cask", "install", app).Run()
}

func IsHomebrewInstalled() bool {
	return IsHealthy("brew --version")
}

func CommandLineApps() []CommandLineApp {
	apps := []CommandLineApp{}

	apps = append(apps, CommandLineApp{BrewPackageName: "node", HealthCheckCommand: "node --version"})
	apps = append(apps, CommandLineApp{BrewPackageName: "jdk8", HealthCheckCommand: "java -version"})
	apps = append(apps, CommandLineApp{BrewPackageName: "ansible", HealthCheckCommand: "ansible --version"})
	apps = append(apps, CommandLineApp{BrewPackageName: "python3", HealthCheckCommand: "python3 --version"})
	apps = append(apps, CommandLineApp{BrewPackageName: "maven", HealthCheckCommand: "mvn --version"})
	apps = append(apps, CommandLineApp{BrewPackageName: "gradle", HealthCheckCommand: "gradle --version"})
	apps = append(apps, CommandLineApp{BrewPackageName: "go", HealthCheckCommand: "go version"})
	apps = append(apps, CommandLineApp{BrewPackageName: "thefuck", HealthCheckCommand: "thefuck --version"})
	apps = append(apps, CommandLineApp{BrewPackageName: "kubernetes-cli", HealthCheckCommand: "kubectl config current-context"})
	apps = append(apps, CommandLineApp{BrewPackageName: "kubectx", HealthCheckCommand: "kubectx --help"})
	apps = append(apps, CommandLineApp{BrewPackageName: "kubernetes-helm", HealthCheckCommand: "helm --version"})

	return apps
}

func IsHealthy(command string) bool {
	tokens := strings.Split(command, " ")

	var err error

	if len(tokens) == 1 {
		err = exec.Command(tokens[0]).Run()
	} else {
		err = exec.Command(tokens[0], tokens[1:]...).Run()
	}

	if err != nil {
		fmt.Println(err.Error())

	}

	return err == nil
}

type ApplicationHealth struct {
	AppName string
	Health  bool
}

func IsHealthyAsync(appName string, command string, responseChannel chan ApplicationHealth) {
	isHealthy := IsHealthy(command)
	responseChannel <- ApplicationHealth{Health: isHealthy, AppName: appName}
}
