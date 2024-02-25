package cmd

import (
	"fmt"
	"log"
	"os"

	config "github.com/rendick/pem/settings"
)

var (
	pem = `		
    dMMMMb  dMMMMMP dMMMMMMMMb
   dMP.dMP dMP     dMP"dMP"dMP	%s
  dMMMMP" dMMMP   dMP dMP dMP	%s
 dMP     dMP     dMP dMP dMP	%s
dMP     dMMMMMP dMP dMP dMP

`
)

var Packages_output string

func PackageStats() {
	stats, err := os.ReadDir(config.PackageDir)
	if err != nil {
		log.Fatal(err)
	}
	Packages_output = fmt.Sprintf("%d", len(stats)-1)

	if len(stats) > 0 {
		// Convert the integer to a string before concatenating
		fmt.Printf(pem,
			config.Yellow+config.Bold+"PACKAGES:  "+config.Reset+Packages_output,
			config.Yellow+config.Bold+"VERSION:   "+config.Reset+config.Version,
			config.Yellow+config.Bold+"HLOS:      "+config.Reset+"SOON")
	}
}
