package cursorio

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var reTextLineColumnString = regexp.MustCompile(`^L(\d+)C(\d+)$`)

//

// TextLineColumn is the zero-based line and column position of a text. The very first symbol of a text document has a
// machine value of 0 for both line and column. When represented to humans, these values are 1-based.
type TextLineColumn [2]int64

// ParseTextLineColumn converts a human-friendly string based on the pattern of `L{line}C{column}` to a TextLineColumn.
func ParseTextLineColumn(v string) (TextLineColumn, error) {
	match := reTextLineColumnString.FindStringSubmatch(v)
	if match == nil {
		return TextLineColumn{}, errors.New("invalid text position")
	}

	matchLine, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		return TextLineColumn{}, fmt.Errorf("parse line: %w", err)
	}

	matchLineColumn, err := strconv.ParseInt(match[2], 10, 64)
	if err != nil {
		return TextLineColumn{}, fmt.Errorf("parse line column: %w", err)
	}

	return TextLineColumn{matchLine - 1, matchLineColumn - 1}, nil
}

func (tlc TextLineColumn) String() string {
	return fmt.Sprintf("L%dC%d", tlc[0]+1, tlc[1]+1)
}
