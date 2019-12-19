package main

import ("testing")

func TestCalcCost( t *testing.T) {
	cost, _ := calcCost(3.5, 5, 5)
	if cost != 87.5 {
		t.Errorf("calcCost(3.5,5,5) = %.2f; should be 3.141592653", cost)
	}

}
