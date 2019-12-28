package main

import (
	"../pkg/rbinput"
	"fmt"
	"strconv"
)

//this is the object that the change calculator function will return.
type change struct {
	hundred int
	fifty   int
	twenty  int
	ten     int
	single  int
	half    int
	quarter int
	dime    int
	nickle  int
	penny   int
}

func main() {
	fmt.Println("Welcome to the Change Return Application\n")

	t, _ := strconv.ParseFloat(rbinput.GetInput("Enter the total amount charged? "), 64)
	g, _ := strconv.ParseFloat(rbinput.GetInput("Enter the total money given? "), 64)

	ch := calculateChange(t, g)

	fmt.Println((ch))
}

//calculateChange given a totla price and given money will tell a cashier what bills and coins to return.
func calculateChange(total float64, given float64) change {

	c := change{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	return c
}
