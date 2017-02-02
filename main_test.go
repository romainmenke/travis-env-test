package main

import "testing"

func TestDummy(t *testing.T) {
	if dummy(37) != 42 {
		t.Fail()
	}
}
