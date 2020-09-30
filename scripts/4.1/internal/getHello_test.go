package internal

import (
	"testing"
)

func TestGetHello(t *testing.T){
	var v = GetHello()

	if v != "Hello World" {
		t.Errorf("Should have returned 'Hello World' rather than '%s'", v)
	}
}