package ohsnap

import (
	"math/rand"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	t.Parallel()

	iterations := 100000

	seed := time.Now().UnixNano()

	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewSource(seed))
	arb := ArbitraryString(rnd, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 3, 10)

	t.Run("concatenation", func(t *testing.T) {
		t.Parallel()

		t.Run("associative result", func(t *testing.T) {
			Check(t, iterations, arb, func(a string) bool {
				b := arb.Generate()
				c := arb.Generate()
				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("associative length of result", func(t *testing.T) {
			Check(t, iterations, arb, func(a string) bool {
				b := arb.Generate()
				return len(a+b) == len(a)+len(b)
			})
		})
	})
}
