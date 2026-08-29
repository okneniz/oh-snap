package shrink

import (
	"iter"
)

// Peel returns a strategy that removes characters from the end of the
// string one by one (the default behavior of the string generator).
func Peel() Shrinker[string] {
	return func(value string) iter.Seq[string] {
		return func(yield func(string) bool) {
			for len(value) > 0 {
				value = value[:len(value)-1]
				if !yield(value) {
					return
				}
			}
		}
	}
}

// HalveLength returns a strategy that halves the string length toward zero.
// Compared to Peel it reduces the length in O(log n) steps instead of O(n).
func HalveLength() Shrinker[string] {
	return func(value string) iter.Seq[string] {
		return func(yield func(string) bool) {
			for len(value) > 0 {
				value = value[:len(value)/2]
				if !yield(value) {
					return
				}
			}
		}
	}
}

// MinimizeChars returns a strategy that replaces one character at a time
// with min. Use the minimal character of the generator alphabet.
// Positions that already contain min are skipped.
func MinimizeChars(min byte) Shrinker[string] {
	return func(value string) iter.Seq[string] {
		return func(yield func(string) bool) {
			for i := range value {
				if value[i] == min {
					continue
				}

				candidate := value[:i] + string(min) + value[i+1:]
				if !yield(candidate) {
					return
				}
			}
		}
	}
}
