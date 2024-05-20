package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetPackageInformation(pkg string) {
	var packageData map[string]interface{}
	pkgURL := "https://hubschluft.github.io/db/" + strings.TrimSpace(pkg) + ".json"

	fmt.Printf("[1/1] connecting to a given URL\n")
	client := &http.Client{}
	req, err := http.NewRequest("GET", pkgURL, nil)
	if err != nil {
		fmt.Printf("[0/1] error while connecting to a given URL: %s\n", err)
		return
	}

	fmt.Printf("[1/1] creating a URL request\n")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[0/1] error while creating a URL request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyText))

	err = json.Unmarshal(bodyText, &packageData)
	if err != nil {
		fmt.Printf("[0/1] error while unmarshaling a JSON file: %s\n", err)
		return
	}

	version, ok := packageData["version"].(string)
	if !ok {
		fmt.Printf("[0/1] error while parsing JSON file (version)\n")
		return
	}

	maintainer, ok := packageData["maintainer"].(string)
	if !ok {
		fmt.Printf("[0/1] error while parsing JSON file (maintainer)\n")
	}

	dependencies, ok := packageData["dependencies"].([]interface{})
	if !ok {
		fmt.Printf("[0/1] Error while parsing JSON file (dependencies)\n")
		return
	}

	source, ok := packageData["source"].(string)
	if !ok {
		fmt.Printf("[0/1] error while parsing JSON file (source)\n")
	}

	path, ok := packageData["path"].(string)
	if !ok {
		fmt.Printf("[0/1] error while parsing JSON file (source)\n")
	}
	fmt.Println(source)

	configPath := "/home/rendick" + "/.config/hpm/" + pkg + ".json"
	fmt.Println(configPath)

	packageJsonInformation, err := os.Create(configPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(packageJsonInformation)

	if err := os.WriteFile(configPath, []byte(bodyText), 0755); err != nil {
		fmt.Println(err)
	}

	Sync(pkg, version, maintainer, dependencies, source, path)
	return
}
