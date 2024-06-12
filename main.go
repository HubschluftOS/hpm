package main

import (
	"fmt"
	"hpm/cmd"
	"runtime"
)

func main() {
	os_slice := []string{"linux"}
	os_type := false
	for _, str := range os_slice {
		if str == runtime.GOOS {
			os_type = true
			break
		}
	}

	if os_type == true {
		cmd.Cli()
	} else {
		fmt.Printf("You are not running UNIX-like systme right now: %s %s\n", runtime.GOOS, runtime.GOARCH)
	}
}
