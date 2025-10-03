package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestInt8(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryInt8(rnd, 0, math.MaxInt8),
			ArbitraryInt8(rnd, 0, math.MaxInt8),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int8, int8]) bool {
				a := p.First
				b := p.Second
				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int8, int8]) bool {
				a := p.First
				b := p.Second
				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryInt8(rnd, 0, math.MaxInt8),
			Combine(
				ArbitraryInt8(rnd, 0, math.MaxInt8),
				ArbitraryInt8(rnd, 0, math.MaxInt8),
			),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int8, Pair[int8, int8]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int8, Pair[int8, int8]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arb := Combine(
			ArbitraryInt8(rnd, 0, math.MaxInt8),
			Combine(
				ArbitraryInt8(rnd, 0, math.MaxInt8),
				ArbitraryInt8(rnd, 0, math.MaxInt8),
			),
		)

		Check(t, iterations, arb, func(p Pair[int8, Pair[int8, int8]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second
			return a*(b+c) == a*b+a*c
		})
	})

	t.Run("same result for single number in range", func(t *testing.T) {
		for i := int8(0); true; i++ {
			arb := ArbitraryInt8(rnd, i, i)
			Check(t, 3, arb, func(x int8) bool { return x == i })

			if i == math.MaxInt8 {
				break
			}
		}
	})
}
