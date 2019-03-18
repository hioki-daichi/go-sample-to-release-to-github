package main

import "testing"

func TestFoo(t *testing.T) {
	expected := "Hello, statik!!\n"

	actual := foo()
	if actual != expected {
		t.Errorf("expected: %q, actual: %q", expected, actual)
	}
}
