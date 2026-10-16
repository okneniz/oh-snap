package arbitrary

import (
	"math/rand/v2"
	"testing"

	ohsnap "github.com/okneniz/oh-snap"
)

type num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func checkBounds[T num](
	t *testing.T,
	rnd *rand.Rand,
	arbConstructor func(*rand.Rand, T, T) ohsnap.Arbitrary[T],
	from, to T,
	max T,
) {
	t.Helper()

	t.Run("same result for single number in range", func(t *testing.T) {
		const iterations = 3

		for i := from; true; i++ {
			arb := arbConstructor(rnd, i, i)

			ohsnap.Check(t, iterations, arb, func(x T) bool { return x == i })

			if i == to {
				break
			}
		}

		arb := arbConstructor(rnd, max, max) // because int64 have huge bounds
		ohsnap.Check(t, iterations, arb, func(x T) bool { return x == max })
	})
}
