package cmd

import (
	"log"
	"os"
)

func Desync(pkg string) {
	ConfigurateManager()
	if err := os.Remove("/usr/bin/" + pkg); err != nil {
		log.Fatal(err)
	}
}
