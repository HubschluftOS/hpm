package cmd

import (
	"flag"
	"fmt"
)

var (
	flagInput  string
	desyncFlag string
	updateFlag string
	newsFlag   bool
	configFlag string
	logsFlag   string
	searchFlag string
	testFlag   bool
)

func Args() {
	flag.StringVar(&flagInput, "sync", "", "sync the package with your system")
	flag.StringVar(&desyncFlag, "desync", "", "desync the package with your system")
	flag.StringVar(&updateFlag, "update", "", "update the package or system")
	flag.BoolVar(&newsFlag, "news", false, "display operating system news")
	flag.StringVar(&configFlag, "config", "", "configure the package manager")
	flag.StringVar(&logsFlag, "logs", "", "logs")
	flag.StringVar(&searchFlag, "search", "", "search")
	flag.BoolVar(&testFlag, "test", false, "run tests")
	flag.Parse()

	if flagInput != "" {
		GetPackageInformation(flagInput)
		return
	}

	if desyncFlag != "" {
		Desync(desyncFlag)
		return
	}

	if updateFlag == "@world" {
		UpdateSystem()
		return
	} else if updateFlag != "" {
		UpdatePackage(updateFlag)
		return
	}

	if newsFlag {
		News()
		return
	}

	if configFlag == "reload" {
		reloadConfig()
		return
	} else if configFlag == "create" {
		ConfigurateManager()
		return
	} else if configFlag == "remove" {
		removeConfig()
		return
	}

	if logsFlag == "display" {
		DisplayLogs()
		return
	} else if logsFlag == "delete" {
		fmt.Println("delete")
		return
	}

	if searchFlag != "" {
		Search(searchFlag)
		return
	}

	if testFlag {
		Test()
		return
	}

	nonFlagsArgs := flag.Args()
	if len(nonFlagsArgs) == 0 {
		fmt.Printf("hpm: not enough arguments\nTry 'hpm --help' for more information.\n")
		return
	}
}
