package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestUint64(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryUint64(rnd, 0, math.MaxUint64),
			ArbitraryUint64(rnd, 0, math.MaxUint64),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint64, uint64]) bool {
				a := p.First
				b := p.Second
				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint64, uint64]) bool {
				a := p.First
				b := p.Second
				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryUint64(rnd, 0, math.MaxUint64),
			Combine(
				ArbitraryUint64(rnd, 0, math.MaxUint64),
				ArbitraryUint64(rnd, 0, math.MaxUint64),
			),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint64, Pair[uint64, uint64]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint64, Pair[uint64, uint64]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arb := Combine(
			ArbitraryUint64(rnd, 0, math.MaxUint64),
			Combine(
				ArbitraryUint64(rnd, 0, math.MaxUint64),
				ArbitraryUint64(rnd, 0, math.MaxUint64),
			),
		)

		Check(t, iterations, arb, func(p Pair[uint64, Pair[uint64, uint64]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second
			return a*(b+c) == a*b+a*c
		})
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryUint64, 0, math.MaxUint16*2, math.MaxUint64)

		t.Run("shringking", func(t *testing.T) {
			checkSrinking(t, ArbitraryUint64(rnd, 0, math.MaxUint64), 1, func(x uint64) bool {
				return x == 0
			})
		})
	})
}
