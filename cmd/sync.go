package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	Input        string
	PackageInput string
)

func Sync(pkg, version, maintainer string, dependencies []interface{}, source, path string) {
	fmt.Printf("[1/1] getting the URL\n")
	client := &http.Client{}
	req, err := http.NewRequest("GET", source, nil)
	if err != nil {
		fmt.Printf("[0/1] error while connecting to a given URL: %s\n", err)
		return
	} else {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("[0/1] error while creating a URL request: %s\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("[0/1] bad status: %s\n", resp.Status)
			return
		} else {
			size, err := io.Copy(io.Discard, resp.Body)
			if err != nil {
				fmt.Printf("[0/1] failed to read file size: %s\n", err)
				return
			}

			fmt.Printf(ContinueMSG, pkg, version, maintainer, dependencies, size/1024, source)
			fmt.Print("Continue? [Y/n] ")

			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			input = strings.TrimSpace(strings.ToLower(input))

			fmt.Printf(red + bold + "[!/!] wait\n" + reset)

			if input != "" && input != "y" && input != "yes" {
				fmt.Printf("[0/1] exiting\n")
				return
			} else {
				resp, err = client.Do(req)
				if err != nil {
					fmt.Printf("[0/1] error while creating a URL request: %s\n", err)
					return
				}
				defer resp.Body.Close()

				fmt.Printf("[1/1] creating a file\n")
				URLFileName := "/tmp/" + pkg + "-" + version + ".tar.gz"
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
				} else {
					if err := UntarGzFile(URLFileName); err != nil {
						fmt.Printf("[0/1] Failed to extract file: %s\n", err)
						return
					}

					fmt.Printf("[1/1] deleting %s\n", URLFileName)
					if err := os.Remove(URLFileName); err != nil {
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

					PackageInput = fmt.Sprintf("%s%ssync:%s %s (%s)\n", red, bold, reset, pkg, currentTime)
					Logs()
				}
			}
		}
	}
}
