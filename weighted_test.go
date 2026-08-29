package ohsnap

import (
	"math/rand/v2"
	"testing"
)

func TestWeighted_Generate_Distribution(t *testing.T) {
	rnd := rand.New(rand.NewPCG(42, 54))

	// Two arbitraries: 1 weight for "a", 3 weights for "b"
	type letter string
	arb := Weighted[letter](rnd, map[int]Arbitrary[letter]{
		1: OneOfValue(rnd, letter("a")),
		3: OneOfValue(rnd, letter("b")),
	})

	counts := map[letter]int{
		"a": 0,
		"b": 0,
	}
	const samples = 4000
	for i := 0; i < samples; i++ {
		val := First(arb.Generate())
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

	arb := Weighted[int](rnd, map[int]Arbitrary[int]{
		0:  OneOfValue(rnd, 1),
		-2: OneOfValue(rnd, 2),
		5:  OneOfValue(rnd, 42),
	})

	// Only "42" should ever be generated
	for i := 0; i < 100; i++ {
		val := First(arb.Generate())
		if val != 42 {
			t.Errorf("Expected only 42, got %v", val)
		}
	}
}

func TestWeighted_EmptyOrAllZeroWeightsReturnsZeroValue(t *testing.T) {
	rnd := rand.New(rand.NewPCG(3, 4))

	arb := Weighted[int](rnd, map[int]Arbitrary[int]{})
	val := First(arb.Generate())
	if val != 0 {
		t.Errorf("Expected zero value for empty input, got %v", val)
	}

	arb2 := Weighted[int](rnd, map[int]Arbitrary[int]{0: OneOfValue(rnd, 99)})
	val2 := First(arb2.Generate())
	if val2 != 0 {
		t.Errorf("Expected zero value for all zero weights, got %v", val2)
	}
}

func TestWeightedDelegatesShrinkToLastPick(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(6, 6))

	small := OneOfValue(rnd, 7) // only passing values
	large := ArbitraryInt(rnd, 0, 100000)
	arb := Weighted[int](rnd, map[int]Arbitrary[int]{
		1: small,
		3: large,
	})

	value, shrunk := findSimplestBadCaseWith(1000, arb, func(x int) bool { return x == 7 }, traceShrinking(t))
	if value == nil {
		t.Fatal("expected property to fail")
	}
	if *shrunk != 0 {
		t.Errorf("shrunk = %d, want 0 (delegated halving of the large generator)", *shrunk)
	}
}
