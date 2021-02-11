package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	x := mySum(5, 8)

	if x != 13 {
		t.Error("Expected", 5, "Got", x)
	}
}