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
	configFlag bool
	testFlag   bool
)

func Args() {
	flag.StringVar(&flagInput, "sync", "", "sync the package with your system")
	flag.StringVar(&desyncFlag, "desync", "", "desync the package with ")
	flag.StringVar(&updateFlag, "update", "", "update the package or system")
	flag.BoolVar(&newsFlag, "news", false, "display operating system news")
	flag.BoolVar(&configFlag, "config", false, "configurate package manager")
	flag.BoolVar(&testFlag, "test", false, "test")
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

	if configFlag {
		ConfigurateManager()
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
