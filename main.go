package main

import (
	"log"
	"runtime"

	"github.com/rendick/pem/cmd"
)

func main() {
	if runtime.GOOS == "linux" {
		cmd.Switch()
	} else {
		log.Printf("You are not running Linux!")
	}
}
