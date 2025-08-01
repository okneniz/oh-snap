package ohsnap

import (
	"math/rand/v2"
)

type arbitraryMap[K comparable, V any] struct {
	rand    *rand.Rand
	key     Arbitrary[K]
	value   Arbitrary[V]
	minSize int
	maxSize int
}

func ArbitraryMap[K comparable, V any](
	rnd *rand.Rand,
	key Arbitrary[K],
	value Arbitrary[V],
	minSize, maxSize int,
) Arbitrary[map[K]V] {
	return &arbitraryMap[K, V]{
		rand:    rnd,
		key:     key,
		value:   value,
		minSize: minSize,
		maxSize: maxSize,
	}
}

func (a *arbitraryMap[K, V]) Generate() map[K]V {
	size := a.rand.IntN(a.maxSize-a.minSize+1) + int(a.minSize)
	result := make(map[K]V, size)

	for i := 0; i < size; i++ {
		for {
			key := a.key.Generate()

			if _, exists := result[key]; !exists {
				result[key] = a.value.Generate()
				break
			}
		}
	}
	return result
}

func (a *arbitraryMap[K, V]) Shrink(input map[K]V) []map[K]V {
	var shrunk []map[K]V

	if len(input) > 0 {
		halfSize := len(input) / 2
		smaller := make(map[K]V, halfSize)
		i := 0
		for k, v := range input {
			if i >= halfSize {
				break
			}
			smaller[k] = v
			i++
		}
		shrunk = append(shrunk, smaller)
	}

	for k, v := range input {
		for _, smallerV := range a.value.Shrink(v) {
			newMap := make(map[K]V, len(input))
			for k2, v2 := range input {
				newMap[k2] = v2
			}
			newMap[k] = smallerV
			shrunk = append(shrunk, newMap)
		}
	}

	if len(input) > 0 {
		for k := range input {
			newMap := make(map[K]V, len(input)-1)
			for k2, v := range input {
				if k2 != k {
					newMap[k2] = v
				}
			}
			shrunk = append(shrunk, newMap)
		}
	}

	if len(input) > 0 {
		shrunk = append(shrunk, make(map[K]V))
	}

	return shrunk
}
