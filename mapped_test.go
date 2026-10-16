package ohsnap_test

import (
	"math/rand/v2"
	"testing"
	"time"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/arbitrary"
)

func TestMapped(t *testing.T) {
	t.Parallel()

	const iterations = 100

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)

	rnd := rand.New(rand.NewPCG(0, uint64(seed)))

	type User struct {
		Name string
		Date time.Time
	}

	names := arbitrary.String(
		rnd,
		"abcdefgjihijklmnoprstwuxyz",
		3,
		10,
	)

	year := 365 * 24 * time.Hour

	dates := arbitrary.Time(
		rnd,
		time.Now().Add(-40*year),
		time.Now().Add(-18*year),
	)

	usersData := ohsnap.Combine(names, dates)

	users := ohsnap.Map(
		usersData,
		func(p ohsnap.Pair[string, time.Time]) User {
			return User{
				Name: p.First,
				Date: p.Second,
			}
		},
	)

	ohsnap.Check(t, iterations, users, func(p User) bool {
		t.Log("user", p)
		return true
	})
}
