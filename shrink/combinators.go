package shrink

import (
	"iter"
)

// Concat returns a strategy that yields all candidates of the first
// strategy, then all candidates of the second one, and so on.
func Concat[T any](shrinkers ...Shrinker[T]) Shrinker[T] {
	return func(value T) iter.Seq[T] {
		return func(yield func(T) bool) {
			for _, s := range shrinkers {
				cont := true
				s(value)(func(v T) bool {
					cont = yield(v)
					return cont
				})
				if !cont {
					return
				}
			}
		}
	}
}

// Limit returns a strategy that yields at most n candidates of the
// underlying strategy.
func Limit[T any](n int, s Shrinker[T]) Shrinker[T] {
	return func(value T) iter.Seq[T] {
		return func(yield func(T) bool) {
			count := 0
			s(value)(func(v T) bool {
				if count >= n {
					return false
				}
				count++
				return yield(v)
			})
		}
	}
}

// Interleave returns a strategy that alternates candidates of two
// strategies: a, b, a, b, ... It is useful for pairs and containers
// to approach both components at the same time.
func Interleave[T any](a, b Shrinker[T]) Shrinker[T] {
	return func(value T) iter.Seq[T] {
		return func(yield func(T) bool) {
			nextA, stopA := iter.Pull(a(value))
			defer stopA()

			nextB, stopB := iter.Pull(b(value))
			defer stopB()

			for {
				va, okA := nextA()
				if okA {
					if !yield(va) {
						return
					}
				}

				vb, okB := nextB()
				if okB {
					if !yield(vb) {
						return
					}
				}

				if !okA && !okB {
					return
				}
			}
		}
	}
}
