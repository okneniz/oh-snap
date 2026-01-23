package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestInt16(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("math laws", func(t *testing.T) {
		checkMath(t, ArbitraryInt16(rnd, 0, math.MaxInt16))
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryInt16, 0, math.MaxInt16, math.MaxInt16)

		t.Run("shringking", func(t *testing.T) {
			checkShrinking(t, ArbitraryInt16(rnd, 0, math.MaxInt16), 1, func(x int16) bool {
				return x == 0
			})
		})
	})
}
