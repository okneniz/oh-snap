package ohsnap_test

import (
	"math/rand/v2"
	"testing"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/arbitrary"
	"github.com/okneniz/oh-snap/shrink"
)

func TestWithShrinkerStaysInRange(t *testing.T) {
	rnd := rand.New(rand.NewPCG(7, 7))
	prop := func(x int) bool { return x >= 150 }

	t.Run("default shrink leaves the range", func(t *testing.T) {
		arb := arbitrary.Int(rnd, 100, 200)

		_, shrunk := ohsnap.FindSimplestBadCase(1000, arb, prop, ohsnap.CheckOptions{LogShrinkCandidates: true, Logger: t})
		if shrunk == nil {
			t.Fatal("expected property to fail")
		}

		if *shrunk != 0 {
			t.Errorf("shrunk = %d, want 0 (default halving goes to zero)", *shrunk)
		}
	})

	t.Run("halving to lower bound stays inside", func(t *testing.T) {
		arb := ohsnap.WithShrinker(arbitrary.Int(rnd, 100, 200), shrink.Halving(100))

		_, shrunk := ohsnap.FindSimplestBadCase(1000, arb, prop, ohsnap.CheckOptions{LogShrinkCandidates: true, Logger: t})
		if shrunk == nil {
			t.Fatal("expected property to fail")
		}

		if *shrunk != 100 {
			t.Errorf("shrunk = %d, want 100 (minimal in-range failure)", *shrunk)
		}
	})
}

func TestWithShrinkerKeepsGenerate(t *testing.T) {
	rnd := rand.New(rand.NewPCG(9, 9))
	arb := ohsnap.WithShrinker(arbitrary.Int(rnd, 5, 10), shrink.NoShrink[int]())

	for range 100 {
		value := arb.Generate()
		if value < 5 || value > 10 {
			t.Errorf("value = %d, want it in [5, 10]", value)
		}
	}
}
