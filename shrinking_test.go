package ohsnap

import (
	"testing"
)

// traceShrinking returns options that log every tried shrink candidate
// via t, making the work of the greedy algorithm visible in test output.
func traceShrinking(t testing.TB) CheckOptions {
	return CheckOptions{
		LogShrinkCandidates: true,
		Logger:              t,
	}
}

func checkShrinking[T comparable](
	t testing.TB,
	arb Arbitrary[T],
	expected T,
	prop func(T) bool,
) {
	t.Helper()

	_, simplestValue := findSimplestBadCaseWith(1000, arb, prop, traceShrinking(t))

	if simplestValue == nil {
		t.Error("bad value not found")
		t.FailNow()
	}

	if *simplestValue != expected {
		t.Errorf("expected: %v", expected)
		t.Errorf("actual: %v", *simplestValue)
		t.FailNow()
	}
}
