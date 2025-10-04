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

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryInt, 0, math.MaxInt16*2, math.MaxInt)

		t.Run("shringking", func(t *testing.T) {
			checkSrinking(t, ArbitraryInt(rnd, 0, math.MaxInt), 1, func(x int) bool {
				return x == 0
			})
		})
	})
}
