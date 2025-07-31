package ohsnap

import (
	"math/rand"
	"testing"
	"time"
)

func TestAdditionCommutative(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	arb := ArbitraryInt(rnd, 0, 10000)

	prop := func(a int) bool {
		b := arb.Generate()
		return a+b == b+a
	}

	Check(t, arb, prop, DefaultConfig())
}

func TestStringConcatenationAssociative(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	arb := ArbitraryString(rnd, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 3, 10)

	prop := func(a string) bool {
		b := arb.Generate()
		c := arb.Generate()
		return (a+b)+c == a+(b+c)
	}

	Check(t, arb, prop, DefaultConfig())
}

func TestStringLength(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	arb := ArbitraryString(rnd, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 5, 13)

	prop := func(a string) bool {
		b := arb.Generate()
		return len(a+b) == len(a)+len(b)
	}

	Check(t, arb, prop, DefaultConfig())
}
