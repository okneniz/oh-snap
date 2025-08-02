package ohsnap

import (
	"math/rand/v2"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("bounds", func(t *testing.T) {
		start := time.Now().Add(-100 * time.Hour)
		finish := time.Now().Add(100 * time.Hour)

		arb := ArbitraryTime(rnd, start, finish)

		Check(t, iterations, arb, func(x time.Time) bool {
			return (x.After(start) || x.Equal(start)) &&
				(x.Before(finish) || x.Equal(finish))
		})
	})

	t.Run("format", func(t *testing.T) {
		arb := Combine(
			OneOfValue(
				rnd,
				time.UnixDate,
				time.RubyDate,
				time.RFC850,
				time.RFC1123,
				time.RFC1123Z,
				time.RFC3339,
				time.RFC3339Nano,
			),
			ArbitraryTime(
				rnd,
				time.Now(),
				time.Now().Add(3*time.Hour),
			),
		)

		Check(t, iterations, arb, func(p Pair[string, time.Time]) bool {
			layout := p.First
			before := p.Second

			formatted := before.Format(layout)

			parsed, err := time.Parse(layout, formatted)
			if err != nil {
				t.Log("layout", layout)
				t.Log("time", before)
				t.Log("error", err)

				return false
			}

			before = before.Truncate(time.Second)
			parsed = parsed.Truncate(time.Second)

			return before.Equal(parsed)
		})
	})
}
