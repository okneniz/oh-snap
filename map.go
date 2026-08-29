package ohsnap

import (
	"iter"
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

func (a *arbitraryMap[K, V]) Generate() iter.Seq[map[K]V] {
	return func(yield func(map[K]V) bool) {
		for {
			size := a.rand.IntN(a.maxSize-a.minSize+1) + int(a.minSize)
			result := make(map[K]V, size)

			for i := 0; i < size; i++ {
				for {
					key := First(a.key.Generate())

					if _, exists := result[key]; !exists {
						result[key] = First(a.value.Generate())
						break
					}
				}
			}

			if !yield(result) {
				return
			}
		}
	}
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
				for k2, v2 := range input {
					newMap[k2] = v2
				}
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
