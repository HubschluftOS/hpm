package cmd

import (
	"flag"
	"fmt"
)

var (
	GetFlag    string
	UngetFlag  string
	NewsFlag   bool
	UpdateFlag string
	FindFlag   string
	HelpFlag   bool
)

func Cli() {
	flag.BoolVar(&HelpFlag, "help", false, "Display extra information about package manager.")
	flag.StringVar(&GetFlag, "get", "", "Install the needed package on your Linux system.")
	flag.StringVar(&UngetFlag, "unget", "", "Uninstall the package on your Linux system.")
	flag.StringVar(&UpdateFlag, "update", "", "Update all packages on your Linux system.")
	flag.StringVar(&FindFlag, "find", "", "Find a specific package in the repository.")
	flag.BoolVar(&NewsFlag, "news", false, "Display the latest news of the Hubshluft team and HubshluftOS.")
	flag.Parse()

	if HelpFlag {
		fmt.Printf("Usage: hpm [OPTION]... [PACKAGE]\n" +
			"Lightweight and powerful package manager for Hubshluft Linux and other GNU/Linux distributions\n\n" +
			"get\t Install the needed package on your Linux system\n" +
			"unget\t Uninstall the package on your Linux system\n" +
			"update\t Update all packages on your Linux system\n" +
			"news\tDisplay the latest news of the Hubshluft team and Hubshluft Linux\n" +
			"help\tDisplay extra information about package manager\n\n" +
			"Full documentation <https://github.com/hubshluft/hpm>\n" +
			"Source code <https://github.com/hubshluft/hpm>\n")
		return
	}

	if GetFlag != "" {
		Get(GetFlag)
		return
	}

	if UngetFlag != "" {
		Unget(UngetFlag)
		return
	}

	if UpdateFlag == "@world" {
		UpdateSystem()
		return
	} else if UpdateFlag != "" {
		UpdatePackage(UpdateFlag)
		return
	}

	if FindFlag != "" {
		Find(FindFlag)
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
