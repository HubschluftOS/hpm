package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	showJson string
)

func Search(packageSearch string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", db+strings.TrimSpace(string(packageSearch))+".json", nil)
	if err != nil {
		fmt.Printf("[0/1] failed to get URL: %s\n", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("[0/1] failed to download file: %s\n", resp.Status)
		return
	} else {
		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var packageInfoRmationSerach map[string]interface{}

		if err := json.Unmarshal(bodyText, &packageInfoRmationSerach); err != nil {
			fmt.Println(err)
		}

		description, ok := packageInfoRmationSerach["description"].(string)
		if !ok {
			fmt.Println(ok)
			return
		}

		fmt.Printf(packageInfo, packageSearch, red, bold, description, reset)
	}
}
