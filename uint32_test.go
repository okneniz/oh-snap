package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestUint32(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryUint32(rnd, 0, math.MaxUint32),
			ArbitraryUint32(rnd, 0, math.MaxUint32),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint32, uint32]) bool {
				a := p.First
				b := p.Second
				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint32, uint32]) bool {
				a := p.First
				b := p.Second
				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryUint32(rnd, 0, math.MaxUint32),
			Combine(
				ArbitraryUint32(rnd, 0, math.MaxUint32),
				ArbitraryUint32(rnd, 0, math.MaxUint32),
			),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint32, Pair[uint32, uint32]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[uint32, Pair[uint32, uint32]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arb := Combine(
			ArbitraryUint32(rnd, 0, math.MaxUint32),
			Combine(
				ArbitraryUint32(rnd, 0, math.MaxUint32),
				ArbitraryUint32(rnd, 0, math.MaxUint32),
			),
		)

		Check(t, iterations, arb, func(p Pair[uint32, Pair[uint32, uint32]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second
			return a*(b+c) == a*b+a*c
		})
	})

	t.Run("same result for single number in range", func(t *testing.T) {
		for i := uint32(0); i <= math.MaxUint16; i++ {
			arb := ArbitraryUint32(rnd, i, i)
			Check(t, 3, arb, func(x uint32) bool { return x == i })
		}

		arb := ArbitraryUint32(rnd, math.MaxInt32, math.MaxInt32)
		Check(t, 3, arb, func(x uint32) bool { return x == math.MaxInt32 })
	})
}
