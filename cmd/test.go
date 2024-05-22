package cmd

import (
	"fmt"
	"io"
	"net/http"
)

func Test() {
	url := "https://github.com/dylanaraps/neofetch/archive/refs/tags/7.1.0.tar.gz"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	} else {

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Failed to download file: %s\n", resp.Status)
			return
		} else {
			size, err := io.Copy(io.Discard, resp.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				return
			} else {
				fmt.Printf("size_download=%d\n", size)
			}
		}
	}

}
