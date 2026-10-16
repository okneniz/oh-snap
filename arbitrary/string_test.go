package arbitrary

import (
	"math/rand/v2"
	"testing"
	"time"

	ohsnap "github.com/okneniz/oh-snap"
)

func TestString(t *testing.T) {
	iterations := 100000

	seed := time.Now().UnixNano()

	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))
	arb := String(rnd, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 3, 10)

	t.Run("concatenation", func(t *testing.T) {
		t.Run("associative result", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(a string) bool {
				b := arb.Generate()
				c := arb.Generate()

				return (a+b)+c == a+(b+c)
			})
		})

		t.Run("associative length of result", func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(a string) bool {
				b := arb.Generate()
				return len(a+b) == len(a)+len(b)
			})
		})
	})
}
