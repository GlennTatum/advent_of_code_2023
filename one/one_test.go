package main

import (
	"testing"
)

const (
	x = "ninefivefive2nine5ntvscdfdsmvqgcbxxxt"
	y = "four4eightlnhtvrscbf5gh5"
	z = "jgsrnzzz8cf4eighteight3"
)

const (
	x1 = 25
	y1 = 45
	z1 = 83
)

func TestDecode(t *testing.T) {
	v, err := DecodeCalibrationValue(x)
	if err != nil {
		t.Fatal()
	}
	if v == x1 {
		t.Log(x, v, x1)
		t.Skipped()
	}
}
