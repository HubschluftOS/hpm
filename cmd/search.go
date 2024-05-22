package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	packageInformationSerach map[string]interface{}
	showJson                 string
)

func Search(packageSearch string) {
	fmt.Printf("[1/1] connecting to a given URL\n")
	client := &http.Client{}
	req, err := http.NewRequest("GET", db+strings.TrimSpace(string(packageSearch))+".json", nil)
	if err != nil {
		fmt.Printf("[0/1] failed to get URL: %s\n", err)
		return
	} else {
		fmt.Printf("[1/1] creating a URL request\n")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("[0/1] error while creating a URL request: %s\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("[0/1] failed to download a file: %s\n", resp.Status)
			return
		} else {
			bodyText, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("[0/1] error while displaying a JSON file: %s\n", err)
				return
			}

			if err := json.Unmarshal(bodyText, &packageInformationSerach); err != nil {
				fmt.Printf("[0/1] error while unmarshaling a JSON file: %s\n", err)
				return
			} else {
				description, ok := packageInformationSerach["description"].(string)
				if !ok {
					fmt.Println(ok)
					return
				} else {
					fmt.Printf(packageInfo, packageSearch, red, bold, description, reset)
				}
			}
		}
	}
}
