package shrink_test

import (
	"fmt"
	"iter"
	"slices"

	"github.com/okneniz/oh-snap/shrink"
)

// Halving halves the distance between the value and the target.
// Halving(0) is the default strategy of numeric generators.
func ExampleHalving() {
	fmt.Println(slices.Collect(shrink.Halving(0)(5)))
	// Output: [2 1 0]
}

// Halving toward a non-zero target keeps candidates useful when the
// generator range does not include zero: the chain stays inside [100, 200]
// and walks toward the lower bound.
func ExampleHalving_lowerBound() {
	fmt.Println(slices.Collect(shrink.Halving(100)(200)))
	// Output: [150 125 112 106 103 101 100]
}

// Boundaries yields interesting values first: zeros, ones, range bounds.
// The input value itself is skipped to keep the search terminating.
func ExampleBoundaries() {
	fmt.Println(slices.Collect(shrink.Boundaries(0, 1, -1)(42)))
	// Output: [0 1 -1]
}

// Peel removes characters from the end one by one, like the default
// strategy of the string generator.
func ExamplePeel() {
	fmt.Println(slices.Collect(shrink.Peel()("abc")))
	// Output: [ab a ]
}

// HalveLength reduces the string length in O(log n) steps instead of O(n).
func ExampleHalveLength() {
	fmt.Println(slices.Collect(shrink.HalveLength()("abcdef")))
	// Output: [abc a ]
}

// MinimizeChars replaces one character at a time with the minimal
// character of the generator alphabet.
func ExampleMinimizeChars() {
	fmt.Println(slices.Collect(shrink.MinimizeChars('a')("cba")))
	// Output: [aba caa]
}

// Concat yields all candidates of the first strategy, then of the second.
func ExampleConcat() {
	s := shrink.Concat(shrink.Halving(0), shrink.Boundaries(7))
	fmt.Println(slices.Collect(s(5)))
	// Output: [2 1 0 7]
}

// Limit caps the number of candidates, useful for very wide strategies.
func ExampleLimit() {
	s := shrink.Limit(2, shrink.Halving(0))
	fmt.Println(slices.Collect(s(5)))
	// Output: [2 1]
}

// Interleave alternates candidates of two strategies: a, b, a, b, ...
// Handy for pairs and containers to approach both components at once.
func ExampleInterleave() {
	s := shrink.Interleave(shrink.Boundaries(1, 2), shrink.Boundaries(3, 4))
	fmt.Println(slices.Collect(s(9)))
	// Output: [1 3 2 4]
}

// NoShrink disables shrinking.
func ExampleNoShrink() {
	fmt.Println(slices.Collect(shrink.NoShrink[int]()(5)))
	// Output: []
}

// Shrinker is a plain function, so custom strategies are trivial to write.
// countdown is a naive strategy that tries value-1, value-2, ...
func ExampleShrinker() {
	countdown := shrink.Shrinker[int](func(value int) iter.Seq[int] {
		return func(yield func(int) bool) {
			for v := value - 1; v > 0; v-- {
				if !yield(v) {
					return
				}
			}
		}
	})

	fmt.Println(slices.Collect(countdown(5)))
	// Output: [4 3 2 1]
}
