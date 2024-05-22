package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	payload map[string]interface{}
)

func Desync(pkg string) {
	ConfigurateManager()

	readJsonConfigFileDesync, err := ioutil.ReadFile(configPath + "config.json")
	if err != nil {
		fmt.Printf("[0/1] error while reading a JSON file: %s\n", err)
		return
	} else {
		err = json.Unmarshal(readJsonConfigFileDesync, &payload)
		if err != nil {
			fmt.Printf("[0/1] error while unmarshaling JSON file: %s\n", err)
		} else {
			installationPathToString, ok := payload["installation_path"].(string)
			if !ok {
				fmt.Printf("[0/1] installation_path is not a string")
				return
			}

			if err := os.Remove(installationPathToString + pkg); err != nil {
				fmt.Printf("[0/1] error during package desynchronization %s\n", err)
				return
			}
		}
	}
}
