package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		some_name := "Chris"
		got := Hello(some_name)
		want := "Hello "+some_name

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		the_default_value := "world"
		got := Hello("")
		want := "Hello "+ the_default_value

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}
