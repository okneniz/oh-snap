package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestInt(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryInt(rnd, 0, math.MaxInt),
			ArbitraryInt(rnd, 0, math.MaxInt),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int, int]) bool {
				a := p.First
				b := p.Second
				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int, int]) bool {
				a := p.First
				b := p.Second
				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryInt(rnd, 0, math.MaxInt),
			Combine(
				ArbitraryInt(rnd, 0, math.MaxInt),
				ArbitraryInt(rnd, 0, math.MaxInt),
			),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int, Pair[int, int]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[int, Pair[int, int]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arb := Combine(
			ArbitraryInt(rnd, 0, math.MaxInt),
			Combine(
				ArbitraryInt(rnd, 0, math.MaxInt),
				ArbitraryInt(rnd, 0, math.MaxInt),
			),
		)

		Check(t, iterations, arb, func(p Pair[int, Pair[int, int]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second
			return a*(b+c) == a*b+a*c
		})
	})

	t.Run("same result for single number in range", func(t *testing.T) {
		for i := int(0); i <= math.MaxInt16; i++ {
			arb := ArbitraryInt(rnd, i, i)
			Check(t, 3, arb, func(x int) bool { return x == i })
		}

		arb := ArbitraryInt(rnd, math.MaxInt, math.MaxInt)
		Check(t, 3, arb, func(x int) bool { return x == math.MaxInt })
	})

}
