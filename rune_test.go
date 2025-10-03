package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
	"unicode"
)

func TestRune(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryRune(rnd, 0, unicode.MaxRune),
			ArbitraryRune(rnd, 0, unicode.MaxRune),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[rune, rune]) bool {
				a := p.First
				b := p.Second
				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[rune, rune]) bool {
				a := p.First
				b := p.Second
				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryRune(rnd, 0, unicode.MaxRune),
			Combine(
				ArbitraryRune(rnd, 0, unicode.MaxRune),
				ArbitraryRune(rnd, 0, unicode.MaxRune),
			),
		)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[rune, Pair[rune, rune]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[rune, Pair[rune, rune]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arb := Combine(
			ArbitraryRune(rnd, 0, unicode.MaxRune),
			Combine(
				ArbitraryRune(rnd, 0, unicode.MaxRune),
				ArbitraryRune(rnd, 0, unicode.MaxRune),
			),
		)

		Check(t, iterations, arb, func(p Pair[rune, Pair[rune, rune]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second
			return a*(b+c) == a*b+a*c
		})
	})

	t.Run("same result for single number in range", func(t *testing.T) {
		for i := int32(0); i <= math.MaxInt16; i++ {
			arb := ArbitraryRune(rnd, i, i)
			Check(t, 3, arb, func(x rune) bool { return x == i })
		}

		arb := ArbitraryRune(rnd, math.MaxInt32, math.MaxInt32)
		Check(t, 3, arb, func(x rune) bool { return x == math.MaxInt32 })
	})
}
