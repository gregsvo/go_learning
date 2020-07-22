package main

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello("Greg")
// 	want := "Hello, Greg!"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Greg", "en")
		want := "Hello, Greg!"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "esp")
		want := "Hola, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "fr")
		want := "Bonjour, Elodie!"

		assertCorrectMessage(t, got, want)
	})
}
