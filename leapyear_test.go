package main

import (
	"testing"
)

func TestLeapYear01(t *testing.T) {
	expected := leapyear(1)
	actual := true
	if expected != actual {
		t.Errorf("Error")
	}
}
