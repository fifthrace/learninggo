package main

import (
	"fmt"
	"strconv"
	"strings"
)

//this is the object that the change calculator function will return.
type change struct {
	hundred int
	fifty   int
	twenty  int
	ten     int
	five    int
	single  int
	half    int
	quarter int
	dime    int
	nickel  int
	penny   int
}

func main() {
	p := fmt.Println
	p("Welcome to the Change Return Application")

	//t, _ := strconv.ParseFloat(in.GetInput("Enter the total amount charged? "), 64)
	//g, _ := strconv.ParseFloat(in.GetInput("Enter the total money given? "), 64)
	t := 123.00
	g := 500.26
	b := false

	p("Change due: ", g-t)
	ch := calculateChange(g-t, b)

	fmt.Printf("%+v",ch)
}

//calculateChange given a total price and given money will tell a cashier what bills and coins to return.
func calculateChange(m float64, useHalfDollar bool) change {
	var c change

	//let's change the float to a string, then we'll split it on the decimal
	s := fmt.Sprintf("%f", m)
	sl := strings.Split(s, ".")

	if len(sl) != 2 {
		panic("Expected a single decimal.")
	}

	dollars, _ := strconv.Atoi(sl[0])
	cents, _ := strconv.Atoi(sl[1][:2])

	getBills(dollars, &c)
	getCoins(cents, &c, useHalfDollar)

	return c
}

func getBills(d int, c *change) {
	if d > 100 {
		c.hundred = d/100
		d = d - (c.hundred * 100)
	}

	if d > 50 {
		c.fifty = d/50
		d = d - (c.fifty * 50)
	}

	if d > 20 {
		c.twenty = d/20
		d = d - (c.twenty * 20)
	}

	if d > 10 {
		c.ten = d/10
		d = d - (c.ten * 10)
	}

	if d > 5 {
		c.five = d/5
		d = d - (c.five * 5)
	}

	c.single = d
}

func getCoins(i int, c *change, useHalfDollar bool) {

	if useHalfDollar && i > 50 {
		c.half = 1
		i = i - 50
	}

	if i > 25 {
		c.quarter = 1
		i = i - 25
	}

	if i > 10 {
		c.dime = i/10
		i = i - (c.dime * 10)
	}

	if i > 5 {
		c.nickel = 1
		i = i - 5
	}

	c.penny = i
}