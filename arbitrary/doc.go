// Package arbitrary provides generators — implementations of
// [ohsnap.Arbitrary] — for primitive types, strings, runes, time and
// containers.
//
// Every constructor takes a *rand.Rand, so generation is deterministic
// for a fixed seed:
//
//	rnd := rand.New(rand.NewPCG(1, 2))
//	arb := arbitrary.Int(rnd, -100, 100)
//
// Combinators that build bigger generators from existing ones —
// Combine, Map, OneOf, Weighted — live in the root package ohsnap.
package arbitrary
