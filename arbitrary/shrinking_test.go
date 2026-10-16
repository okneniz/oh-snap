package arbitrary

import (
	"testing"

	ohsnap "github.com/okneniz/oh-snap"
)

// traceShrinking returns options that log every tried shrink candidate
// via t, making the work of the greedy algorithm visible in test output.
func traceShrinking(tb testing.TB) ohsnap.CheckOptions {
	tb.Helper()

	return ohsnap.CheckOptions{
		LogShrinkCandidates: true,
		Logger:              tb,
	}
}

func checkShrinking[T comparable](
	tb testing.TB,
	arb ohsnap.Arbitrary[T],
	expected T,
	prop func(T) bool,
) {
	tb.Helper()

	_, simplestValue := ohsnap.FindSimplestBadCase(1000, arb, prop, traceShrinking(tb))

	if simplestValue == nil {
		tb.Error("bad value not found")
		tb.FailNow()
	}

	if *simplestValue != expected {
		tb.Errorf("expected: %v", expected)
		tb.Errorf("actual: %v", *simplestValue)
		tb.FailNow()
	}
}
