package ohsnap

import (
	"strconv"
	"strings"
)

// Logger receives progress and shrinking messages during CheckWith.
// It is satisfied by *testing.T, *testing.B and any custom logger.
type Logger interface {
	Logf(format string, args ...any)
}

// nopLogger drops all messages. It is the default logger for Check
// and for direct findSimplestBadCase calls.
type nopLogger struct{}

func (nopLogger) Logf(string, ...any) {}

// humanize groups digits of a non-negative number with spaces: 100000 -> "100 000".
func humanize(n int) string {
	digits := strconv.Itoa(n)
	var b strings.Builder
	for i := 0; i < len(digits); i++ {
		if i > 0 && (len(digits)-i)%3 == 0 {
			b.WriteByte(' ')
		}
		b.WriteByte(digits[i])
	}
	return b.String()
}
