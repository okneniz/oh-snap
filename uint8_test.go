package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestUint8(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("math laws", func(t *testing.T) {
		checkMath(t, ArbitraryUint8(rnd, 0, math.MaxUint8))
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryUint8, 0, math.MaxUint8, math.MaxUint8)

		t.Run("shringking", func(t *testing.T) {
			checkSrinking(t, ArbitraryUint8(rnd, 0, math.MaxUint8), 1, func(x uint8) bool {
				return x == 0
			})
		})
	})
}
