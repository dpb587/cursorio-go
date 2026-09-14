package cursorio

import "unicode/utf8"

type DecodedRune struct {
	Size int
	Rune rune
}

func (dr DecodedRune) AsDecodedRunes() DecodedRunes {
	return DecodedRunes{
		Size:  dr.Size,
		Runes: []rune{dr.Rune},
	}
}

//

type DecodedRuneList []DecodedRune

func (dr DecodedRuneList) AsDecodedRunes() DecodedRunes {
	return NewDecodedRunes(dr...)
}

func (dr DecodedRuneList) String() string {
	if len(dr) == 0 {
		return ""
	}

	byteSize := 0

	for _, r := range dr {
		byteSize += r.Size
	}

	b := make([]byte, 0, byteSize)

	for _, r := range dr {
		b = utf8.AppendRune(b, r.Rune)
	}

	return string(b)
}

//

type DecodedRunes struct {
	Size  int
	Runes []rune
}

func NewDecodedRunes(rl ...DecodedRune) DecodedRunes {
	dr := DecodedRunes{
		Size:  0,
		Runes: make([]rune, len(rl)),
	}

	for i, r := range rl {
		dr.Size += r.Size
		dr.Runes[i] = r.Rune
	}

	return dr
}

func (dr DecodedRunes) Append(rl ...DecodedRune) DecodedRunes {
	if len(rl) == 0 {
		return dr
	}

	ndr := DecodedRunes{
		Size:  dr.Size,
		Runes: make([]rune, 0, len(dr.Runes)+len(rl)),
	}

	ndr.Runes = append(ndr.Runes, dr.Runes...)

	for _, r := range rl {
		ndr.Size += r.Size
		ndr.Runes = append(ndr.Runes, r.Rune)
	}

	return ndr
}

func (dr DecodedRunes) String() string {
	return string(dr.Runes)
}
