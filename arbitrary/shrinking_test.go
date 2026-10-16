package arbitrary

import (
	"testing"

	ohsnap "github.com/okneniz/oh-snap"
)

// traceShrinking returns options that log every tried shrink candidate
// via t, making the work of the greedy algorithm visible in test output.
func traceShrinking(t testing.TB) ohsnap.CheckOptions {
	return ohsnap.CheckOptions{
		LogShrinkCandidates: true,
		Logger:              t,
	}
}

func checkShrinking[T comparable](
	t testing.TB,
	arb ohsnap.Arbitrary[T],
	expected T,
	prop func(T) bool,
) {
	t.Helper()

	_, simplestValue := ohsnap.FindSimplestBadCase(1000, arb, prop, traceShrinking(t))

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
