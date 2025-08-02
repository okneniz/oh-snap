package ohsnap

import (
	"math/rand/v2"
	"time"
)

type arbitraryTime struct {
	rand     *rand.Rand
	from, to time.Time
}

func ArbitraryTime(rnd *rand.Rand, from, to time.Time) Arbitrary[time.Time] {
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
