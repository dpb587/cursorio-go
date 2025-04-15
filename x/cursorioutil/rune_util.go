package cursorioutil

import "unicode/utf8"

func RunesBytes(runes []rune) int {
	if runes == nil {
		return 0
	}

	var l int

	for _, r := range runes {
		l += utf8.RuneLen(r)
	}

	return l
}
