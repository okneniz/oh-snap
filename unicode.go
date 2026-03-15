package ohsnap

import (
	"math/rand/v2"
	"unicode"
)

// RuneFromTable - return generator for arbitrary rune values from unicode table.
// rnd - pseudo-random number generator.
// tbl - bounds for generated runes.
func RuneFromTable(rnd *rand.Rand, tbl *unicode.RangeTable) Arbitrary[rune] {
	total := 0
	for _, r := range tbl.R16 {
		count := int((r.Hi-r.Lo)/r.Stride) + 1
		total += count
	}
	for _, r := range tbl.R32 {
		count := int((r.Hi-r.Lo)/r.Stride) + 1
		total += count
	}
	return &arbitraryUnicode{
		rnd:   rnd,
		tbl:   tbl,
		total: total,
	}
}

type arbitraryUnicode struct {
	rnd   *rand.Rand
	tbl   *unicode.RangeTable
	total int
}

func (a *arbitraryUnicode) Generate() rune {
	if a.total == 0 {
		return 0 // fallback: no valid runes
	}
	n := a.rnd.IntN(a.total)
	tbl := a.tbl

	// Pick from R16
	for _, r := range tbl.R16 {
		count := int((r.Hi-r.Lo)/r.Stride) + 1
		if n < count {
			return rune(r.Lo + uint16(n)*r.Stride)
		}
		n -= count
	}
	// Pick from R32
	for _, r := range tbl.R32 {
		count := int((r.Hi-r.Lo)/r.Stride) + 1
		if n < count {
			return rune(r.Lo + uint32(n)*r.Stride)
		}
		n -= count
	}

	return 0 // fallback, should not happen
}

func (a *arbitraryUnicode) Shrink(value rune) []rune {
	// No shrinking for now
	return nil
}
