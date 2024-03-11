package cmd

// alpha version

import (
	"fmt"
	config "hpm/settings"
	"log"
	"os"
)

var Remove string

func UninstallPackage(packageName string) {
	Remove = fmt.Sprintf("%s%s", config.PackageDir, packageName)
	uninstall := os.Remove(Remove)
	if uninstall != nil {
		log.Printf(config.Red+config.Bold+config.ErrorMsg+config.Reset+config.Bold+" sudo hpm remove %s"+config.Reset, packageName, packageName)
	} else {
		log.Printf(config.Green+config.Bold+config.SuccessMsg+config.Reset, packageName)
	}
}
