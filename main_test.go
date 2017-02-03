package main

import (
	"runtime"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Log(runtime.Version())
}

func TestDummy(t *testing.T) {
	if dummy(37) != 42 {
		t.Fail()
	}
}
