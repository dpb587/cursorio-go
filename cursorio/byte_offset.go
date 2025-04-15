package cursorio

import (
	"errors"
	"strconv"
)

// ByteOffset indicates the n-th byte of a stream. The first byte
type ByteOffset int64

var _ Offset = ByteOffset(0)

func ParseByteOffset(s string) (ByteOffset, error) {
	if len(s) < 3 || s[:2] != "0x" {
		return 0, errors.New("invalid cursor")
	}

	i, err := strconv.ParseInt(s[2:], 16, 64)

	return ByteOffset(i), err
}

func (c ByteOffset) ByteOffset() ByteOffset {
	return c
}

func (c ByteOffset) ByteOffsetString() string {
	return "0x" + strconv.FormatInt(int64(c), 16)
}

func (c ByteOffset) OffsetString() string {
	return c.ByteOffsetString()
}

func (c ByteOffset) String() string {
	return c.ByteOffsetString()
}
