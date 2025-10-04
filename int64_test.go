package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestInt64(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("math laws", func(t *testing.T) {
		checkMath(t, ArbitraryInt64(rnd, 0, math.MaxInt64))
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryInt64, 0, math.MaxInt16*2, math.MaxInt64)

		t.Run("shringking", func(t *testing.T) {
			checkSrinking(t, ArbitraryInt64(rnd, 0, math.MaxInt64), 1, func(x int64) bool {
				return x == 0
			})
		})
	})
}
