package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestUint(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryUint(rnd, 0, math.MaxUint),
			ArbitraryUint(rnd, 0, math.MaxUint),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint, uint]) bool {
				a := p.First
				b := p.Second
				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint, uint]) bool {
				a := p.First
				b := p.Second
				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryUint(rnd, 0, 1000),
			Combine(
				ArbitraryUint(rnd, 0, math.MaxUint),
				ArbitraryUint(rnd, 0, math.MaxUint),
			),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint, Pair[uint, uint]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint, Pair[uint, uint]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arb := Combine(
			ArbitraryUint(rnd, 0, math.MaxUint),
			Combine(
				ArbitraryUint(rnd, 0, math.MaxUint),
				ArbitraryUint(rnd, 0, math.MaxUint),
			),
		)

		Check(t, iterations, arb, func(p Pair[uint, Pair[uint, uint]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second
			return a*(b+c) == a*b+a*c
		})
	})
}
