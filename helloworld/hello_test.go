package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Salvador", "")
        want := "Hello, Salvador!"
        assertCorrectMessage(t, got, want)
    })

    t.Run("saying 'Hello, world' when an empty string is supplied", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, world!"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Salvador", "Spanish")
        want := "Hola, Salvador!"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in French", func(t *testing.T) {
        got := Hello("Salvador", "French")
        want := "Bonjour, Salvador!"
        assertCorrectMessage(t, got, want)
    })

}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
