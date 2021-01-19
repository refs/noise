package main

import "testing"

func TestSum(t *testing.T) {
	r := sum(2, 3)
	if r != 5 {
		t.Fail()
	}
}
