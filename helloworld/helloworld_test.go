package helloworld

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Mary", "")
		want := "Hello, Mary"

		assertMessage(t, got, want)
	})

	t.Run("Add 'World' when no name provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertMessage(t, got, want)
	})

	t.Run("Spanish greeting", func(t *testing.T) {
		got := Hello("Maria", "Spanish")
		want := "Hola, Maria"

		assertMessage(t, got, want)
	})

	t.Run("French greeting", func(t *testing.T) {
		got := Hello("Marie", "French")
		want := "Bonjour, Marie"

		assertMessage(t, got, want)
	})

}
