package ohsnap

import (
	"math/rand"
	"testing"
	"time"
)

func TestInt(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()

	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewSource(seed))
	arb := ArbitraryInt(rnd, 0, 1000)

	t.Run("commutative", func(t *testing.T) {
		t.Parallel()

		t.Run("addition", func(t *testing.T) {
			prop := func(a int) bool {
				b := arb.Generate()
				return a+b == b+a
			}

			Check(t, arb, prop, DefaultConfig())
		})

		t.Run("multiplication", func(t *testing.T) {
			prop := func(a int) bool {
				b := arb.Generate()
				return a*b == b*a
			}

			Check(t, arb, prop, DefaultConfig())
		})
	})

	t.Run("associative", func(t *testing.T) {
		t.Parallel()

		t.Run("addition", func(t *testing.T) {
			prop := func(a int) bool {
				b := arb.Generate()
				c := arb.Generate()
				return (a+b)+c == a+(b+c)
			}

			Check(t, arb, prop, DefaultConfig())
		})

		t.Run("multiplication", func(t *testing.T) {
			prop := func(a int) bool {
				b := arb.Generate()
				c := arb.Generate()
				return (a*b)*c == a*(b*c)
			}

			Check(t, arb, prop, DefaultConfig())
		})
	})

	t.Run("distributive", func(t *testing.T) {
		prop := func(a int) bool {
			b := arb.Generate()
			c := arb.Generate()
			return a*(b+c) == a*b+a*c
		}

		Check(t, arb, prop, DefaultConfig())
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()

	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewSource(seed))
	arb := ArbitraryString(rnd, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 3, 10)

	t.Run("concatenation", func(t *testing.T) {
		t.Parallel()

		t.Run("associative result", func(t *testing.T) {
			prop := func(a string) bool {
				b := arb.Generate()
				c := arb.Generate()
				return (a+b)+c == a+(b+c)
			}

			Check(t, arb, prop, DefaultConfig())
		})

		t.Run("associative length of result", func(t *testing.T) {
			prop := func(a string) bool {
				b := arb.Generate()
				return len(a+b) == len(a)+len(b)
			}

			Check(t, arb, prop, DefaultConfig())
		})
	})
}
