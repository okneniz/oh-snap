package ohsnap

import (
	"math/rand/v2"
	"testing"
)

type num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func checkBounds[T num](
	t *testing.T,
	rnd *rand.Rand,
	arbConstructor func(*rand.Rand, T, T) Arbitrary[T],
	from, to T,
	max T,
) {
	t.Helper()

	t.Run("same result for single number in range", func(t *testing.T) {
		const iterations = 3

		for i := from; true; i++ {
			arb := arbConstructor(rnd, i, i)

			Check(t, iterations, arb, func(x T) bool { return x == i })

			if i == to {
				break
			}
		}

		arb := arbConstructor(rnd, max, max) // because int64 have huge bounds
		Check(t, iterations, arb, func(x T) bool { return x == max })
	})
}
