package main

import (
	"testing"
)

func TestCal(t *testing.T) {
	if Sum(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
