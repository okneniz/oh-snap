package ohsnap

import (
	"iter"
	"testing"
	"time"
)

// Arbitrary is an interface for generating random values and shrinking them.
// Generate returns a lazy, effectively infinite sequence of random values.
// Shrink returns a lazy sequence of candidates simpler than the input value.
type Arbitrary[T any] interface {
	Generate() iter.Seq[T]
	Shrink(T) iter.Seq[T]
}

// Property is a function that takes a value of type T and returns a boolean.
type Property[T any] func(T) bool

// CheckOptions configures CheckWith.
type CheckOptions struct {
	// Budget limits the total number of property calls made during shrinking.
	// When the budget is exhausted, shrinking stops and returns the best
	// value found so far. Zero or negative means unlimited.
	Budget int

	// ProgressEvery logs a message after every N passed cases during
	// the generation phase. Zero means disabled.
	ProgressEvery int

	// LogShrinkSteps logs every accepted shrinking step: a candidate that
	// also fails the property and becomes the new minimum.
	// Disabled by default.
	LogShrinkSteps bool

	// LogShrinkCandidates logs every tried shrink candidate with its
	// outcome, including rejected ones. It shows the full work of the
	// greedy algorithm and includes the LogShrinkSteps output.
	// Disabled by default.
	LogShrinkCandidates bool

	// Logger receives the messages. When nil, the *testing.T passed
	// to CheckWith is used.
	Logger Logger
}

// Check runs property-based tests with random values and shrinking.
// It is silent: use CheckWith to enable progress and shrink logging.
func Check[T any](t *testing.T, iterations int, arb Arbitrary[T], prop Property[T]) {
	value, shrunk := findSimplestBadCase(iterations, arb, prop)
	reportFailure(t, value, shrunk)
}

// CheckWith runs property-based tests with custom options.
// When opts.Logger is nil, the *testing.T is used for logging.
func CheckWith[T any](t *testing.T, iterations int, arb Arbitrary[T], prop Property[T], opts CheckOptions) {
	if opts.Logger == nil {
		opts.Logger = t
	}

	value, shrunk := findSimplestBadCaseWith(iterations, arb, prop, opts)
	reportFailure(t, value, shrunk)
}

func reportFailure[T any](t *testing.T, value, shrunk *T) {
	if value != nil {
		t.Errorf("Property failed for value: %v (shrunk: %v)", *value, *shrunk)
	}
}

// findSimplestBadCase find simplest bad case of input value for property func
func findSimplestBadCase[T any](iterations int, arb Arbitrary[T], prop Property[T]) (*T, *T) {
	return findSimplestBadCaseWith(iterations, arb, prop, CheckOptions{})
}

func findSimplestBadCaseWith[T any](iterations int, arb Arbitrary[T], prop Property[T], opts CheckOptions) (*T, *T) {
	if opts.Logger == nil {
		opts.Logger = nopLogger{}
	}

	if iterations <= 0 {
		return nil, nil
	}

	start := time.Now()

	i := 0
	passed := 0
	for value := range arb.Generate() {
		if i >= iterations {
			break
		}
		i++

		if !prop(value) {
			opts.Logger.Logf("property failed for %v, shrinking...", value)

			st := &shrinkState{
				logger:        opts.Logger,
				logSteps:      opts.LogShrinkSteps,
				logCandidates: opts.LogShrinkCandidates,
				unlimited:     opts.Budget <= 0,
				budget:        opts.Budget,
			}
			shrunk := shrinkValue(arb, value, prop, st)

			if st.unlimited {
				opts.Logger.Logf("minimal counterexample: %v (from %v, %d steps)", shrunk, value, st.steps)
			} else {
				used := opts.Budget - st.budget
				opts.Logger.Logf("minimal counterexample: %v (from %v, %d steps, budget used %d/%d)", shrunk, value, st.steps, used, opts.Budget)
			}

			return &value, &shrunk
		}

		passed++
		if opts.ProgressEvery > 0 && passed%opts.ProgressEvery == 0 {
			elapsed := time.Since(start).Round(time.Millisecond)
			opts.Logger.Logf("%s cases passed (%v)", humanize(passed), elapsed)
		}
	}

	return nil, nil
}

// shrinkState carries shrinking phase configuration and accounting.
type shrinkState struct {
	logger        Logger
	logSteps      bool // log accepted descent steps
	logCandidates bool // log every tried candidate, including rejected
	unlimited     bool
	budget        int
	steps         int
}

// shrinkValue attempts to shrink a failing value to its minimal form.
// When the budget is exhausted it returns the best value found so far.
func shrinkValue[T any](arb Arbitrary[T], value T, prop Property[T], st *shrinkState) T {
	for smaller := range arb.Shrink(value) {
		if !st.unlimited {
			if st.budget <= 0 {
				return value
			}
			st.budget--
		}

		fails := !prop(smaller)

		if st.logCandidates {
			if fails {
				st.logger.Logf("shrink %v → %v (fails)", value, smaller)
			} else {
				st.logger.Logf("shrink %v → %v (holds)", value, smaller)
			}
		}

		if fails {
			st.steps++
			if st.logSteps && !st.logCandidates {
				st.logger.Logf("shrink %v → %v", value, smaller)
			}
			return shrinkValue(arb, smaller, prop, st)
		}
	}

	return value
}
