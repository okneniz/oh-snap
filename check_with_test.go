package ohsnap

import (
	"math/rand/v2"
	"testing"
)

func TestCheckWithShrinkBudget(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(11, 11))
	arb := ArbitraryInt(rnd, 1000, 2000)
	prop := func(int) bool { return false } // every candidate fails

	value, budgeted := findSimplestBadCaseWith(1, arb, prop, ShrinkOptions{Budget: 2})
	if value == nil {
		t.Fatal("expected property to fail")
	}

	// two evaluated candidates: first halving, then halving again
	want := *value / 4
	if *budgeted != want {
		t.Errorf("budgeted shrunk = %d, want %d", *budgeted, want)
	}

	rnd2 := rand.New(rand.NewPCG(11, 11))
	arb2 := ArbitraryInt(rnd2, 1000, 2000)
	_, unlimited := findSimplestBadCaseWith(1, arb2, prop, ShrinkOptions{})
	if *unlimited != 0 {
		t.Errorf("unlimited shrunk = %d, want 0", *unlimited)
	}
}

func TestCheckWithPassingProperty(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(12, 12))
	arb := ArbitraryInt(rnd, 0, 100)

	CheckWith(t, 100, arb, func(x int) bool { return x >= 0 }, ShrinkOptions{Budget: 10})
}
