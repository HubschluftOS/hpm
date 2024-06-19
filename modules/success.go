package modules

import (
	"fmt"
)

func Success(success_text string, a ...interface{}) {
	fmt.Printf("%s%s-> Done: %s%s\n", Green, Bold, fmt.Sprintf(success_text, a...), Reset)
}
