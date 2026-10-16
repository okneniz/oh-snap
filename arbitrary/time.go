package arbitrary

import (
	"iter"
	"math/rand/v2"
	"time"

	ohsnap "github.com/okneniz/oh-snap"
)

type arbitraryTime struct {
	rand     *rand.Rand
	from, to time.Time
}

// Time - return generator for arbitrary time values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func Time(rnd *rand.Rand, from, to time.Time) ohsnap.Arbitrary[time.Time] {
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
