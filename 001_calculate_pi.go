package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How many digits of Pi do you want? ")
	text, _ := reader.ReadString('\n')

	digits, _ := strconv.ParseInt(text, 10, 64)
	getPi(digits)

}

func getPi(digits int64) string {
	digits += 1

	fmt.Println(digits)

	return "stuff"
}
