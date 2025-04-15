package cursorio

import (
	"errors"
	"fmt"
	"regexp"
)

type TextLineColumnRange struct {
	From  TextLineColumn
	Until TextLineColumn
}

func ParseTextLineColumnRange(v string) (TextLineColumnRange, error) {
	split := regexp.MustCompile(`:`).Split(v, 2)
	if len(split) != 2 {
		return TextLineColumnRange{}, errors.New("invalid text range")
	}

	from, err := ParseTextLineColumn(split[0])
	if err != nil {
		return TextLineColumnRange{}, fmt.Errorf("parse from: %w", err)
	}

	until, err := ParseTextLineColumn(split[1])
	if err != nil {
		return TextLineColumnRange{}, fmt.Errorf("parse until: %w", err)
	}

	return TextLineColumnRange{from, until}, nil
}

func (tlcr TextLineColumnRange) String() string {
	return fmt.Sprintf("%s:%s", tlcr.From.String(), tlcr.Until.String())
}
