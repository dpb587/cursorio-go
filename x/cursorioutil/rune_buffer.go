package cursorioutil

import (
	"bufio"
	"io"

	"github.com/dpb587/cursorio-go/cursorio"
)

type RuneBuffer struct {
	r io.RuneReader
	o int64

	buf  cursorio.DecodedRuneList
	bufi int
}

func NewRuneBuffer(r io.Reader) *RuneBuffer {
	rr, ok := r.(io.RuneReader)
	if !ok {
		rr = bufio.NewReader(r)
	}

	return &RuneBuffer{
		r: rr,
	}
}

func (d *RuneBuffer) GetByteOffset() cursorio.ByteOffset {
	return cursorio.ByteOffset(d.o)
}

func (d *RuneBuffer) NextRune() (cursorio.DecodedRune, error) {
	if d.bufi > 0 {
		r0 := d.buf[0]

		d.buf = d.buf[1:]
		d.bufi--

		d.o += int64(r0.Size)

		return r0, nil
	}

	r0, r0s, err := d.r.ReadRune()
	if err != nil {
		return cursorio.DecodedRune{}, err
	}

	d.o += int64(r0s)

	return cursorio.DecodedRune{
		Size: r0s,
		Rune: r0,
	}, nil
}

func (d *RuneBuffer) BacktrackRunes(runes ...cursorio.DecodedRune) {
	d.buf = append(runes, d.buf...)
	d.bufi += len(runes)

	for _, r := range runes {
		d.o -= int64(r.Size)
	}
}
