package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {

	for _, logRune := range log {
		if logRune == '❗' {
			return "recommendation"
		}
		if logRune == '🔍' {
			return "search"
		}

		if logRune == '☀' {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := make([]rune, 0, utf8.RuneCountInString(log))
	for _, logRune := range log {
		if logRune == oldRune {
			runes = append(runes, newRune)
		} else {
			runes = append(runes, logRune)
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.

func WithinLimit(log string, limit int) bool {
	return (utf8.RuneCountInString(log) <= limit)
}
