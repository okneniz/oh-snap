package ohsnap

import (
	"math/rand"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewSource(seed))

	keys := ArbitraryString(rnd, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 3, 10)
	values := ArbitraryInt(rnd, 5, 100)
	arb := ArbitraryMap(rnd, keys, values, 3, 10)

	t.Run("check get / put", func(t *testing.T) {
		Check(t, iterations, arb, func(m map[string]int) bool {
			var key string

			lenBefore := len(m)

			for k, _ := range m {
				key = k
				break
			}

			value := m[key]
			delete(m, key)
			m[key] = value

			lenAfter := len(m)

			return lenBefore == lenAfter
		})
	})
}
