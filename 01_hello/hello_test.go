package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		some_name := "Chris"
		got := Hello(some_name, "")
		want := "Hello " + some_name

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		the_default_value := "world"
		got := Hello("", "")
		want := "Hello " + the_default_value

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		some_name := "Elodie"
		got := Hello(some_name, "Spanish")
		want := "Hola " + some_name

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		some_name := "Elodie"
		got := Hello(some_name, "French")
		want := "Bonjour " + some_name

		assertCorrectMessage(t, got, want)
	})
	t.Run("explicit in english", func(t *testing.T){
		some_name := "Elodie"
		got := Hello(some_name, "English")
		want := "Hello " + some_name

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
