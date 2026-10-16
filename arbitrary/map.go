package arbitrary

import (
	"iter"
	"maps"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
)

type arbitraryMap[K comparable, V any] struct {
	rand    *rand.Rand
	key     ohsnap.Arbitrary[K]
	value   ohsnap.Arbitrary[V]
	minSize int
	maxSize int
}

// Map - return generator for arbitrary maps.
// rnd - pseudo-random number generator.
// key and value - generators of arbitrary keys and values.
// minSize and maxSize - bounds of size of generated values.
func Map[K comparable, V any](
	rnd *rand.Rand,
	key ohsnap.Arbitrary[K],
	value ohsnap.Arbitrary[V],
	minSize, maxSize int,
) ohsnap.Arbitrary[map[K]V] {
	return &arbitraryMap[K, V]{
		rand:    rnd,
		key:     key,
		value:   value,
		minSize: minSize,
		maxSize: maxSize,
	}
}

func (a *arbitraryMap[K, V]) Generate() map[K]V {
	size := a.rand.IntN(a.maxSize-a.minSize+1) + a.minSize
	result := make(map[K]V, size)

	for range size {
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

func (a *arbitraryMap[K, V]) Shrink(input map[K]V) iter.Seq[map[K]V] {
	return func(yield func(map[K]V) bool) {
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

			if !yield(smaller) {
				return
			}
		}

		for k, v := range input {
			for smallerV := range a.value.Shrink(v) {
				newMap := make(map[K]V, len(input))
				maps.Copy(newMap, input)

				newMap[k] = smallerV
				if !yield(newMap) {
					return
				}
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

				if !yield(newMap) {
					return
				}
			}

			empty := make(map[K]V)
			if !yield(empty) {
				return
			}
		}
	}
}
