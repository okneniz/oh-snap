package ohsnap

import (
	"math/rand/v2"
	"testing"
	"time"
	"unicode"
)

func TestRuneFromTable(t *testing.T) {
	t.Parallel()

	const iterations = 10_000_000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	arbs := make([]Arbitrary[Pair[rune, *unicode.RangeTable]], 0)

	for _, x := range unicode.Categories {
		tbl := x

		arb := Map(
			RuneFromTable(rnd, tbl),
			func(x rune) Pair[rune, *unicode.RangeTable] {
				return Pair[rune, *unicode.RangeTable]{
					First:  x,
					Second: tbl,
				}
			},
		)

		arbs = append(arbs, arb)
	}

	arb := OneOf(rnd, arbs)

	Check(t, iterations, arb, func(p Pair[rune, *unicode.RangeTable]) bool {
		char := p.First
		unicodeTable := p.Second
		return unicode.In(char, unicodeTable)
	})
}
