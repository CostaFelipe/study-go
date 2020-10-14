package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 6)
	if total != 11 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 11)
		fmt.Println("oi")
	}
}
