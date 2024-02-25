package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	config "github.com/rendick/pem/settings"
)

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), "\n")
}

func Scrapper() {
	c := colly.NewCollector()

	c.OnHTML("ul#myMenu", func(e *colly.HTMLElement) {
		fmt.Println(config.Red + config.Bold + StandardizeSpaces(e.Text) + config.Reset)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf(config.Red+config.Bold+config.Reset+"All packages are available on: "+config.Bold+"%s\n"+config.Reset, r.URL)
	})

	c.Visit(config.PackagesURL)
}
