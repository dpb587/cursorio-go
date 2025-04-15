package cursorio

import (
	"errors"
	"fmt"
)

type OffsetError struct {
	Offset Offset
	Err    error
}

func (e OffsetError) Error() string {
	return fmt.Sprintf("offset %s: %s", e.Offset.OffsetString(), e.Err.Error())
}

func (e OffsetError) Is(target error) bool {
	return errors.Is(e.Err, target)
}

func (e OffsetError) As(target interface{}) bool {
	if err, ok := target.(*OffsetError); ok {
		*err = e
		return true
	}

	return false
}

func (e OffsetError) Unwrap() error {
	return e.Err
}

//

type OffsetRangeError struct {
	OffsetRange OffsetRange
	Err         error
}

func (e OffsetRangeError) Error() string {
	return fmt.Sprintf("offset %s: %s", e.OffsetRange.OffsetRangeString(), e.Err.Error())
}

func (e OffsetRangeError) Is(target error) bool {
	return errors.Is(e.Err, target)
}

func (e OffsetRangeError) As(target interface{}) bool {
	if err, ok := target.(*OffsetRangeError); ok {
		*err = e
		return true
	}

	return false
}

func (e OffsetRangeError) Unwrap() error {
	return e.Err
}
