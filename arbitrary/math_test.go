package arbitrary

import (
	"testing"

	ohsnap "github.com/okneniz/oh-snap"
)

func checkMath[T num](t *testing.T, arb ohsnap.Arbitrary[T]) {
	t.Helper()

	const iterations = 1000

	t.Run("commutative", func(t *testing.T) {
		arbC := ohsnap.Combine(arb, arb)

		t.Run("addition", func(t *testing.T) {
			ohsnap.Check(t, iterations, arbC, func(p ohsnap.Pair[T, T]) bool {
				a := p.First
				b := p.Second

				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			ohsnap.Check(t, iterations, arbC, func(p ohsnap.Pair[T, T]) bool {
				a := p.First
				b := p.Second

				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		arbC := ohsnap.Combine(arb, ohsnap.Combine(arb, arb))

		t.Run("addition", func(t *testing.T) {
			ohsnap.Check(t, iterations, arbC, func(p ohsnap.Pair[T, ohsnap.Pair[T, T]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second

				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			ohsnap.Check(t, iterations, arbC, func(p ohsnap.Pair[T, ohsnap.Pair[T, T]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second

				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arbC := ohsnap.Combine(arb, ohsnap.Combine(arb, arb))

		ohsnap.Check(t, iterations, arbC, func(p ohsnap.Pair[T, ohsnap.Pair[T, T]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second

			return a*(b+c) == a*b+a*c
		})
	})
}
