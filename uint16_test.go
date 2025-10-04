package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestUint16(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("math laws", func(t *testing.T) {
		checkMath(t, ArbitraryUint16(rnd, 0, math.MaxUint16))
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryUint16, 0, math.MaxUint16, math.MaxUint16)

		t.Run("shringking", func(t *testing.T) {
			checkSrinking(t, ArbitraryUint16(rnd, 0, math.MaxUint16), 1, func(x uint16) bool {
				return x == 0
			})
		})
	})
}
