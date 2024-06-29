package cmd

import (
	"bufio"
	"fmt"
	"hpm/modules"
	"net/http"
	"os"
	"strings"
)

func UpdateSystem() {
	readFolder, err := os.ReadDir("/hl-bin/")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range readFolder {
		if file.IsDir() {
			continue
		}

		url := Repository + file.Name() + ".json"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", url, err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Non-OK HTTP status for %s: %s\n", url, resp.Status)
			resp.Body.Close()
			continue
		}

		UpdatePackage(file.Name())
	}
}

func UpdatePackage(pkg string) {
	if modules.IsSudo() == true {
		if Curl(pkg) {
			UnmarshalPackage()

			fmt.Printf(PackageInformationUpdate,
				modules.Bold, modules.Reset, name,
				modules.Bold, modules.Reset, description,
				modules.Bold, modules.Reset, version,
				modules.Bold, modules.Reset, maintainer,
				modules.Bold, modules.Reset, strings.Join(installation, modules.Bold+" - "+modules.Reset),
				modules.Bold, modules.Reset, strings.Join(uninstallation, modules.Bold+" - "+modules.Reset))
			fmt.Print(modules.Bold + "Continue? [Y/n] " + modules.Reset)

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
					ExecuteShell(installation)
				} else {
					modules.Error("Exiting.")
					return
				}
			}
		}
	} else {
		return
	}
}
