package arbitrary

import (
	"math/rand/v2"
	"testing"
	"time"
	"unicode"

	ohsnap "github.com/okneniz/oh-snap"
)

func TestRuneFromTable(t *testing.T) {
	t.Parallel()

	const iterations = 10_000_000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	arbs := make([]ohsnap.Arbitrary[ohsnap.Pair[rune, *unicode.RangeTable]], 0)

	for _, x := range unicode.Categories {
		tbl := x

		arb := ohsnap.Map(
			RuneFromTable(rnd, tbl),
			func(x rune) ohsnap.Pair[rune, *unicode.RangeTable] {
				return ohsnap.Pair[rune, *unicode.RangeTable]{
					First:  x,
					Second: tbl,
				}
			},
		)

		arbs = append(arbs, arb)
	}

	arb := ohsnap.OneOf(rnd, arbs...)

	ohsnap.Check(t, iterations, arb, func(p ohsnap.Pair[rune, *unicode.RangeTable]) bool {
		char := p.First
		unicodeTable := p.Second
		return unicode.In(char, unicodeTable)
	})
}
