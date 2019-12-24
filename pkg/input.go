package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//function to take a message, and gather input from the user based on the message
func GetInput(message string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	return text
}