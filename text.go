package frenchspace

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	spaceRunes     = regexp.MustCompile(` ?[;!?%»›”]`)
	quotationMark  = regexp.MustCompile(`[«‹“] ?`)
	colon          = regexp.MustCompile(` ?:`)
	multipleSpaces = regexp.MustCompile(`\s+`)
)

// Add space (thin and non-breaking) before or after ponctuation symbols.
func Text(s string) string {
	s = spaces(s)
	s = spaceRunes.ReplaceAllStringFunc(s, func(s string) string {
		return "\u202F" + strings.TrimSpace(s)
	})
	s = quotationMark.ReplaceAllStringFunc(s, func(s string) string {
		return strings.TrimSpace(s) + "\u202F"
	})
	s = colon.ReplaceAllStringFunc(s, func(s string) string {
		return "\u00A0" + strings.TrimSpace(s)
	})
	return s
}

// Remove spaces at the begin and at the end of the string; then replace all
// space runes series by one space.
func spaces(s string) string {
	s = strings.TrimSpace(s)
	buff := strings.Builder{}
	buff.Grow(len(s))

	beforeSpace := false
	for _, r := range s {
		if unicode.IsSpace(r) {
			beforeSpace = true
		} else if beforeSpace {
			beforeSpace = false
			buff.WriteByte(' ')
			buff.WriteRune(r)
		} else {
			buff.WriteRune(r)
		}
	}
	return buff.String()
}
