package cursorio

import (
	"errors"
	"fmt"
	"strings"
)

type TextOffsetRange struct {
	From  TextOffset
	Until TextOffset
}

var _ OffsetRange = TextOffsetRange{}

func ParseTextOffsetRange(v string) (TextOffsetRange, error) {
	fields := strings.SplitN(v, ";", 2)
	if len(fields) != 2 {
		return TextOffsetRange{}, errors.New("invalid text cursor range")
	}

	textLineColumnRange, err := ParseTextLineColumnRange(fields[0])
	if err != nil {
		return TextOffsetRange{}, fmt.Errorf("text range: %v", err)
	}

	byteRange, err := ParseByteOffsetRange(fields[1])
	if err != nil {
		return TextOffsetRange{}, fmt.Errorf("byte range: %v", err)
	}

	return TextOffsetRange{
		TextOffset{
			Byte:       byteRange.From,
			LineColumn: textLineColumnRange.From,
		},
		TextOffset{
			Byte:       byteRange.Until,
			LineColumn: textLineColumnRange.Until,
		},
	}, nil
}

func (o TextOffsetRange) OffsetRangeFrom() Offset {
	return o.From
}

func (o TextOffsetRange) OffsetRangeUntil() Offset {
	return o.Until
}

func (cr TextOffsetRange) ByteOffsetRangeString() string {
	return cr.From.Byte.ByteOffsetString() + ":" + cr.Until.Byte.ByteOffsetString()
}

func (cr TextOffsetRange) TextOffsetRangeString() string {
	return cr.From.LineColumn.String() + ":" + cr.Until.LineColumn.String()
}

func (cr TextOffsetRange) OffsetRangeString() string {
	return cr.TextOffsetRangeString() + ";" + cr.ByteOffsetRangeString()
}

func (cr TextOffsetRange) String() string {
	return cr.OffsetRangeString()
}
