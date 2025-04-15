package cursorioutil

import (
	"fmt"
	"strconv"
)

var errUnexpectedRune = "unexpected rune"

type UnexpectedRuneError struct {
	Rune rune
}

func (e UnexpectedRuneError) Error() string {
	return fmt.Sprintf("%s (%s)", errUnexpectedRune, strconv.QuoteRuneToASCII(e.Rune))
}
