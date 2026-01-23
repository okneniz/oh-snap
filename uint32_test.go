package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestUint32(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("math laws", func(t *testing.T) {
		checkMath(t, ArbitraryUint32(rnd, 0, math.MaxUint32))
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryUint32, 0, math.MaxUint16*2, math.MaxUint32)

		t.Run("shringking", func(t *testing.T) {
			checkShrinking(t, ArbitraryUint32(rnd, 0, math.MaxUint32), 1, func(x uint32) bool {
				return x == 0
			})
		})
	})
}
