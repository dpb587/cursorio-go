package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dpb587/cursorio-go/cursorio"
)

func main() {
	all := bytes.NewBuffer(nil)
	cursor := cursorio.NewTextWriter(cursorio.TextOffset{})

	reader := bufio.NewReader(os.Stdin)

	var exiting bool
	var trailing []cursorio.TextOffsetRange

	for {
		r, rs, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			panic(err)
		} else if exiting {
			fmt.Fprintf(os.Stderr, "Exiting after first line...\n")

			break
		} else if r == '\n' {
			exiting = true

			continue
		}

		all.Write([]byte(string(r)))
		fmt.Fprintf(os.Stdout, "%s\n", all.String())

		pos := cursor.WriteRunesForOffsetRange([]rune{r}, rs)
		fmt.Fprintf(os.Stdout,
			"%s^ byte-offset %d; byte-count %d; text-range %s\n",
			strings.Repeat(" ", int(pos.From.LineColumn[1])),
			pos.From.Byte,
			pos.Until.Byte-pos.From.Byte,
			pos.TextOffsetRangeString(),
		)

		trailing = append(trailing, pos)

		if len(trailing) > 3 {
			fmt.Fprintf(os.Stdout,
				"%s==== range %s\n",
				strings.Repeat(" ", int(trailing[0].From.LineColumn[1])),
				cursorio.TextOffsetRange{
					From:  trailing[0].From,
					Until: pos.Until,
				},
			)

			trailing = trailing[1:]
		}

		fmt.Fprintf(os.Stdout, "\n")
	}
}
