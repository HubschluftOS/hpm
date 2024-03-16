package cmd

import (
	"fmt"
	config "hpm/settings"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var URL string

func InstallPackage(packageName string) {
	URL = fmt.Sprintf(config.PackageLink, os.Args[2])
	fmt.Println(URL)
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer response.Body.Close()

	tempFile, err := os.CreateTemp("", os.Args[2])
	if err != nil {
		fmt.Println("Error creating temporar  file:", err)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, response.Body)
	if err != nil {
		fmt.Println("Error copying content to temporary file:", err)
		return
	}

	command, err := exec.Command("bash", tempFile.Name()).CombinedOutput()

	if err != nil {
		fmt.Println("Error executing the command:", err)
		return
	}

	fmt.Println(strings.TrimSpace(string(command)))
}
