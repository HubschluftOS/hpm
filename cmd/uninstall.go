package cmd

// alpha version

import (
	"fmt"
	"log"
	"os"

	config "github.com/rendick/pem/settings"
)

var Remove string

func UninstallPacage(packageName string) {
	Remove = fmt.Sprintf("%s%s", config.PackageDir, packageName)
	uninstall := os.Remove(Remove)
	if uninstall != nil {
		log.Printf(config.Red+config.Bold+config.ErrorMsg+config.Reset, packageName)
	} else {
		log.Printf(config.Green+config.Bold+config.SuccessMsg+config.Reset, packageName)
	}
}
