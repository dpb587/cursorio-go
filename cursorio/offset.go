package cursorio

type Offset interface {
	ByteOffset() ByteOffset

	OffsetString() string
}

type OffsetRange interface {
	OffsetRangeFrom() Offset
	OffsetRangeUntil() Offset

	OffsetRangeString() string
}
