// Package ohsnap is a property-based testing library: instead of writing
// individual test cases, you declare a property that must hold for any
// input, and the library searches for a counterexample using random values.
// When a failing value is found, it is shrunk to the simplest form that
// still fails the property, which makes failures easy to understand.
//
// # Quick start
//
//	func TestSortSorts(t *testing.T) {
//		rnd := rand.New(rand.NewPCG(1, 2))
//		arb := ohsnap.ArbitrarySlice(rnd, ohsnap.ArbitraryInt(rnd, 0, 100), 0, 10)
//
//		ohsnap.Check(t, 1000, arb, func(xs []int) bool {
//			dst := slices.Clone(xs)
//			slices.Reverse(dst)
//			slices.Reverse(dst)
//			return slices.Equal(xs, dst)
//		})
//	}
//
// Check runs the property on iterations random values. On the first failure
// it reports both the original value and its shrunk minimal form.
//
// # Arbitraries
//
// The central abstraction is [Arbitrary]. Generate returns a lazy,
// effectively infinite stream of random values ([iter.Seq]): pull as many
// as you need with a range loop, or take a single value with [First].
// Shrink returns a lazy stream of candidates simpler than a given value.
//
// Generators for primitives cover all numeric types, bool, string, time and
// runes (including [RuneFromTable] for unicode tables):
//
//	arb := ohsnap.ArbitraryInt(rnd, -100, 100)
//	x := ohsnap.First(arb.Generate())
//
// Containers and combinators build bigger arbitraries from smaller ones:
// [ArbitrarySlice], [ArbitraryMap], [Combine] (pairs), [Map] (transform),
// [OneOf], [OneOfValue] and [Weighted] (choice). The [Builder] configures
// bounds fluently and produces any primitive or slice generator:
//
//	b := ohsnap.NewBuilder(rnd).MinInt(10).MaxInt(20).MaxSliceLen(5)
//	arb := b.IntSlice()
//
// # Shrinking
//
// When the property fails, the engine walks the candidates produced by
// Shrink, descending into the first candidate that also fails, until no
// simpler failing value is left. Numbers halve toward zero, strings peel
// characters from the end, slices and maps shrink structurally.
//
// Shrinking strategies are first-class values in the [shrink] subpackage
// and compose: [shrink.Concat], [shrink.Limit], [shrink.Interleave].
// Any generator's strategy can be replaced with [WithShrinker], for
// example to keep candidates inside the generator range:
//
//	arb := ohsnap.WithShrinker(
//	    ohsnap.ArbitraryInt(rnd, 100, 200),
//	    shrink.Halving(100), // shrink toward the lower bound
//	)
//
// # Logging and budgets
//
// Check is silent. CheckWith accepts [CheckOptions]: log progress every N
// passed cases, trace every tried shrink candidate, limit the total number
// of property calls spent on shrinking, and inject a custom [Logger]
// (satisfied by *testing.T and *testing.B by default):
//
//	ohsnap.CheckWith(t, 1_000_000, arb, prop, ohsnap.CheckOptions{
//	    ProgressEvery:       100_000,
//	    LogShrinkCandidates: true,
//	    Budget:              500,
//	})
//
// # Subpackages
//
// The [shrink] subpackage provides composable shrinking strategies.
// The [github.com/okneniz/oh-snap/json] subpackage generates arbitrary
// JSON documents, useful for testing parsers and serializers.
package ohsnap
