package ohsnap

import (
	"math/rand"
)

type stringArbitrary struct {
	rand     *rand.Rand
	letters  string
	from, to int
}

func ArbitraryString(rnd *rand.Rand, letters string, from, to int) Arbitrary[string] {
	return &stringArbitrary{
		rand:    rnd,
		letters: letters,
		from:    from,
		to:      to,
	}
}

func (a stringArbitrary) Generate() string {
	length := a.rand.Intn(a.to-a.from+1) + int(a.from)
	result := make([]byte, length)

	for i := range result {
		result[i] = a.letters[rand.Intn(len(a.letters))]
	}

	return string(result)
}

func (stringArbitrary) Shrink(value string) []string {
	var results []string

	for len(value) > 0 {
		value = value[:len(value)-1] // Remove the last character
		results = append(results, value)
	}

	return results
}
