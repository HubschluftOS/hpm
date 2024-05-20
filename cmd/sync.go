package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const ContinueMSG = `
Packages:		%s
Version:		%s
Maintainer:		%s
Dependencies:		%s
Source:			%s

Continue? [Y/n] `

var (
	Input        string
	PackageInput string
)

func Sync(pkg string, version string, maintainer string, dependencies []interface{}, source string, path string) {
	ConfigurateManager()
	PackageInput = fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), pkg)
	Logs()
	URLFileName := "/tmp/" + pkg + "-" + version + ".tar.gz"
	fmt.Printf(ContinueMSG, pkg, version, maintainer, dependencies, source)
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
			fmt.Printf("[0/1] failed to get URL: %s\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("[0/1] failed to download file: %s\n", resp.Status)
		}

		fmt.Printf("[1/1] creating a file\n")
		out, err := os.Create(URLFileName)
		if err != nil {
			fmt.Printf("[0/1] failed to create file: %s\n", err)
			return
		}
		defer out.Close()

		fmt.Printf("[1/1] copying a file\n")
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Printf("[0/1] failed to copy file: %s\n", err)
			return
		}

		if err := UntarGzFile(URLFileName); err != nil {
			fmt.Printf("[0/1] Failed to extract file: %s\n", err)
			return
		}

		fmt.Printf("[1/1] deleting a %s\n", URLFileName)
		if removeTar := os.Remove(URLFileName); removeTar != nil {
			fmt.Printf("[0/1] Failed to delete %s\n", URLFileName)
			return
		}

		sourcePath := path + pkg
		if err := os.Rename("/home/rendick/programming/hpm/neofetch-7.1.0/neofetch", sourcePath); err != nil {
			fmt.Print(err)
			return
		}

		if err := os.Chmod(sourcePath, 0755); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("[1/1] %s successfully installed\n", pkg)
		return
	} else {
		fmt.Printf("[1/1] exiting \n")
		return
	}
}
