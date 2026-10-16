package ohsnap_test

import (
	"math/rand/v2"
	"testing"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/arbitrary"
)

func TestWeighted_Generate_Distribution(t *testing.T) {
	rnd := rand.New(rand.NewPCG(42, 54))

	// Two arbitraries: 1 weight for "a", 3 weights for "b"
	type letter string

	arb := ohsnap.Weighted[letter](rnd, map[int]ohsnap.Arbitrary[letter]{
		1: ohsnap.OneOfValue(rnd, letter("a")),
		3: ohsnap.OneOfValue(rnd, letter("b")),
	})

	counts := map[letter]int{
		"a": 0,
		"b": 0,
	}

	const samples = 4000
	for range samples {
		val := arb.Generate()
		counts[val]++
	}

	// "b" should appear about 3x as often as "a"
	ratio := float64(counts["b"]) / float64(counts["a"])
	if ratio < 2.5 || ratio > 3.5 {
		t.Errorf("Expected ratio of b to a to be about 3:1, got %.2f (counts: %+v)", ratio, counts)
	}
}

func TestWeighted_ZeroOrNegativeWeightsAreIgnored(t *testing.T) {
	rnd := rand.New(rand.NewPCG(1, 2))

	arb := ohsnap.Weighted[int](rnd, map[int]ohsnap.Arbitrary[int]{
		0:  ohsnap.OneOfValue(rnd, 1),
		-2: ohsnap.OneOfValue(rnd, 2),
		5:  ohsnap.OneOfValue(rnd, 42),
	})

	// Only "42" should ever be generated
	for range 100 {
		val := arb.Generate()
		if val != 42 {
			t.Errorf("Expected only 42, got %v", val)
		}
	}
}

func TestWeighted_EmptyOrAllZeroWeightsReturnsZeroValue(t *testing.T) {
	rnd := rand.New(rand.NewPCG(3, 4))

	arb := ohsnap.Weighted[int](rnd, map[int]ohsnap.Arbitrary[int]{})

	val := arb.Generate()
	if val != 0 {
		t.Errorf("Expected zero value for empty input, got %v", val)
	}

	arb2 := ohsnap.Weighted[int](rnd, map[int]ohsnap.Arbitrary[int]{0: ohsnap.OneOfValue(rnd, 99)})

	val2 := arb2.Generate()
	if val2 != 0 {
		t.Errorf("Expected zero value for all zero weights, got %v", val2)
	}
}

func TestWeightedDelegatesShrinkToLastPick(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(6, 6))

	small := ohsnap.OneOfValue(rnd, 7) // only passing values
	large := arbitrary.Int(rnd, 0, 100000)
	arb := ohsnap.Weighted[int](rnd, map[int]ohsnap.Arbitrary[int]{
		1: small,
		3: large,
	})

	value, shrunk := ohsnap.FindSimplestBadCase(
		1000,
		arb,
		func(x int) bool { return x == 7 },
		ohsnap.CheckOptions{LogShrinkCandidates: true, Logger: t},
	)
	if value == nil {
		t.Fatal("expected property to fail")
	}

	if *shrunk != 0 {
		t.Errorf("shrunk = %d, want 0 (delegated halving of the large generator)", *shrunk)
	}
}
