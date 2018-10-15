package applications

import (
	"github.com/everdev/mack"
)

func OpenApplication(name string) {
	mack.Tell(name, "open")
}

func OpenApplicationAsync(name string, responseChannel chan string) {
	OpenApplication(name)
	responseChannel <- name
}

func CloseApplication(name string) {
	mack.Tell(name, "quit")
}
