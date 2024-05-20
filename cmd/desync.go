package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	payload map[string]interface{}
)

func Desync(pkg string) {
	ConfigurateManager()

	readJsonConfigFileDesync, err := ioutil.ReadFile("/etc/hpm/config.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(readJsonConfigFileDesync, &payload)
	if err != nil {
		fmt.Println(err)
	}

	installationPathToString, ok := payload["installation_path"].(string)
	if !ok {
		fmt.Printf("[0/1] installation_path is not a string")
	}

	if err := os.Remove(installationPathToString + pkg); err != nil {
		log.Fatal(err)
	}
}
