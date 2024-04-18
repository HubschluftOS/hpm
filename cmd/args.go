package cmd

import (
	"flag"
	"fmt"
)

func Args() {
	var flagInput string
	flag.StringVar(&flagInput, "sync", "", "sync the package with your system")
	flag.Parse()

	if flagInput == "" {
		nonFlagArgument := flag.Args()
		if len(nonFlagArgument) > 0 {
			firstFlag := nonFlagArgument[0]
			fmt.Printf("Invalid flag '%s'\n", firstFlag)
		}
		return
	}

	if flagInput != "" {
		PkgInformation(flagInput)
	}

	// Multi-packages
	// nonFlagsArgs := flag.Args()
	// if len(nonFlagsArgs) == 0 {
	// 	return
	// } else {
	// 	for _, arg := range nonFlagsArgs {
	// 		fmt.Println(strings.ReplaceAll(arg, "\n", " "))
	// 	}
	// }
}
