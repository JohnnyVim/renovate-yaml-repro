package main

import "testing"

func TestSomething(t *testing.T) {
	if 1 == 2 {
		t.Fatal("this cannot happen")
	}
}
