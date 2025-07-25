package cursorio

import (
	"fmt"
	"io"
	"os"

	"github.com/apparentlymart/go-textseg/v16/textseg"
)

type TextWriter struct {
	offset TextOffset
	buf    []byte
}

var _ io.Writer = (*TextWriter)(nil)

func NewTextWriter(o TextOffset) *TextWriter {
	return &TextWriter{
		offset: o,
	}
}

func (w *TextWriter) Clone() *TextWriter {
	return &TextWriter{
		offset: w.offset,
		buf:    w.buf[:],
	}
}

func (w *TextWriter) GetByteOffset() ByteOffset {
	return ByteOffset(w.offset.Byte)
}

func (w *TextWriter) GetTextOffset() TextOffset {
	return w.offset
}

func (w *TextWriter) WriteRunes(p []rune, size int) (int, error) {
	w.write([]byte(string(p)), size, false)

	return size, nil
}

func (w *TextWriter) WriteRunesForOffset(p []rune, size int) TextOffset {
	w.write([]byte(string(p)), size, false)

	return w.offset
}

func (w *TextWriter) WriteRunesForOffsetRange(p []rune, size int) TextOffsetRange {
	from := w.offset

	w.write([]byte(string(p)), size, false)

	return TextOffsetRange{
		From:  from,
		Until: w.offset,
	}
}

func (w *TextWriter) WriteForOffset(p []byte) TextOffset {
	w.write(p, len(p), false)

	return w.offset
}

func (w *TextWriter) WriteForOffsetRange(p []byte) TextOffsetRange {
	from := w.offset

	w.write(p, len(p), false)

	return TextOffsetRange{
		From:  from,
		Until: w.offset,
	}
}

func (w *TextWriter) Write(p []byte) (int, error) {
	w.write(p, len(p), false)

	return len(p), nil
}

func (w *TextWriter) WriteEOF() {
	if len(w.buf) == 0 {
		return
	}

	w.write(nil, 0, true)
}

func (w *TextWriter) write(p []byte, psize int, atEOF bool) {
	if len(w.buf) > 0 {
		p = append(w.buf, p...)

		w.buf = nil
	}

	var graphemeByteCount int

	for len(p) > 0 {
		if p[0] == '\n' {
			w.offset.LineColumn[0]++
			w.offset.LineColumn[1] = 0

			p = p[1:]

			continue
		} else if p[0] == '\r' {
			if len(p) > 1 && p[1] == '\n' {
				w.offset.LineColumn[0]++
				w.offset.LineColumn[1] = 0

				p = p[2:]

				continue
			} else {
				// treat this as a hidden character
				// hacky? otherwise textseg waits indefinitely for the newline
				p = p[1:]

				continue
			}
		}

		graphemeByteCount, _, _ = textseg.ScanGraphemeClusters(p, atEOF)
		if graphemeByteCount == 0 {
			fmt.Fprintf(os.Stderr, "FATAL: no grapheme cluster found for bytes: %q\n", string(p))
			panic("no grapheme cluster found") // TODO possible?

			// tc.buf = buf

			// return
		}

		w.offset.LineColumn[1]++

		p = p[graphemeByteCount:]
	}

	w.offset.Byte += ByteOffset(psize)
}
