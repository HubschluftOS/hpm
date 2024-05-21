package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Test() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Press Enter to continue, or type 'y' or 'n' and press Enter:")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("You pressed Enter. Continuing program...")
	} else if input == "y" {
		fmt.Println("You entered 'y'. Proceeding with Yes.")
	} else if input == "n" {
		fmt.Println("You entered 'n'. Proceeding with No.")
	} else {
		fmt.Println("Invalid input. Please enter 'y', 'n', or press Enter.")
	}
}
