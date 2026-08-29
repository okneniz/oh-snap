package ohsnap

import (
	"math/rand/v2"
	"testing"

	"github.com/okneniz/oh-snap/shrink"
)

func TestWithShrinkerStaysInRange(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(7, 7))
	prop := func(x int) bool { return x >= 150 }

	t.Run("default shrink leaves the range", func(t *testing.T) {
		t.Parallel()

		arb := ArbitraryInt(rnd, 100, 200)

		_, shrunk := findSimplestBadCase(1000, arb, prop)
		if shrunk == nil {
			t.Fatal("expected property to fail")
		}
		if *shrunk != 0 {
			t.Errorf("shrunk = %d, want 0 (default halving goes to zero)", *shrunk)
		}
	})

	t.Run("halving to lower bound stays inside", func(t *testing.T) {
		t.Parallel()

		arb := WithShrinker(ArbitraryInt(rnd, 100, 200), shrink.Halving(100))

		_, shrunk := findSimplestBadCase(1000, arb, prop)
		if shrunk == nil {
			t.Fatal("expected property to fail")
		}
		if *shrunk != 100 {
			t.Errorf("shrunk = %d, want 100 (minimal in-range failure)", *shrunk)
		}
	})
}

func TestWithShrinkerKeepsGenerate(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(9, 9))
	arb := WithShrinker(ArbitraryInt(rnd, 5, 10), shrink.NoShrink[int]())

	for i := 0; i < 100; i++ {
		value := First(arb.Generate())
		if value < 5 || value > 10 {
			t.Errorf("value = %d, want it in [5, 10]", value)
		}
	}
}
