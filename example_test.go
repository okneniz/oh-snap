package ohsnap_test

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"
	"time"
	"unicode"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

// Check needs a *testing.T; inside real tests you pass the one you got.
// The property below holds for every generated value, so nothing is reported.
func ExampleCheck() {
	rnd := rand.New(rand.NewPCG(1, 1))
	arb := ohsnap.ArbitraryInt(rnd, 0, 1000)

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
	arb := ohsnap.ArbitraryInt(rnd, 0, 100)
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

// Generate is a lazy, effectively infinite stream of random values:
// pull as many as you need with a range loop.
func Example_stream() {
	rnd := rand.New(rand.NewPCG(3, 3))
	arb := ohsnap.ArbitraryInt(rnd, 1, 6)

	i := 0
	for value := range arb.Generate() {
		fmt.Print(value, " ")
		i++
		if i == 10 {
			break
		}
	}
	fmt.Println()
	// Output:
	// 2 3 1 4 4 2 5 3 5 2
}

// First pulls a single value from the stream.
func ExampleFirst() {
	rnd := rand.New(rand.NewPCG(4, 4))
	dice := ohsnap.ArbitraryInt(rnd, 1, 6)

	fmt.Println("dice:", ohsnap.First(dice.Generate()))
	// Output:
	// dice: 4
}

func ExampleArbitraryInt() {
	rnd := rand.New(rand.NewPCG(5, 5))
	arb := ohsnap.ArbitraryInt(rnd, -10, 10)

	values := []int{}
	for value := range arb.Generate() {
		values = append(values, value)
		if len(values) == 5 {
			break
		}
	}
	fmt.Println(values)
	// Output:
	// [4 0 -6 -7 -10]
}

func ExampleArbitraryString() {
	rnd := rand.New(rand.NewPCG(6, 6))
	arb := ohsnap.ArbitraryString(rnd, "ab", 8, 8)

	fmt.Println(ohsnap.First(arb.Generate()))
	// Output:
	// bbabbbba
}

func ExampleArbitraryTime() {
	rnd := rand.New(rand.NewPCG(7, 7))
	from := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, time.December, 31, 23, 59, 59, 0, time.UTC)
	arb := ohsnap.ArbitraryTime(rnd, from, to)

	moment := ohsnap.First(arb.Generate())
	fmt.Println(moment.Year(), moment.Month(), moment.Day())
	// Output:
	// 2024 December 17
}

func ExampleArbitrarySlice() {
	rnd := rand.New(rand.NewPCG(8, 8))
	elem := ohsnap.ArbitraryInt(rnd, 0, 9)
	arb := ohsnap.ArbitrarySlice(rnd, elem, 3, 5)

	fmt.Println(ohsnap.First(arb.Generate()))
	// Output:
	// [2 6 5]
}

// Go prints maps with sorted keys, so the output is deterministic for a seed.
func ExampleArbitraryMap() {
	rnd := rand.New(rand.NewPCG(9, 9))
	keys := ohsnap.ArbitraryString(rnd, "abc", 1, 1)
	values := ohsnap.ArbitraryInt(rnd, 0, 9)
	arb := ohsnap.ArbitraryMap(rnd, keys, values, 1, 3)

	fmt.Println(ohsnap.First(arb.Generate()))
	// Output:
	// map[a:3 b:2 c:2]
}

// Combine pairs two generators; the pair shrinks component-wise.
func ExampleCombine() {
	rnd := rand.New(rand.NewPCG(10, 10))
	ages := ohsnap.ArbitraryInt(rnd, 18, 99)
	active := ohsnap.ArbitraryBool(rnd)
	users := ohsnap.Combine(ages, active)

	user := ohsnap.First(users.Generate())
	fmt.Printf("age=%d active=%v\n", user.First, user.Second)
	// Output:
	// age=87 active=true
}

// Map transforms generated values with a plain function.
func ExampleMap() {
	rnd := rand.New(rand.NewPCG(11, 11))
	times := ohsnap.ArbitraryInt(rnd, 1, 3)
	words := ohsnap.Map(times, func(n int) string { return strings.Repeat("ha", n) })

	fmt.Println(ohsnap.First(words.Generate()) + "!")
	// Output:
	// ha!
}

// OneOfValue picks from a fixed list of values.
func ExampleOneOfValue() {
	rnd := rand.New(rand.NewPCG(12, 12))
	colors := ohsnap.OneOfValue(rnd, "red", "green", "blue")

	fmt.Println(ohsnap.First(colors.Generate()))
	// Output:
	// red
}

// OneOf picks a nested generator on every pull.
func ExampleOneOf() {
	rnd := rand.New(rand.NewPCG(13, 13))
	small := ohsnap.ArbitraryInt(rnd, 1, 6)
	large := ohsnap.ArbitraryInt(rnd, 100, 200)
	die := ohsnap.OneOf(rnd, []ohsnap.Arbitrary[int]{small, large})

	fmt.Println(ohsnap.First(die.Generate()))
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
		counts[ohsnap.First(arb.Generate())]++
	}
	fmt.Println(counts)
	// Output:
	// map[a:234 b:766]
}

// RuneFromTable generates runes from a unicode table.
func ExampleRuneFromTable() {
	rnd := rand.New(rand.NewPCG(15, 15))
	letters := ohsnap.RuneFromTable(rnd, unicode.Letter)

	r := ohsnap.First(letters.Generate())
	fmt.Printf("U+%04X %q\n", r, r)
	// Output:
	// U+28817 '𨠗'
}

// The Builder configures bounds fluently and produces any generator.
func ExampleBuilder() {
	rnd := rand.New(rand.NewPCG(16, 16))
	b := ohsnap.NewBuilder(rnd).
		MinInt(10).MaxInt(20).
		MinSliceLen(2).MaxSliceLen(4)

	fmt.Println(ohsnap.First(b.Int().Generate()))
	fmt.Println(ohsnap.First(b.IntSlice().Generate()))
	// Output:
	// 11
	// [17 12 15]
}

// Shrink returns candidates simpler than the value: ints halve toward zero.
func Example_shrinking() {
	rnd := rand.New(rand.NewPCG(17, 17))
	arb := ohsnap.ArbitraryInt(rnd, 0, 1000)

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
		ohsnap.ArbitraryInt(rnd, 100, 200),
		shrink.Halving(100),
	)

	for candidate := range arb.Shrink(180) {
		fmt.Print(candidate, " ")
	}
	fmt.Println()
	// Output:
	// 140 120 110 105 102 101 100
}
