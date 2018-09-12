package main

import (
	"testing"
)

func Add(x, y int) (z int) {
	z = x + y
	return
}

func TestAdd(t *testing.T) {
	if Add(1, 2) == 3 {
		t.Errorf("1+2=3777")
	}

	if Add(1, 1) == 3 {
		t.Errorf("1+1=35555")
	}
}
