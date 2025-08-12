package ohsnap

import (
	"math/rand/v2"
)

type arbitraryRune struct {
	rand     *rand.Rand
	from, to rune
}

// ArbitraryRune - return generator for arbitrary rune values.
// - rnd - pseudo-random number generator.
// - from and to - bounds of generated values.
func ArbitraryRune(rnd *rand.Rand, from, to rune) Arbitrary[rune] {
	return &arbitraryRune{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryRune) Generate() rune {
	return a.rand.Int32N(a.to-a.from) + a.from
}

func (arbitraryRune) Shrink(value rune) []rune {
	var results []rune

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
