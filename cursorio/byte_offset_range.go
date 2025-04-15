package cursorio

import (
	"errors"
	"strings"
)

type ByteOffsetRange struct {
	From  ByteOffset
	Until ByteOffset
}

var _ OffsetRange = ByteOffsetRange{}

func ParseByteOffsetRange(s string) (ByteOffsetRange, error) {
	split := strings.SplitN(s, ":", 2)
	if len(split) != 2 {
		return ByteOffsetRange{}, errors.New("invalid cursor range")
	}

	from, err := ParseByteOffset(split[0])
	if err != nil {
		return ByteOffsetRange{}, err
	}

	until, err := ParseByteOffset(split[1])
	if err != nil {
		return ByteOffsetRange{}, err
	}

	return ByteOffsetRange{from, until}, nil
}

func (c ByteOffsetRange) OffsetRangeFrom() Offset {
	return c.From
}

func (c ByteOffsetRange) OffsetRangeUntil() Offset {
	return c.Until
}

func (c ByteOffsetRange) ByteOffsetRangeString() string {
	return c.From.ByteOffsetString() + ":" + c.Until.ByteOffsetString()
}

func (c ByteOffsetRange) OffsetRangeString() string {
	return c.ByteOffsetRangeString()
}

func (c ByteOffsetRange) String() string {
	return c.ByteOffsetRangeString()
}
