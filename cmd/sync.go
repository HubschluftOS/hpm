package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const ContinueMSG = `
Packages:	  %s
Version:      %s
Maintainer:   %s
Dependencies: %s

Continue? [Y/n] `

var Input string

func Sync(pkg string, version string, maintainer string, dependencies []interface{}, source string) {
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
			fmt.Printf("[0/1] failed to get URL: %s\n", err)
			return
		}
		defer resp.Body.Close()

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

		fmt.Println(URLFileName)

		if err := UntarGzFile(URLFileName); err != nil {
			fmt.Printf("[0/1] Failed to extract file: %s\n", err)
			return
		}

		fmt.Printf("[1/1] %s successfully installed\n", pkg)
	} else {
		fmt.Printf("[1/1] exiting \n")
		return
	}
}
