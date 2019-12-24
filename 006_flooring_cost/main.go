package main

import (
	"fmt"
	"strconv"
	"../pkg/rbinput"
)

func main() {

	price, _ := strconv.ParseFloat(rbinput.GetInput("Enter cost per square unit: "), 64)

	h, _ := strconv.ParseFloat(rbinput.GetInput("Enter height in units: "), 64)

	w, _ := strconv.ParseFloat(rbinput.GetInput("Enter width in units: "), 64)

	totalCost, _ := calcCost(price, h, w)

	fmt.Println("Total Cost: " + fmt.Sprintf("%f", totalCost))
}

func calcCost(price float64, height float64, width float64) (float64, error) {
	total := height*width*price

	return total, nil
}

