package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	price, _ := strconv.ParseFloat(getInput("Enter cost per square unit: "), 64)

	h, _ := strconv.ParseFloat(getInput("Enter height in units: "), 64)

	w, _ := strconv.ParseFloat(getInput("Enter width in units: "), 64)

	totalCost, _ := calcCost(price, h, w)

	fmt.Println("Total Cost: " + fmt.Sprintf("%f", totalCost))
}

func calcCost(price float64, height float64, width float64) (float64, error) {
	total := height*width*price

	return total, nil
}

//function to take a message, and gather input from the user based on the message
func getInput(message string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	return text
}