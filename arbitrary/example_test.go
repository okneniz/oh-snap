package arbitrary_test

import (
	"fmt"
	"github.com/okneniz/oh-snap/arbitrary"
	"math/rand/v2"
	"time"
	"unicode"
)

// Generate returns a new random value on every call: loop to pull as
// many values as you need.
func Example_stream() {
	rnd := rand.New(rand.NewPCG(3, 3))
	arb := arbitrary.Int(rnd, 1, 6)

	for range 10 {
		value := arb.Generate()
		fmt.Print(value, " ")
	}

	fmt.Println()
	// Output:
	// 2 3 1 4 4 2 5 3 5 2
}

// Generate returns a single random value.
func Example_singleValue() {
	rnd := rand.New(rand.NewPCG(4, 4))
	dice := arbitrary.Int(rnd, 1, 6)

	fmt.Println("dice:", dice.Generate())
	// Output:
	// dice: 4
}

func ExampleInt() {
	rnd := rand.New(rand.NewPCG(5, 5))
	arb := arbitrary.Int(rnd, -10, 10)

	values := make([]int, 0, 5)
	for range 5 {
		value := arb.Generate()
		values = append(values, value)
	}

	fmt.Println(values)
	// Output:
	// [4 0 -6 -7 -10]
}

func ExampleString() {
	rnd := rand.New(rand.NewPCG(6, 6))
	arb := arbitrary.String(rnd, "ab", 8, 8)

	fmt.Println(arb.Generate())
	// Output:
	// bbabbbba
}

func ExampleTime() {
	rnd := rand.New(rand.NewPCG(7, 7))
	from := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, time.December, 31, 23, 59, 59, 0, time.UTC)
	arb := arbitrary.Time(rnd, from, to)

	moment := arb.Generate()
	fmt.Println(moment.Year(), moment.Month(), moment.Day())
	// Output:
	// 2024 December 17
}

func ExampleSlice() {
	rnd := rand.New(rand.NewPCG(8, 8))
	elem := arbitrary.Int(rnd, 0, 9)
	arb := arbitrary.Slice(rnd, elem, 3, 5)

	fmt.Println(arb.Generate())
	// Output:
	// [2 6 5]
}

// Go prints maps with sorted keys, so the output is deterministic for a seed.
func ExampleMap() {
	rnd := rand.New(rand.NewPCG(9, 9))
	keys := arbitrary.String(rnd, "abc", 1, 1)
	values := arbitrary.Int(rnd, 0, 9)
	arb := arbitrary.Map(rnd, keys, values, 1, 3)

	fmt.Println(arb.Generate())
	// Output:
	// map[a:3 b:2 c:2]
}

// RuneFromTable generates runes from a unicode table.
func ExampleRuneFromTable() {
	rnd := rand.New(rand.NewPCG(15, 15))
	letters := arbitrary.RuneFromTable(rnd, unicode.Letter)

	r := letters.Generate()
	fmt.Printf("U+%04X %q\n", r, r)
	// Output:
	// U+28817 '𨠗'
}
