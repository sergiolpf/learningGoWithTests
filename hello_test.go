package main

import "testing"

func TestHello(t *testing.T){


	assetCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"

		assetCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assetCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assetCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got := Hello("Jean Jack", "French")
		want := "Bonjour, Jean Jack!"
		assetCorrectMessage(t, got, want)
	})
}
