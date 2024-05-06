package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func Test() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://hubschluftos.github.io/db/packages/neofetch.json", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", strings.TrimSpace(string(bodyText)))
}
