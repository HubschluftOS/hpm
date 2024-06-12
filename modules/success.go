package modules

import "fmt"

func Success(success_text string) string {
	return fmt.Sprintf("%s%s-> Done: %s%s\n", Green, Bold, success_text, Reset)
}
