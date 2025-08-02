package ohsnap

import (
	"math/rand/v2"
	"testing"
	"time"
)

func TestMapped(t *testing.T) {
	t.Parallel()

	const iterations = 100000

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	type User struct {
		Name string
		Date time.Time
	}

	names := ArbitraryString(
		rnd,
		"abcdefgjihijklmnoprstwuxyz",
		3,
		10,
	)

	year := 365 * 24 * time.Hour

	dates := ArbitraryTime(
		rnd,
		time.Now().Add(-40*year),
		time.Now().Add(-18*year),
	)

	usersData := Combine(names, dates)

	users := Map(
		usersData,
		func(p Pair[string, time.Time]) User {
			return User{
				Name: p.First,
				Date: p.Second,
			}
		},
	)

	CheckM(t, 10, users, func(p User) bool {
		t.Log("user", p)
		return true
	})
}
