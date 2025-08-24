package ohsnap

import (
	"math/rand/v2"
	"unicode"
)

// RuneFromTable - return generator for arbitrary rune values from unicode table.
// rnd - pseudo-random number generator.
// tbl - bounds for generated runes.
func RuneFromTable(rnd *rand.Rand, tbl *unicode.RangeTable) Arbitrary[rune] {
	allowed := make([]rune, 0) // not memory efficient, but works

	for r := rune(0); r <= unicode.MaxRune; r++ {
		if unicode.In(r, tbl) {
			allowed = append(allowed, r)
		}
	}

	return OneOfValue(rnd, allowed...)
}
