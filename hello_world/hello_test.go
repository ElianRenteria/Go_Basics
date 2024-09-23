package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello(Salutation{name:"Elian"})
		want := "Hello, Elian"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello(Salutation{})
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello(Salutation{name:"Elian", language:"Spanish"})
		want := "Hola, Elian"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello(Salutation{name:"Elian", language:"French"})
		want := "Bonjour, Elian"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}