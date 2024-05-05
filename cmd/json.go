package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func handleError(msg string, err error) {
	if err != nil {
		fmt.Printf("Error: %s: %s\n", msg, err)
	}
}

func GetPackageInformation(pkg string) {
	pkgURL := "https://hubschluftos.github.io/db/packages/" + strings.TrimSpace(pkg) + ".json"

	fmt.Println("[1/1] Connecting to the given URL")
	resp, err := http.Get(pkgURL)
	defer func() {
		handleError("closing response body", resp.Body.Close())
	}()
	handleError("creating a URL request", err)

	var packageData map[string]interface{}
	handleError("reading the package information", json.NewDecoder(resp.Body).Decode(&packageData))

	version, _ := packageData["version"].(string)
	maintainer, _ := packageData["maintainer"].(string)
	path, _ := packageData["path"].(string)

	var dependencies []string
	if dep, ok := packageData["dependencies"].([]interface{}); ok {
		for _, d := range dep {
			if s, ok := d.(string); ok {
				dependencies = append(dependencies, s)
			}
		}
	} else {
		fmt.Println("Error while parsing JSON file (dependencies)")
		return
	}

	source, _ := packageData["source"].(string)

	fmt.Println(source)

	var interfaceDeps []interface{}
	for _, dep := range dependencies {
		interfaceDeps = append(interfaceDeps, dep)
	}

	Sync(pkg, version, maintainer, interfaceDeps, source, path)
}
