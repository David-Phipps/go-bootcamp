package main

import "testing"

// has to be a file _test.go and same package
// best pracice to capitalize
func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected ", 5, " Got ", x)
	}
}
