package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"hpm/modules"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	PackageJsonOutput []byte

	PackageData    map[string]interface{}
	name           string
	description    string
	version        string
	maintainer     string
	installation   []string
	uninstallation []string

	err error
	ok  bool
)

func Curl(pkg string) bool {
	modules.Info("Getting the URL.")
	client := &http.Client{}
	repo := Repository + pkg + ".json"
	req, err := http.NewRequest("GET", repo, nil)
	if err != nil {
		modules.Error("Error while connecting to a given URL: %s", err)
		return false
	} else {
		modules.Info("Making a request.")
		resp, err := client.Do(req)
		if err != nil {
			modules.Error("Error while creating a URL request: %s", err)
			return false
		}
		defer resp.Body.Close()

		modules.Info("Checking the status code.")
		if resp.StatusCode != http.StatusOK {
			modules.Error("Bad status code: %s not found", pkg)
			return false
		}

		modules.Info("Reading a JSON file.")
		PackageJsonOutput, err = io.ReadAll(resp.Body)
		if err != nil {
			modules.Error("Error while reading a JSON file: %s", err)
			return false
		}
	}

	return true
}

func UnmarshalPackage() {
	modules.Info("Unmarshalling a JSON file.")
	err := json.Unmarshal(PackageJsonOutput, &PackageData)
	if err != nil {
		modules.Error("Error unmarshalling JSON: %s", err)
		return
	} else {
		name, ok = PackageData["name"].(string)
		if !ok {
			modules.Error("Error while parsing JSON file.")
			return
		}

		description, ok = PackageData["description"].(string)
		if !ok {
			modules.Error("Error while parsing JSON file.")
			return
		}

		version, ok = PackageData["version"].(string)
		if !ok {
			modules.Error("Error while parsing JSON file.")
			return
		}

		maintainer, ok = PackageData["maintainer"].(string)
		if !ok {
			modules.Error("Error while parsing JSON file.")
			return
		}

		installationInterface, ok := PackageData["installation"].([]interface{})
		if !ok {
			modules.Error("Error while parsing JSON file.")
			return
		}
		for _, cmd := range installationInterface {
			installation = append(installation, cmd.(string))
		}

		uninstallationInterface, ok := PackageData["uninstallation"].([]interface{})
		if !ok {
			modules.Error("Error while parsing JSON file.")
			return
		}
		for _, cmd := range uninstallationInterface {
			uninstallation = append(uninstallation, cmd.(string))
		}
	}
}

func ExecuteShell(commands []string) {
	for _, command := range commands {
		cmd := exec.Command("sh", "-c", command)
		output, err := cmd.CombinedOutput()
		if err != nil {
			modules.Error("Error while executing the command: %s", err)
			fmt.Printf("%s\n", command)
			fmt.Printf("%s\n", string(output))
		} else {
			fmt.Printf("%s\n", command)
			fmt.Printf("%s\n", strings.TrimSpace(string(output)))
		}
	}
	modules.Success("%s successfully installed", name)
	return
}

func Get(pkg string) {
	if modules.IsSudo() == true {
		if _, err := os.Stat("/hl-bin/" + strings.TrimSpace(strings.ToLower(string(pkg)))); err == nil {
			modules.Error("File exists. Aborting.")
		} else if errors.Is(err, os.ErrNotExist) {
			if Curl(pkg) {
				UnmarshalPackage()

				getIndex := -1
				for i, arg := range os.Args {
					if arg == "-get" || arg == "--get" {
						getIndex = i
						break
					}
				}

				if getIndex != -1 {
					getArgsCount := 0
					for i := getIndex + 1; i < len(os.Args); i++ {
						if os.Args[i][0] == '-' {
							break
						}
						getArgsCount++
					}

					fmt.Printf(PackageInformation,
						modules.Bold, modules.Reset, name, getArgsCount,
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
							return
						} else {
							modules.Error("Exiting.")
							return
						}
					}
				}
			} else {
				modules.Error("Failed to fetch package information. Aborting.")
				return
			}
		}
	} else {
		modules.Error("Unable to get current user. Aborting.")
		return
	}
}
