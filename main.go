package main

import (
	"hpm/cmd"
	"log"
	"runtime"
)

func main() {
	if runtime.GOOS == "linux" {
		cmd.Switch()
	} else {
		log.Printf("You are not running Linux!")
	}
}
