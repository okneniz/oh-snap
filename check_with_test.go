package ohsnap

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"strings"
	"testing"
)

type bufferLogger struct {
	buf bytes.Buffer
}

func (b *bufferLogger) Logf(format string, args ...any) {
	fmt.Fprintf(&b.buf, format, args...)
	b.buf.WriteByte('\n')
}

func TestCheckWithShrinkBudget(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(11, 11))
	arb := ArbitraryInt(rnd, 1000, 2000)
	prop := func(int) bool { return false } // every candidate fails

	value, budgeted := findSimplestBadCaseWith(1, arb, prop, CheckOptions{
		Budget:              2,
		LogShrinkCandidates: true,
		Logger:              t,
	})
	if value == nil {
		t.Fatal("expected property to fail")
	}

	// two evaluated candidates: first halving, then halving again
	want := *value / 4
	if *budgeted != want {
		t.Errorf("budgeted shrunk = %d, want %d", *budgeted, want)
	}

	rnd2 := rand.New(rand.NewPCG(11, 11))
	arb2 := ArbitraryInt(rnd2, 1000, 2000)
	_, unlimited := findSimplestBadCaseWith(1, arb2, prop, CheckOptions{
		LogShrinkCandidates: true,
		Logger:              t,
	})
	if *unlimited != 0 {
		t.Errorf("unlimited shrunk = %d, want 0", *unlimited)
	}
}

func TestCheckWithPassingProperty(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(12, 12))
	arb := ArbitraryInt(rnd, 0, 100)

	CheckWith(t, 100, arb, func(x int) bool { return x >= 0 }, CheckOptions{Budget: 10})
}

func TestCheckWithProgressLogging(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(21, 21))
	arb := ArbitraryInt(rnd, 0, 100)

	log := &bufferLogger{}
	CheckWith(t, 1200, arb, func(int) bool { return true }, CheckOptions{
		ProgressEvery: 1000,
		Logger:        log,
	})

	out := log.buf.String()
	if !strings.Contains(out, "1 000 cases passed (") {
		t.Errorf("expected progress message with humanized count, got:\n%s", out)
	}
	if strings.Count(out, "cases passed") != 1 {
		t.Errorf("expected exactly one progress message for 1200 iterations, got:\n%s", out)
	}
}

func TestCheckWithShrinkStepLogging(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(22, 22))
	arb := ArbitraryInt(rnd, 1000, 1000000)
	prop := func(int) bool { return false } // every candidate fails

	log := &bufferLogger{}
	value, shrunk := findSimplestBadCaseWith(10, arb, prop, CheckOptions{
		LogShrinkSteps: true,
		Logger:         log,
	})
	if value == nil {
		t.Fatal("expected property to fail")
	}
	if *shrunk != 0 {
		t.Errorf("shrunk = %d, want 0", *shrunk)
	}

	out := log.buf.String()
	if !strings.Contains(out, ", shrinking...") {
		t.Errorf("expected failure header, got:\n%s", out)
	}
	if !strings.Contains(out, "shrink ") {
		t.Errorf("expected shrink steps, got:\n%s", out)
	}
	if !strings.Contains(out, "minimal counterexample: 0 ") {
		t.Errorf("expected final counterexample, got:\n%s", out)
	}
}

func TestCheckWithBudgetLogging(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewPCG(23, 23))
	arb := ArbitraryInt(rnd, 1000, 2000)
	prop := func(int) bool { return false } // every candidate fails

	log := &bufferLogger{}
	_, shrunk := findSimplestBadCaseWith(1, arb, prop, CheckOptions{
		Budget: 2,
		Logger: log,
	})
	if *shrunk == 0 {
		t.Errorf("shrunk = 0, want a value cut off by the budget")
	}

	out := log.buf.String()
	if !strings.Contains(out, "budget used 2/2") {
		t.Errorf("expected budget summary, got:\n%s", out)
	}
}
