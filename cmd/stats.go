package cmd

import (
	"fmt"
	config "hpm/settings"
	"log"
	"os"
)

var (
	hpm = `		
    dMP dMP dMMMMb  dMMMMMMMMb
   dMP dMP dMP.dMP dMP"dMP"dMP	%s
  dMMMMMP dMMMMP" dMP dMP dMP	%s 
 dMP dMP dMP     dMP dMP dMP	%s 
dMP dMP dMP     dMP dMP dMP   

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
		fmt.Printf(hpm,
			config.Yellow+config.Bold+"PACKAGES:  "+config.Reset+Packages_output,
			config.Yellow+config.Bold+"VERSION:   "+config.Reset+config.Version,
			config.Yellow+config.Bold+"HLOS:      "+config.Reset+"SOON")
	}
}
