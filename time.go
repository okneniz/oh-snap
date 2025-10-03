package ohsnap

import (
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

func (a arbitraryTime) Generate() time.Time {
	n := a.to.UnixNano() - a.from.UnixNano()
	nanoSeconds := a.rand.Int64N(n)
	return time.Unix(0, a.from.UnixNano()+nanoSeconds)
}

func (arbitraryTime) Shrink(t time.Time) []time.Time {
	var results []time.Time

	for value := t.UnixNano(); value != 0; {
		value /= 2
		t = time.Unix(0, value)
		results = append(results, t)
	}

	return results
}
