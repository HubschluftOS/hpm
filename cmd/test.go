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

	// Trim the input to remove leading and trailing whitespace (including the newline)
	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("You pressed Enter. Continuing program...")
		// Add logic for Enter key action
	} else if input == "y" {
		fmt.Println("You entered 'y'. Proceeding with Yes.")
		// Add logic for 'y' action
	} else if input == "n" {
		fmt.Println("You entered 'n'. Proceeding with No.")
		// Add logic for 'n' action
	} else {
		fmt.Println("Invalid input. Please enter 'y', 'n', or press Enter.")
	}
}
