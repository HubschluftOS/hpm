package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	config "github.com/rendick/pem/settings"
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

	command := exec.Command("bash", tempFile.Name())

	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing the command:", err)
		return
	}

	fmt.Println(string(output))
}
