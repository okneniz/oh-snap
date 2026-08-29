package ohsnap

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestOneOfDelegatesShrinkToLastPick(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(5, 5))
	small := OneOfValue(rnd, 7) // only passes values
	large := ArbitraryInt(rnd, 0, 100000)
	arb := OneOf(rnd, []Arbitrary[int]{small, large})

	value, shrunk := findSimplestBadCaseWith(1000, arb, func(x int) bool { return x == 7 }, traceShrinking(t))
	if value == nil {
		t.Fatal("expected property to fail")
	}
	if *shrunk != 0 {
		t.Errorf("shrunk = %d, want 0 (delegated halving of the large generator)", *shrunk)
	}
}

func TestOneOfShrinkBeforeGenerate(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(5, 5))
	arb := OneOf(rnd, []Arbitrary[int]{ArbitraryInt(rnd, 0, 10)})

	candidates := slices.Collect(arb.Shrink(5))
	if len(candidates) != 0 {
		t.Errorf("expected no candidates before Generate, got %v", candidates)
	}
}
