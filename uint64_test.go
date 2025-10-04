package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestUint64(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("math laws", func(t *testing.T) {
		checkMath(t, ArbitraryUint64(rnd, 0, math.MaxUint))
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryUint64, 0, math.MaxUint16*2, math.MaxUint64)

		t.Run("shringking", func(t *testing.T) {
			checkSrinking(t, ArbitraryUint64(rnd, 0, math.MaxUint64), 1, func(x uint64) bool {
				return x == 0
			})
		})
	})
}
