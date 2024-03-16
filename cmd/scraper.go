package cmd

import (
	"fmt"
	config "hpm/settings"
	"log"
	"strings"

	"github.com/gocolly/colly"
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
		log.Printf(config.Bold+"%s\n"+config.Reset, r.URL)
	})

	fmt.Printf("%s\n\n", c.Visit(config.PackagesURL))
}
