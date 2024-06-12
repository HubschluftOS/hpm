package cmd

import (
	"flag"
	"fmt"
)

var (
	GetFlag   string
	UngetFlag string
	NewsFlag  bool
)

func Cli() {
	flag.StringVar(&GetFlag, "get", "", "Install the needed package on your Linux system")
	flag.StringVar(&UngetFlag, "unget", "", "Uninstall the package on your Linux system")
	flag.BoolVar(&NewsFlag, "news", false, "Display the latest news of the Hubshluft team and HubshluftOS.")
	flag.Parse()

	if GetFlag != "" {
		Get(GetFlag)
		return
	}

	if UngetFlag != "" {
		Unget(UngetFlag)
		return
	}

	if NewsFlag {
		News()
		return
	}

	nonFlagsArgs := flag.Args()
	if len(nonFlagsArgs) == 0 {
		fmt.Printf("hpm: not enough arguments\nTry 'hpm --help' for more information.\n")
		return
	}
}
