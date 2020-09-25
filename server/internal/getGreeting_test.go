package internal

import (
	"testing"
)

func TestGetGreeting(t *testing.T){
	var result = GetGreeting("Tom")
	var expected = "Hello, Tom"

	if result != expected {
		t.Errorf("Should have return '%s' rathern tahn '%s'\n", expected, result)
	}
}
