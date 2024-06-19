package cmd

import (
	"encoding/json"
	"fmt"
	"hpm/modules"
	"io"
	"net/http"
)

var (
	NewsData map[string]interface{}
)

func News() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://hubshluft.github.io/news.json", nil)
	if err != nil {
		modules.Error("Error while connecting to a given URL: %s", err)
		return
	} else {
		resp, err := client.Do(req)
		if err != nil {
			modules.Error("Error while creating a URL request: %s", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			modules.Error("Bad status code: %s", resp.Status)
			return
		}

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			modules.Error("Error while reading a JSON file: %s", err)
			return
		} else {
			err = json.Unmarshal(bodyText, &NewsData)
			if err != nil {
				modules.Error("Error unmarshalling JSON: %s", err)
				return
			} else {
				channel, ok := NewsData["channel"].(map[string]interface{})
				if !ok {
					modules.Error("Error while parsing JSON file.")
					return
				}

				newsArr, ok := channel["news"].([]interface{})
				if !ok {
					modules.Error("Error while parsing JSON file.")
					return
				}

				for _, newsItem := range newsArr {
					news, ok := newsItem.(string)
					if !ok {
						modules.Error("Error while parsing news item.")
						return

					} else {
						fmt.Printf("%s\n", news)
						return
					}
				}
			}
		}
	}
}
