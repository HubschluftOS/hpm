package cmd

import (
	"flag"
	"fmt"
)

func Args() {
	var flagInput string
	var desyncFlag string
	var newsFlag bool
	flag.StringVar(&flagInput, "sync", "", "sync the package with your system")
	flag.StringVar(&desyncFlag, "desync", "", "desync the package with ")
	flag.BoolVar(&newsFlag, "news", false, "display operating system news")
	flag.Parse()

	if flagInput != "" {
		GetPackageInformation(flagInput)
	}

	if desyncFlag != "" {
		Desync()
	}

	if newsFlag {
		News()
		return
	}

	nonFlagsArgs := flag.Args()
	if len(nonFlagsArgs) == 0 {
		fmt.Printf("hpm: not enough arguments\nTry 'hpm --help' for more information.\n")
		return
	}
}
