package output

import (
	"os/exec"
)

func Speak(output string) {
	exec.Command("say", output).Run()
}
