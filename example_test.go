package ohsnap_test

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/arbitrary"
	"github.com/okneniz/oh-snap/shrink"
)

// Check needs a *testing.T; inside real tests you pass the one you got.
// The property below holds for every generated value, so nothing is reported.
func ExampleCheck() {
	rnd := rand.New(rand.NewPCG(1, 1))
	arb := arbitrary.Int(rnd, 0, 1000)

	ohsnap.Check(new(testing.T), 1000, arb, func(x int) bool {
		return x >= 0 && x <= 1000
	})
	// Output:
}

// clockFreeLogger drops the elapsed time suffix so the example output is stable.
type clockFreeLogger struct{ lines []string }

func (l *clockFreeLogger) Logf(format string, args ...any) {
	line := fmt.Sprintf(format, args...)
	if i := strings.Index(line, " ("); i >= 0 {
		line = line[:i]
	}
	l.lines = append(l.lines, line)
}

func ExampleCheckWith_progress() {
	rnd := rand.New(rand.NewPCG(2, 2))
	arb := arbitrary.Int(rnd, 0, 100)
	logger := &clockFreeLogger{}

	ohsnap.CheckWith(new(testing.T), 500, arb, func(int) bool { return true }, ohsnap.CheckOptions{
		ProgressEvery: 250,
		Logger:        logger,
	})

	for _, line := range logger.lines {
		fmt.Println(line)
	}
	// Output:
	// 250 cases passed
	// 500 cases passed
}

// Combine pairs two generators; the pair shrinks component-wise.
func ExampleCombine() {
	rnd := rand.New(rand.NewPCG(10, 10))
	ages := arbitrary.Int(rnd, 18, 99)
	active := arbitrary.Bool(rnd)
	users := ohsnap.Combine(ages, active)

	user := users.Generate()
	fmt.Printf("age=%d active=%v\n", user.First, user.Second)
	// Output:
	// age=87 active=true
}

// Map transforms generated values with a plain function.
func ExampleMap() {
	rnd := rand.New(rand.NewPCG(11, 11))
	times := arbitrary.Int(rnd, 1, 3)
	words := ohsnap.Map(times, func(n int) string { return strings.Repeat("ha", n) })

	fmt.Println(words.Generate() + "!")
	// Output:
	// ha!
}

// OneOfValue picks from a fixed list of values.
func ExampleOneOfValue() {
	rnd := rand.New(rand.NewPCG(12, 12))
	colors := ohsnap.OneOfValue(rnd, "red", "green", "blue")

	fmt.Println(colors.Generate())
	// Output:
	// red
}

// OneOf picks a nested generator on every pull.
func ExampleOneOf() {
	rnd := rand.New(rand.NewPCG(13, 13))
	small := arbitrary.Int(rnd, 1, 6)
	large := arbitrary.Int(rnd, 100, 200)
	die := ohsnap.OneOf(rnd, small, large)

	fmt.Println(die.Generate())
	// Output:
	// 1
}

// Weighted picks nested generators with the given weights; here "b" is
// expected three times as often as "a".
func ExampleWeighted() {
	rnd := rand.New(rand.NewPCG(14, 14))
	arb := ohsnap.Weighted[string](rnd, map[int]ohsnap.Arbitrary[string]{
		1: ohsnap.OneOfValue(rnd, "a"),
		3: ohsnap.OneOfValue(rnd, "b"),
	})

	counts := map[string]int{}
	for i := 0; i < 1000; i++ {
		counts[arb.Generate()]++
	}
	fmt.Println(counts)
	// Output:
	// map[a:234 b:766]
}

// Shrink returns candidates simpler than the value: ints halve toward zero.
func Example_shrinking() {
	rnd := rand.New(rand.NewPCG(17, 17))
	arb := arbitrary.Int(rnd, 0, 1000)

	fmt.Println(slices.Collect(arb.Shrink(37)))
	// Output:
	// [18 9 4 2 1 0]
}

// WithShrinker replaces the shrinking strategy of any generator.
// Halving(100) keeps candidates inside the range and walks toward its
// lower bound, unlike the default halving toward zero.
func ExampleWithShrinker() {
	rnd := rand.New(rand.NewPCG(18, 18))
	arb := ohsnap.WithShrinker(
		arbitrary.Int(rnd, 100, 200),
		shrink.Halving(100),
	)

	for candidate := range arb.Shrink(180) {
		fmt.Print(candidate, " ")
	}
	fmt.Println()
	// Output:
	// 140 120 110 105 102 101 100
}
