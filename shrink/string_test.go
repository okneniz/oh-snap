package shrink

import (
	"slices"
	"testing"
)

func TestPeel(t *testing.T) {
	t.Parallel()

	got := slices.Collect(Peel()("abc"))
	want := []string{"ab", "a", ""}
	if !slices.Equal(got, want) {
		t.Errorf("Peel()(abc) = %v, want %v", got, want)
	}
}

func TestHalveLength(t *testing.T) {
	t.Parallel()

	got := slices.Collect(HalveLength()("abcdef"))
	want := []string{"abc", "a", ""}
	if !slices.Equal(got, want) {
		t.Errorf("HalveLength()(abcdef) = %v, want %v", got, want)
	}
}

func TestMinimizeChars(t *testing.T) {
	t.Parallel()

	got := slices.Collect(MinimizeChars('a')("cba"))
	want := []string{"aba", "caa"}
	if !slices.Equal(got, want) {
		t.Errorf("MinimizeChars(a)(cba) = %v, want %v", got, want)
	}
}
