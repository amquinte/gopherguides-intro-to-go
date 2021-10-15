package main

import "testing"

func TestSimple(t *testing.T) {
	if simple() == true {
		t.Error("expected false, got true")
	}
}
