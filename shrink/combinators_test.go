package shrink

import (
	"slices"
	"testing"
)

func TestConcat(t *testing.T) {
	t.Parallel()

	s := Concat(Halving(0), Boundaries(7))

	got := slices.Collect(s(5))
	want := []int{2, 1, 0, 7}
	if !slices.Equal(got, want) {
		t.Errorf("Concat(Halving, Boundaries)(5) = %v, want %v", got, want)
	}
}

func TestLimit(t *testing.T) {
	t.Parallel()

	s := Limit(2, Halving(0))

	got := slices.Collect(s(5))
	want := []int{2, 1}
	if !slices.Equal(got, want) {
		t.Errorf("Limit(2, Halving)(5) = %v, want %v", got, want)
	}
}

func TestInterleave(t *testing.T) {
	t.Parallel()

	s := Interleave(Boundaries(1, 2), Boundaries(3, 4))

	got := slices.Collect(s(9))
	want := []int{1, 3, 2, 4}
	if !slices.Equal(got, want) {
		t.Errorf("Interleave(Boundaries, Boundaries)(9) = %v, want %v", got, want)
	}
}

func TestNoShrink(t *testing.T) {
	t.Parallel()

	got := slices.Collect(NoShrink[int]()(5))
	if len(got) != 0 {
		t.Errorf("NoShrink[int]()(5) = %v, want no candidates", got)
	}
}
