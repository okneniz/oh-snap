package arbitrary

import (
	"math"
	"math/rand/v2"
	"testing"
	"time"
)

func TestRune(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	t.Run("math laws", func(t *testing.T) {
		checkMath(t, Rune(rnd, 0, math.MaxInt32))
	})

	t.Run("functional props", func(t *testing.T) {
		checkBounds(t, rnd, Rune, 0, math.MaxInt16*2, math.MaxInt32)

		t.Run("shringking", func(t *testing.T) {
			checkShrinking(t, Rune(rnd, 0, math.MaxInt32), 1, func(x rune) bool {
				return x == 0
			})
		})
	})
}
