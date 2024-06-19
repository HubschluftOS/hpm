package cmd

import (
	"errors"
	"hpm/modules"
	"os"
	"strings"
)

func CheckIfFileExist(pkg string) bool {
	if _, err := os.Stat("/usr/bin/" + strings.TrimSpace(strings.ToLower(string(pkg)))); err == nil {
		if RemovePackage := os.Remove("/usr/bin/" + strings.TrimSpace(strings.ToLower(string(pkg)))); RemovePackage != nil {
			modules.Error("Error while removing %s", pkg)
			return false
		} else {
			modules.Success("%s successfully uninstalled.", pkg)
			return true
		}
	} else if errors.Is(err, os.ErrNotExist) {
		modules.Error("File does not exists. Aborting.")
		return false
	} else {
		modules.Error("Error checking if file exists: %v", err)
		return false
	}
}

func Unget(pkg string) {
	if modules.IsSudo() == true {
		CheckIfFileExist(pkg)
	} else {
		modules.Error("Unable to get current user. Aborting.")
		return
	}
}
