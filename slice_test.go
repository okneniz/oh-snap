package ohsnap

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestSlice(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewSource(seed))

	ints := ArbitraryInt(rnd, 3, 100)
	arb := ArbitrarySlice(rnd, ints, 3, 10)

	t.Run("double reverse change nothing", func(t *testing.T) {
		reverse := func(slice []int) {
			for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}

		Check(t, iterations, arb, func(source []int) bool {
			destination := make([]int, len(source))
			_ = copy(destination, source)

			reverse(destination)
			reverse(destination)

			return reflect.DeepEqual(source, destination)
		})
	})
}
