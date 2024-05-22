package cmd

import (
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

func Sync(pkg string, version string, maintainer string, dependencies []interface{}, source string, path string) {
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
			fmt.Scan(&Input)
			Input = strings.TrimSpace(strings.ToLower(Input))

			inputSlice := []string{"y", "Y", "yes", "YES", "ye", "YE"}
			inputTypes := false
			for _, str := range inputSlice {
				if str == Input {
					inputTypes = true
					break
				}
			}

			if !inputTypes {
				fmt.Printf("[1/1] exiting\n")
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
					} else {
						fmt.Printf("[1/1] deleting %s\n", URLFileName)
						if err := os.Remove(URLFileName); err != nil {
							fmt.Printf("[0/1] Failed to delete %s\n", URLFileName)
							return
						} else {

							sourcePath := path + pkg
							if err := os.Rename("/home/rendick/programming/hpm/neofetch-7.1.0/neofetch", sourcePath); err != nil {
								fmt.Print(err)
								return
							}

							if err := os.Chmod(sourcePath, 0755); err != nil {
								fmt.Println(err)
								return
							} else {
								fmt.Printf("[1/1] %s successfully installed\n", pkg)
							}
						}
					}
				}
			}
		}
	}
}
