package main

import (
	"fmt"
	"../pkg/rbinput"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the Mortgage Calculator Application\n")

	r, _ := strconv.ParseFloat(rbinput.GetInput("What is your interest rate?"), 64)
	t, _ := strconv.ParseFloat(rbinput.GetInput("How many months is the term?"), 64)

	//compounded daily by default
	m := getMonthlyPayment(r,t)

	fmt.Println("Your monthly payment is: " + fmt.Sprintf("%f", m))
}

func getMonthlyPayment(rate float64, term float64) float64{


	return 100
}