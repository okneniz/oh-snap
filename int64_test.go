package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestInt64(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryInt64(rnd, 0, math.MaxInt64),
			ArbitraryInt64(rnd, 0, math.MaxInt64),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int64, int64]) bool {
				a := p.First
				b := p.Second
				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int64, int64]) bool {
				a := p.First
				b := p.Second
				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryInt64(rnd, 0, math.MaxInt64),
			Combine(
				ArbitraryInt64(rnd, 0, math.MaxInt64),
				ArbitraryInt64(rnd, 0, math.MaxInt64),
			),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int64, Pair[int64, int64]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int64, Pair[int64, int64]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arb := Combine(
			ArbitraryInt64(rnd, 0, math.MaxInt64),
			Combine(
				ArbitraryInt64(rnd, 0, math.MaxInt64),
				ArbitraryInt64(rnd, 0, math.MaxInt64),
			),
		)

		Check(t, iterations, arb, func(p Pair[int64, Pair[int64, int64]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second
			return a*(b+c) == a*b+a*c
		})
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryInt64, 0, math.MaxInt16*2, math.MaxInt64)

		t.Run("shringking", func(t *testing.T) {
			checkSrinking(t, ArbitraryInt64(rnd, 0, math.MaxInt64), 1, func(x int64) bool {
				return x == 0
			})
		})
	})
}
