package config

import (
	"fmt"
	"time"
)

var (
	// style
	Red    = "\033[31m"
	Bold   = "\033[1m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
	Time   = time.Now().Format("2006-01-02 15:04:05")
)

const (
	Version     = "1.1v hpm"
	PackageLink = "https://hubschluftos.github.io/db/packages/%s.html"
	PackagesURL = "https://hubschluftos.github.io/pages/packages.html"
	PackageDir  = "/bin/"
	ErrorMsg    = "%s not successfully uninstalled! Could not find file or use"
	SuccessMsg  = "%s successfully uninstalled!\n"
)

func Help() {
	fmt.Println(Bold + "Usage: hpm [OPTION] [FILE]" + Reset +
		"\n\nsync:	\tinstallation of the entered package" +
		"\nremove: \tuninstalling of the entered package" +
		"\nlist:   \tprinting of all available packages" +
		"\ninfo:   \tprint all information about the package manager" +
		"\n\nhelp:    \tprint the help menu" +
		"\nversion:\tprint the version" +
		"\n\nhpm:    \t<https://github.com/HubschluftOS/hpm>")
}
