package main

import (
	"fmt"
	"hpm/cmd"
	"runtime"
)

func main() {
	os_slice := []string{"linux", "openbsd", "netbsd", "freebsd", "dragonfly"}
	slice_types := false
	for _, str := range os_slice {
		if str == runtime.GOOS {
			slice_types = true
			break
		}
	}

	if slice_types == true {
		cmd.Args()
	} else {
		fmt.Printf("You are not running UNIX-like system right now: %s\n", runtime.GOOS)
		return
	}
}
