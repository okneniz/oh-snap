package arbitrary

import (
	"math/rand/v2"
	"testing"
	"time"

	ohsnap "github.com/okneniz/oh-snap"
)

// and, or and not wrap boolean operators in functions, so algebra laws like
// idempotence can be stated without tripping redundant-expression checks.
func and(a, b bool) bool { return a && b }

func or(a, b bool) bool { return a || b }

func not(b bool) bool { return !b }

func TestBool(t *testing.T) {
	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("commutative", func(t *testing.T) {
		arb := ohsnap.Combine(
			Bool(rnd),
			Bool(rnd),
		)

		t.Run("conjunction (AND)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(p ohsnap.Pair[bool, bool]) bool {
				a := p.First
				b := p.Second

				return (a && b) == (b && a)
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(p ohsnap.Pair[bool, bool]) bool {
				a := p.First
				b := p.Second

				return (a || b) == (b || a)
			})
		})
	})

	t.Run("associative", func(t *testing.T) {
		arb := ohsnap.Combine(
			Bool(rnd),
			ohsnap.Combine(
				Bool(rnd),
				Bool(rnd),
			),
		)

		t.Run("conjunction (AND)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(p ohsnap.Pair[bool, ohsnap.Pair[bool, bool]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second

				return ((a && b) && c) == (a && (b && c))
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(p ohsnap.Pair[bool, ohsnap.Pair[bool, bool]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second

				return ((a || b) || c) == (a || (b || c))
			})
		})
	})

	t.Run("distributive", func(t *testing.T) {
		arb := ohsnap.Combine(
			Bool(rnd),
			ohsnap.Combine(
				Bool(rnd),
				Bool(rnd),
			),
		)

		t.Run("conjunction (AND)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(p ohsnap.Pair[bool, ohsnap.Pair[bool, bool]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second

				return (a && (b || c)) == ((a && b) || (a && c))
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(p ohsnap.Pair[bool, ohsnap.Pair[bool, bool]]) bool {
				a := p.First
				b := p.Second.First
				c := p.Second.Second

				return (a || (b && c)) == ((a || b) && (a || c))
			})
		})
	})

	t.Run("identity", func(t *testing.T) {
		arb := Bool(rnd)

		t.Run("conjunction (AND)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(b bool) bool {
				return (b && true) == b
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(b bool) bool {
				return (b || false) == b
			})
		})
	})

	t.Run("idempotent", func(t *testing.T) {
		arb := Bool(rnd)

		t.Run("conjunction (AND)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(b bool) bool {
				return and(b, b) == b
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(b bool) bool {
				return or(b, b) == b
			})
		})
	})

	t.Run("complementary", func(t *testing.T) {
		arb := Bool(rnd)

		t.Run("conjunction (AND)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(b bool) bool {
				return (b && (!b)) == false
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(b bool) bool {
				return (b || (!b)) == true
			})
		})
	})

	t.Run("De Morgan's Laws", func(t *testing.T) {
		arb := ohsnap.Combine(
			Bool(rnd),
			Bool(rnd),
		)

		t.Run("conjunction (AND)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(p ohsnap.Pair[bool, bool]) bool {
				a := p.First
				b := p.Second

				return not(and(a, b)) == or(not(a), not(b))
			})
		})

		t.Run("disjunction (OR)", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(p ohsnap.Pair[bool, bool]) bool {
				a := p.First
				b := p.Second

				return not(or(a, b)) == and(not(a), not(b))
			})
		})
	})

	t.Run("involution", func(t *testing.T) {
		arb := Bool(rnd)

		t.Run("check", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(b bool) bool {
				return not(not(b)) == b
			})
		})
	})
}
