package ohsnap

import (
	"testing"
)

func checkShrinking[T comparable](
	t testing.TB,
	arb Arbitrary[T],
	expected T,
	prop func(T) bool,
) {
	t.Helper()

	_, simplestValue := findSimplestBadCase(1000, arb, prop)

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
