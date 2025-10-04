package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestUint8(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryUint8(rnd, 0, math.MaxUint8),
			ArbitraryUint8(rnd, 0, math.MaxUint8),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint8, uint8]) bool {
				a := p.First
				b := p.Second
				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint8, uint8]) bool {
				a := p.First
				b := p.Second
				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryUint8(rnd, 0, math.MaxUint8),
			Combine(
				ArbitraryUint8(rnd, 0, math.MaxUint8),
				ArbitraryUint8(rnd, 0, math.MaxUint8),
			),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint8, Pair[uint8, uint8]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint8, Pair[uint8, uint8]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arb := Combine(
			ArbitraryUint8(rnd, 0, math.MaxUint8),
			Combine(
				ArbitraryUint8(rnd, 0, math.MaxUint8),
				ArbitraryUint8(rnd, 0, math.MaxUint8),
			),
		)

		Check(t, iterations, arb, func(p Pair[uint8, Pair[uint8, uint8]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second
			return a*(b+c) == a*b+a*c
		})
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryUint8, 0, math.MaxUint8, math.MaxUint8)

		t.Run("shringking", func(t *testing.T) {
			checkSrinking(t, ArbitraryUint8(rnd, 0, math.MaxUint8), 1, func(x uint8) bool {
				return x == 0
			})
		})
	})
}
