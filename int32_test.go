package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestInt32(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("math laws", func(t *testing.T) {
		checkMath(t, ArbitraryInt32(rnd, 0, math.MaxInt32))
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryInt32, 0, math.MaxInt16*2, math.MaxInt32)

		t.Run("shringking", func(t *testing.T) {
			checkSrinking(t, ArbitraryInt32(rnd, 0, math.MaxInt32), 1, func(x int32) bool {
				return x == 0
			})
		})
	})
}
