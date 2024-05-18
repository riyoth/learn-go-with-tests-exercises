package interactions_test

import (
	"testing"

	"github.com/quii/go-specs-greet/domain/interactions"
	"github.com/quii/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))
}
func TestGreetEmptyString(t *testing.T) {
	got := interactions.Greet("")
	want := "Hello, World"

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
