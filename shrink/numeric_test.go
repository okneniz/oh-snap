package shrink

import (
	"slices"
	"testing"
)

func TestHalvingTowardZero(t *testing.T) {
	t.Parallel()

	got := slices.Collect(Halving(0)(5))
	want := []int{2, 1, 0}
	if !slices.Equal(got, want) {
		t.Errorf("Halving(0)(5) = %v, want %v", got, want)
	}
}

func TestHalvingNegativeTowardZero(t *testing.T) {
	t.Parallel()

	got := slices.Collect(Halving(0)(-5))
	want := []int{-2, -1, 0}
	if !slices.Equal(got, want) {
		t.Errorf("Halving(0)(-5) = %v, want %v", got, want)
	}
}

func TestHalvingStaysInRange(t *testing.T) {
	t.Parallel()

	got := slices.Collect(Halving(100)(200))
	want := []int{150, 125, 112, 106, 103, 101, 100}
	if !slices.Equal(got, want) {
		t.Errorf("Halving(100)(200) = %v, want %v", got, want)
	}

	for _, candidate := range got {
		if candidate < 100 {
			t.Errorf("candidate %d is out of range [100, 200]", candidate)
		}
	}
}

func TestHalvingUnsignedUpward(t *testing.T) {
	t.Parallel()

	got := slices.Collect(Halving[uint](10)(3))
	want := []uint{7, 9, 10}
	if !slices.Equal(got, want) {
		t.Errorf("Halving[uint](10)(3) = %v, want %v", got, want)
	}
}

func TestBoundaries(t *testing.T) {
	t.Parallel()

	got := slices.Collect(Boundaries(0, 1, -1)(42))
	want := []int{0, 1, -1}
	if !slices.Equal(got, want) {
		t.Errorf("Boundaries(0, 1, -1)(42) = %v, want %v", got, want)
	}
}

func TestBoundariesSkipsInputValue(t *testing.T) {
	t.Parallel()

	got := slices.Collect(Boundaries(0, 1, 2)(1))
	want := []int{0, 2}
	if !slices.Equal(got, want) {
		t.Errorf("Boundaries(0, 1, 2)(1) = %v, want %v", got, want)
	}
}
