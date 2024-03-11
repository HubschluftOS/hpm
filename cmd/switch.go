package cmd

import (
	"fmt"
	config "hpm/settings"
	"os"
)

var Log string

var (
	InstallLog = "\n%s: missing package\nUsage: hpm install [PACKAGE]\nTry help for more information\n"
	RemoveLog  = "\n%s: missing package\nUsage: hpm remove [PACKAGE]\nTry help for more information\n"
)

func Switch() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "install":
			Log = fmt.Sprintf(config.Red+"(%s)"+config.Reset+InstallLog, config.Time, os.Args[1])
			fmt.Println(Log)
			Logs()
			os.Exit(0)

		case "list":
			Scrapper()

		case "update":
			fmt.Println("Soon.")
			os.Exit(0)

		case "remove":
			Log = fmt.Sprintf(config.Red+"(%s)"+config.Reset+RemoveLog, config.Time, os.Args[1])
			fmt.Println(Log)
			Logs()
			os.Exit(0)
		case "stats":
			PackageStats()

		case "help":
			config.Help()
			os.Exit(0)
		case "version":
			fmt.Println(config.Version)
			os.Exit(0)

		}
	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		case "install":
			InstallPackage(os.Args[2])
			Log = fmt.Sprintf(config.Red+"\n%s: "+config.Reset+"%s\n", config.Time, os.Args[2])
			fmt.Println(Log)
			Logs()
			os.Exit(0)

		case "remove":
			UninstallPackage(os.Args[2])
			// Log = fmt.Sprintf(config.Red+"\n%s: "+config.Reset+"%s\n", config.Time, os.Args[2])
			// fmt.Println(Log)
			// Logs()
			os.Exit(0)

		}
	} else {
		config.Missing()
	}
}
