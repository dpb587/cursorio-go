package cursorio

type TextOffset struct {
	Byte       ByteOffset
	LineColumn TextLineColumn
}

var _ Offset = TextOffset{}

func (o TextOffset) ByteOffset() ByteOffset {
	return o.Byte
}

func (o TextOffset) OffsetString() string {
	return o.LineColumn.String() + ";" + o.Byte.OffsetString()
}

func (o TextOffset) String() string {
	return o.OffsetString()
}

func (o TextOffset) IsZero() bool {
	return o.Byte == 0 && o.LineColumn[0] == 0 && o.LineColumn[1] == 0
}
