package ohsnap

import (
	"math/rand"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewSource(seed))

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryBool(rnd),
			ArbitraryBool(rnd),
		)

		t.Run("conjunction (AND)", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[bool, bool]) bool {
				a := p.First
				b := p.Second
				return (a && b) == (b && a)
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[bool, bool]) bool {
				a := p.First
				b := p.Second
				return (a || b) == (b || a)
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryBool(rnd),
			Combine(
				ArbitraryBool(rnd),
				ArbitraryBool(rnd),
			),
		)

		t.Run("conjunction (AND)", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[bool, Pair[bool, bool]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return ((a && b) && c) == (a && (b && c))
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[bool, Pair[bool, bool]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return ((a || b) || c) == (a || (b || c))
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryBool(rnd),
			Combine(
				ArbitraryBool(rnd),
				ArbitraryBool(rnd),
			),
		)

		t.Run("conjunction (AND)", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[bool, Pair[bool, bool]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a && (b || c)) == ((a && b) || (a && c))
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[bool, Pair[bool, bool]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second
				return (a || (b && c)) == ((a || b) && (a || c))
			})
		})
	})

	t.Run("identity", func(t *testing.T) {
		t.Parallel()

		arb := ArbitraryBool(rnd)

		t.Run("conjunction (AND)", func(t *testing.T) {
			Check(t, iterations, arb, func(b bool) bool {
				return (b && true) == b
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			Check(t, iterations, arb, func(b bool) bool {
				return (b || false) == b
			})
		})
	})

	t.Run("idempotent", func(t *testing.T) {
		t.Parallel()

		arb := ArbitraryBool(rnd)

		t.Run("conjunction (AND)", func(t *testing.T) {
			Check(t, iterations, arb, func(b bool) bool {
				return (b && b) == b
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			Check(t, iterations, arb, func(b bool) bool {
				return (b || b) == b
			})
		})
	})

	t.Run("complementary", func(t *testing.T) {
		t.Parallel()

		arb := ArbitraryBool(rnd)

		t.Run("conjunction (AND)", func(t *testing.T) {
			Check(t, iterations, arb, func(b bool) bool {
				return (b && (!b)) == false
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			Check(t, iterations, arb, func(b bool) bool {
				return (b || (!b)) == true
			})
		})
	})

	t.Run("De Morgan's Laws", func(t *testing.T) {
		t.Parallel()

		arb := Combine(
			ArbitraryBool(rnd),
			ArbitraryBool(rnd),
		)

		t.Run("conjunction (AND)", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[bool, bool]) bool {
				a := p.First
				b := p.Second
				return !(a && b) == (!a || !b)
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			Check(t, iterations, arb, func(p Pair[bool, bool]) bool {
				a := p.First
				b := p.Second
				return !(a || b) == (!a && !b)
			})
		})
	})

	t.Run("involution", func(t *testing.T) {
		t.Parallel()

		arb := ArbitraryBool(rnd)

		t.Run("check", func(t *testing.T) {
			Check(t, iterations, arb, func(b bool) bool {
				return !(!b) == b
			})
		})
	})
}
