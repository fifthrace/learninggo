package main

import (
	"strconv"
	"testing"
)

func TestGetPi( t *testing.T) {
	got := getPi(10)
	if got != "3.141592653" {
		t.Errorf("getPi(10) = %s; should be 3.141592653", got)
	}

	got = getPi(5)
	if got != "3.1415" {
		t.Errorf("getPi(5) = %s; should be 3.1415", got)
	}

	got = getPi(105)
	gotlen := strconv.Itoa(len(got))
	if gotlen != "101" {
		t.Errorf("getPi(105) length of string was %s; should have limited to 100",gotlen)
	}
}
