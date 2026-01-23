package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestInt8(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("math laws", func(t *testing.T) {
		checkMath(t, ArbitraryInt8(rnd, 0, math.MaxInt8))
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryInt8, 0, math.MaxInt8, math.MaxInt8)

		t.Run("shringking", func(t *testing.T) {
			checkShrinking(t, ArbitraryInt8(rnd, 0, math.MaxInt8), 1, func(x int8) bool {
				return x == 0
			})
		})
	})
}
