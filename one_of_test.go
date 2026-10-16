package ohsnap_test

import (
	"math/rand/v2"
	"slices"
	"testing"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/arbitrary"
)

func TestOneOfDelegatesShrinkToLastPick(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(5, 5))
	small := ohsnap.OneOfValue(rnd, 7) // only passes values
	large := arbitrary.Int(rnd, 0, 100000)
	arb := ohsnap.OneOf(rnd, small, large)

	value, shrunk := ohsnap.FindSimplestBadCase(1000, arb, func(x int) bool { return x == 7 }, ohsnap.CheckOptions{LogShrinkCandidates: true, Logger: t})
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
	arb := ohsnap.OneOf(rnd, arbitrary.Int(rnd, 0, 10), arbitrary.Int(rnd, 0, 10))

	candidates := slices.Collect(arb.Shrink(5))
	if len(candidates) != 0 {
		t.Errorf("expected no candidates before Generate, got %v", candidates)
	}
}
