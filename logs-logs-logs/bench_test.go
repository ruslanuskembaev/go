package logs

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// Вариант 1: += string(...) (медленный)
func replaceSlow(log string, oldRune, newRune rune) string {
	var newLog string
	for _, logRune := range log {
		if logRune == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(logRune)
		}
	}
	return newLog
}

// Вариант 2: []rune + append (эффективный)
func replaceFast(log string, oldRune, newRune rune) string {
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

// Подготовим длинную строку для честного сравнения
var longLog = strings.Repeat("exercism❗", 50000)

func BenchmarkReplaceSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = replaceSlow(longLog, '❗', '!')
	}
}

func BenchmarkReplaceFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = replaceFast(longLog, '❗', '!')
	}
}
