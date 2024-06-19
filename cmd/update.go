package cmd

import (
	"bufio"
	"fmt"
	"hpm/modules"
	"os"
	"strings"
)

func UpdateSystem() {
	fmt.Printf("update system\n")
}

func UpdatePackage(pkg string) {
	if modules.IsSudo() == true {
		if CheckIfFileExist(pkg) == true {
			if Curl(pkg) {
				UnmarshalPackage()

				fmt.Printf(PackageInformationUpdate,
					modules.Bold, modules.Reset, name,
					modules.Bold, modules.Reset, description,
					modules.Bold, modules.Reset, version,
					modules.Bold, modules.Reset, maintainer,
					modules.Bold, modules.Reset, installation)
				fmt.Print(modules.Bold + "Continue? [Y/n] \n" + modules.Reset)

				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					modules.Error("Error reading input: %s", err)
					return
				} else {
					input = strings.TrimSpace(strings.ToLower(string(input)))

					input_slice := []string{"", "yes", "y"}
					input_types := false

					for _, str := range input_slice {
						if str == input {
							input_types = true
							break
						}
					}

					if input_types == true {
						ExecuteShell()
					} else {
						modules.Error("Exiting.")
						return
					}
				}
			}
		} else {
			return
		}
	} else {
		modules.Error("Unable to get current user. Aborting.")
		return
	}
}
