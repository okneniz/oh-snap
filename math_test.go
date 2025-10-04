package ohsnap

import "testing"

func checkMath[T num](t *testing.T, arb Arbitrary[T]) {
	t.Helper()

	const iterations = 1000

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arbC := Combine(arb, arb)

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arbC, func(p Pair[T, T]) bool {
				a := p.First
				b := p.Second
				return a+b == b+a
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arbC, func(p Pair[T, T]) bool {
				a := p.First
				b := p.Second
				return a*b == b*a
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arbC := Combine(arb, Combine(arb, arb))

		t.Run("addition", func(t *testing.T) {
			Check(t, iterations, arbC, func(p Pair[T, Pair[T, T]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("multiplication", func(t *testing.T) {
			Check(t, iterations, arbC, func(p Pair[T, Pair[T, T]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a*b)*c == a*(b*c)
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arbC := Combine(arb, Combine(arb, arb))

		Check(t, iterations, arbC, func(p Pair[T, Pair[T, T]]) bool {
			a := p.First
			b := p.Second.First
			c := p.Second.Second
			return a*(b+c) == a*b+a*c
		})
	})
}
