package main

import "testing"

func TestHello(t *testing.T) {
	some_name := "Chris"
	got := Hello(some_name)
	want := "Hello "+some_name

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
