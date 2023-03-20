package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"
	if got != want {
		t.Fatalf("Wanted: %q, got: %q", want, got)
	}
}
