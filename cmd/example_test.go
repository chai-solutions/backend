package main

import (
	"testing"
)

// Remove this file once real testing is implemented.
func TestOnePlusOne(t *testing.T) {
	actual := 1 + 1
	expected := 2

	if expected != actual {
		t.Errorf("expected %d, received %d", expected, actual)
	}
}
