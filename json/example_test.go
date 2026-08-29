package json_test

import (
	stdjson "encoding/json"
	"fmt"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	ohjson "github.com/okneniz/oh-snap/json"
)

// ArbitraryJSON generates random JSON documents with bounded nesting depth
// and collection size. A typical use is checking parsers and serializers
// with properties instead of fixed fixtures.
func ExampleArbitraryJSON() {
	rnd := rand.New(rand.NewPCG(7, 7))
	arb := ohjson.ArbitraryJSON(rnd, 3, 5)

	const documents = 100
	roundtripped := 0

	for i := 0; i < documents; i++ {
		value := ohsnap.First(arb.Generate())

		data, err := stdjson.Marshal(value)
		if err != nil {
			panic(err)
		}

		var reparsed any
		if err := stdjson.Unmarshal(data, &reparsed); err != nil {
			panic(err)
		}

		rewrapped, err := stdjson.Marshal(reparsed)
		if err != nil {
			panic(err)
		}

		if string(data) == string(rewrapped) {
			roundtripped++
		}
	}

	fmt.Printf("%d of %d documents survived a JSON round-trip\n", roundtripped, documents)
	// Output:
	// 100 of 100 documents survived a JSON round-trip
}
