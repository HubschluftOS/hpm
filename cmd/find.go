package cmd

import (
	"hpm/modules"
)

func Find(pkg string) {
	if Curl(pkg) == true {
		modules.Success("%s found", pkg)
		return
	}
}
