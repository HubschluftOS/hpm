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

const ContinueMSG = `
Package name: %s
Version:      %s
Maintainer:   %s
Dependencies: %s

Continue? [Y/n] `

var Input string

func PkgInformation(pkg string) {
	var packageData map[string]interface{}
	pkgULR := "https://hubschluftos.github.io/db/packages/" + strings.TrimSpace(pkg) + ".json"

	fmt.Printf("[1/1] connecting to a given URL\n")
	client := &http.Client{}
	req, err := http.NewRequest("GET", pkgULR, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("[1/1] creating a URL request\n")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	var pkgData []byte
	fmt.Printf("[1/1] reading the package information\n")
	pkgData, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(pkgData, &packageData)
	if err != nil {
		fmt.Println(err)
		return
	}

	version, ok := packageData["version"].(string)
	if !ok {
		fmt.Printf("Error while parsing JSON file (version)\n")
		return
	}

	maintainer, ok := packageData["maintainer"].(string)
	if !ok {
		fmt.Printf("Error while parsing JSON file (maintainer)\n")
	}

	dependencies, ok := packageData["dependencies"].([]interface{})
	if !ok {
		fmt.Printf("Error while parsing JSON file (dependencies)\n")
		return
	}

	source, ok := packageData["source"].(string)
	if !ok {
		fmt.Printf("Error while parsing JSON file (source)\n")
	}

	Sync(pkg, version, maintainer, dependencies, source)
}

func Sync(pkg string, version string, maintainer string, dependencies []interface{}, source string) {
	fmt.Println("Sync function called with:", pkg)

	URLFileName := pkg + "-" + version + ".tar.gz"
	fmt.Printf(ContinueMSG, pkg, version, maintainer, dependencies)
	fmt.Scan(&Input)
	Input = strings.TrimSpace(strings.ToLower(Input))

	input_slice := []string{"y", "Y", "yes", "YES", "ye", "YE"}
	input_types := false
	for _, str := range input_slice {
		if str == strings.TrimSpace(Input) {
			input_types = true
			break
		}
	}

	if input_types == true {
		fmt.Printf("[1/1] getting the URL\n")
		resp, err := http.Get(source)
		if err != nil {
			fmt.Printf("[0/1] failed to get URL %s\n", err)
			return
		}
		defer resp.Body.Close()

		fmt.Printf("[1/1] creating a file\n")
		out, err := os.Create(URLFileName)
		if err != nil {
			fmt.Printf("[0/1] failed to create file %s\n", err)
			return
		}
		defer out.Close()

		fmt.Printf("[1/1] copying a file\n")
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Printf("failed to copy file %s\n", err)
			return
		}

		fmt.Printf("[1/1] %s successfully installed\n", pkg)
	} else {
		fmt.Printf("[1/1] exiting \n")
		return
	}
}
