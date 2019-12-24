package main

import (
	"fmt"
	"../pkg/rbinput"
	"strconv"
	"math"
)

func main() {
	fmt.Println("Welcome to the Mortgage Calculator Application\n")

	p, _ := strconv.ParseFloat(rbinput.GetInput("What is the principal amount?"), 64)
	r, _ := strconv.ParseFloat(rbinput.GetInput("What is your annual interest rate?"), 64)
	t, _ := strconv.ParseFloat(rbinput.GetInput("How many years is the term?"), 64)
	//compounded daily by default
	m := getMonthlyPayment(p,r,t)

	fmt.Println("Your monthly payment is: " + fmt.Sprintf("%f", m))
}

// getMonthlyPayment calculates the monthly mortgage payment needed to pay off a loan in the given
// amount of years. Formula taken from https://www.bankrate.com/calculators/mortgages/mortgage-calculator.aspx
func getMonthlyPayment(principal float64, rate float64, term float64) float64{
	//take the annual rate and change it to monthly
	r := rate/12
	//number of monthly payments of the term (years * 12)
	n := term * 12
	// m = principal[r(1+r)^n/((1+r)^n)-1)]
	pow := math.Pow(1+r,n)

	m := principal * ((r * pow)/(pow-1))

	return m
}