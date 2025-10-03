package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestInt32(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryInt32(rnd, 0, math.MaxInt32),
			ArbitraryInt32(rnd, 0, math.MaxInt32),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int32, int32]) bool {
				a := p.First
				b := p.Second
				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int32, int32]) bool {
				a := p.First
				b := p.Second
				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryInt32(rnd, 0, math.MaxInt32),
			Combine(
				ArbitraryInt32(rnd, 0, math.MaxInt32),
				ArbitraryInt32(rnd, 0, math.MaxInt32),
			),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int32, Pair[int32, int32]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int32, Pair[int32, int32]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arb := Combine(
			ArbitraryInt32(rnd, 0, math.MaxInt32),
			Combine(
				ArbitraryInt32(rnd, 0, math.MaxInt32),
				ArbitraryInt32(rnd, 0, math.MaxInt32),
			),
		)

		Check(t, iterations, arb, func(p Pair[int32, Pair[int32, int32]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second
			return a*(b+c) == a*b+a*c
		})
	})

	t.Run("same result for single number in range", func(t *testing.T) {
		for i := int32(0); i <= math.MaxInt16; i++ {
			arb := ArbitraryInt32(rnd, i, i)
			Check(t, 3, arb, func(x int32) bool { return x == i })
		}

		arb := ArbitraryInt32(rnd, math.MaxInt32, math.MaxInt32)
		Check(t, 3, arb, func(x int32) bool { return x == math.MaxInt32 })
	})
}
