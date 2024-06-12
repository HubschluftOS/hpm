package modules

import "fmt"

func Error(format string, a ...interface{}) {
	fmt.Printf("%s%s==> %s%s\n", Red, Bold, fmt.Sprintf(format, a...), Reset)
}
