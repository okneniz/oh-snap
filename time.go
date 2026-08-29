package ohsnap

import (
	"iter"
	"math/rand/v2"
	"time"
)

type arbitraryTime struct {
	rand     *rand.Rand
	from, to time.Time
}

// ArbitraryTime - return generator for arbitrary time values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryTime(rnd *rand.Rand, from, to time.Time) Arbitrary[time.Time] {
	if from.After(to) {
		from, to = to, from
	}

	return &arbitraryTime{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryTime) Generate() iter.Seq[time.Time] {
	return func(yield func(time.Time) bool) {
		n := a.to.UnixNano() - a.from.UnixNano()

		for {
			nanoSeconds := a.rand.Int64N(n)
			value := time.Unix(0, a.from.UnixNano()+nanoSeconds)
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryTime) Shrink(t time.Time) iter.Seq[time.Time] {
	return func(yield func(time.Time) bool) {
		for value := t.UnixNano(); value != 0; {
			value /= 2
			shrunk := time.Unix(0, value)
			if !yield(shrunk) {
				return
			}
		}
	}
}
