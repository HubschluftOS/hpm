package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/mmcdole/gofeed/rss"
)

func News() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", db+"news.xml", nil)
	if err != nil {
		fmt.Printf("[1/1] error while connecting to a given URL: %s\n", err)
		return
	} else {
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			fp := rss.Parser{}
			rssFeed, _ := fp.Parse(strings.NewReader(string(bodyText)))
			fmt.Println(rssFeed.WebMaster)
		}
	}
}
