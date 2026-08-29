# OhSnap!

![Downloads](https://img.shields.io/github/downloads/okneniz/oh-snap/total) ![Contributors](https://img.shields.io/github/contributors/okneniz/oh-snap?color=dark-green) ![Forks](https://img.shields.io/github/forks/okneniz/oh-snap?style=social) ![Stargazers](https://img.shields.io/github/stars/okneniz/oh-snap?style=social) ![Issues](https://img.shields.io/github/issues/okneniz/oh-snap) ![License](https://img.shields.io/github/license/okneniz/oh-snap)

Simplest property-based testing library for Go.

Instead of writing individual test cases you declare a **property** that must
hold for any input. OhSnap feeds it random values and, when it finds a
counterexample, **shrinks** it to the simplest form that still fails — so
failures are obvious at a glance.

```go
func TestDoubleReverseIsIdentity(t *testing.T) {
	rnd := rand.New(rand.NewPCG(1, 2))
	arb := ohsnap.ArbitrarySlice(rnd, ohsnap.ArbitraryInt(rnd, 0, 100), 0, 10)

	ohsnap.Check(t, 1000, arb, func(xs []int) bool {
		dst := slices.Clone(xs)
		slices.Reverse(dst)
		slices.Reverse(dst)
		return slices.Equal(xs, dst)
	})
}
```

## Installation

```bash
go get github.com/okneniz/oh-snap
```

## Why OhSnap

- **Lazy by design.** `Generate` returns an `iter.Seq` — an infinite stream
  of random values. Pull one with `ohsnap.First`, or range over it and take
  as many as you need.
- **Real shrinking.** Counterexamples are minimized automatically: numbers
  halve toward zero, strings peel characters, slices and maps shrink
  structurally.
- **Shrinking as a strategy.** The [`shrink`](https://pkg.go.dev/github.com/okneniz/oh-snap/shrink)
  subpackage makes strategies first-class, composable values
  (`Halving`, `Boundaries`, `Peel`, `Concat`, `Interleave`, ...) and
  `WithShrinker` plugs any of them into any generator.
- **Readable diagnostics.** `CheckWith` logs progress every N cases and can
  trace every shrink candidate, showing exactly how the algorithm works.
- **Zero dependencies**, stdlib only (Go 1.23+).

## Tour

### Generators

```go
ohsnap.ArbitraryInt(rnd, -100, 100)                    // all numeric types
ohsnap.ArbitraryString(rnd, "abcdef", 1, 10)           // strings
ohsnap.ArbitraryTime(rnd, from, to)                    // time.Time
ohsnap.ArbitrarySlice(rnd, elem, 0, 10)                // []T
ohsnap.ArbitraryMap(rnd, keys, values, 0, 10)          // map[K]V
ohsnap.RuneFromTable(rnd, unicode.Letter)              // unicode tables
```

Combinators build bigger generators from smaller ones:

```go
pairs   := ohsnap.Combine(arbInt, arbBool)             // Pair[T, U]
words   := ohsnap.Map(arbInt, func(n int) string {...}) // transform
pick    := ohsnap.OneOf(rnd, []ohsnap.Arbitrary[int]{small, large})
choice  := ohsnap.OneOfValue(rnd, "red", "green", "blue")
biased  := ohsnap.Weighted(rnd, map[int]ohsnap.Arbitrary[int]{
	1: small, // "small" ~25% of the time
	3: large, // "large" ~75% of the time
})
```

The `Builder` configures bounds fluently:

```go
b := ohsnap.NewBuilder(rnd).MinInt(10).MaxInt(20).MaxSliceLen(5)
ints := b.Int()
slicesOfInts := b.IntSlice()
```

### Shrinking strategies

```go
// keep candidates inside [100, 200] and walk toward its lower bound:
arb := ohsnap.WithShrinker(
	ohsnap.ArbitraryInt(rnd, 100, 200),
	shrink.Halving(100),
)
```

```
shrink 114 → 107 (fails)
shrink 107 → 103 (fails)
shrink 103 → 101 (fails)
shrink 101 → 100 (fails)
minimal counterexample: 100 (from 114, 4 steps)
```

Custom strategies are plain functions over sequences:

```go
countdown := shrink.Shrinker[int](func(v int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for c := v - 1; c > 0; c-- {
			if !yield(c) {
				return
			}
		}
	}
})
```

### Logging and budgets

```go
ohsnap.CheckWith(t, 1_000_000, arb, prop, ohsnap.CheckOptions{
	ProgressEvery:       100_000, // "100 000 cases passed (124ms)"
	LogShrinkCandidates: true,    // full trace of the greedy search
	Budget:              500,     // cap property calls spent on shrinking
	Logger:              myLogger, // optional; *testing.T by default
})
```

### Arbitrary JSON

```go
arb := json.ArbitraryJSON(rnd, 4, 6) // depth ≤ 4, collections ≤ 6

ohsnap.Check(t, 1000, arb, func(v json.Value) bool {
	data, err := encodingjson.Marshal(v)
	return err == nil && encodingjson.Valid(data)
})
```

## Documentation

- [GoDoc: ohsnap](https://pkg.go.dev/github.com/okneniz/oh-snap) — package
  overview and runnable examples for every generator and combinator
- [GoDoc: shrink](https://pkg.go.dev/github.com/okneniz/oh-snap/shrink) —
  shrinking strategies
- [GoDoc: json](https://pkg.go.dev/github.com/okneniz/oh-snap/json) —
  arbitrary JSON documents

External examples:

- [assembly](https://github.com/okneniz/assembly) — low-level machine-code
  toolkit (ARM64 / RISC-V / LoongArch assemblers, disassemblers, ELF and
  Mach-O parsers) with property tests built on oh-snap
- [timestamps parsing](https://github.com/okneniz/parsec/blob/master/examples/strings/timestamps/timestamps_test.go)
- [json parsing](https://github.com/okneniz/parsec/blob/master/examples/strings/json/json_test.go)

## Roadmap

See the [open issues](https://github.com/okneniz/oh-snap/issues) for a list
of proposed features (and known issues).

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more
information.

## Contributing

Contributions are what make the open source community such an amazing place
to be learn, inspire, and create. Any contributions you make are **greatly
appreciated**.

* If you have suggestions for adding or removing projects, feel free to
  [open an issue](https://github.com/okneniz/oh-snap/issues/new) to discuss
  it, or directly create a pull request after you edit the *README.md* file
  with necessary changes.
* Please make sure you check your spelling and grammar.
* Create individual PR for each suggestion.

### Creating A Pull Request

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
