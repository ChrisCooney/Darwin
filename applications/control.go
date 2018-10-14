package applications

import (
	"github.com/everdev/mack"
)

func OpenApplication(name string) {
	mack.Tell(name, "open")
}

func CloseApplication(name string) {
	mack.Tell(name, "quit")
}
