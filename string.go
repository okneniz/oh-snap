package ohsnap

import (
	"math/rand/v2"
)

type arbitrationString struct {
	rand     *rand.Rand
	letters  string
	from, to int
}

// ArbitraryString - return generator for arbitrary strings.
// rnd - pseudo-random number generator.
// letters - string with allowed runes.
// from and to - bounds of length of strings.
func ArbitraryString(rnd *rand.Rand, letters string, from, to int) Arbitrary[string] {
	if from > to {
		from, to = to, from
	}

	return &arbitrationString{
		rand:    rnd,
		letters: letters,
		from:    from,
		to:      to,
	}
}

func (a arbitrationString) Generate() string {
	length := a.rand.IntN(a.to-a.from+1) + int(a.from)
	result := make([]byte, length)

	for i := range result {
		result[i] = a.letters[rand.IntN(len(a.letters))]
	}

	return string(result)
}

func (arbitrationString) Shrink(value string) []string {
	var results []string

	for len(value) > 0 {
		value = value[:len(value)-1] // Remove the last character
		results = append(results, value)
	}

	return results
}
