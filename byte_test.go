package ohsnap

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestByte(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("math laws", func(t *testing.T) {
		checkMath(t, ArbitraryByte(rnd, 0, math.MaxUint8))
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, ArbitraryByte, 0, math.MaxUint8, math.MaxUint8)

		t.Run("shringking", func(t *testing.T) {
			checkSrinking(t, ArbitraryByte(rnd, 0, math.MaxUint8), 1, func(x byte) bool {
				return x == 0
			})
		})
	})
}
