package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world 2023"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
