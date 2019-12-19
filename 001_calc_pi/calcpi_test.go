package main

import "testing"

func TestGetPi( t *testing.T) {
	got := getPi(10)
	if got != "3.14159265" {
		t.Errorf("getPi(10) = %s; should be 3.141592653", got)
	}
}
