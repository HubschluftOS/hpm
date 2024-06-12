package cmd

import (
	"errors"
	"fmt"
	"hpm/modules"
	"os"
	"strings"
)

func Unget(pkg string) {
	if modules.IsBool() == true {
		if _, err := os.Stat("/usr/bin/" + strings.TrimSpace(strings.ToLower(string(pkg)))); err == nil {
			if RemovePackage := os.Remove("/usr/bin/" + strings.TrimSpace(strings.ToLower(string(pkg)))); RemovePackage != nil {
				modules.Error("Error while removing %s", pkg)
				return
			}
		} else if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("File does not exists.\n")
			return
		}
	} else {
		modules.Error("Unable to get current user. Aborting.")
		return
	}
}
