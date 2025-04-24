package to

import (
	"fmt"
	"strconv"

	"github.com/go-softwarelab/common/pkg/types"
)

// BoolFromString will convert any string to bool
func BoolFromString(value string) (bool, error) {
	result, err := strconv.ParseBool(value)
	if err == nil {
		return result, nil
	}
	return false, fmt.Errorf("%w of %s to convert to bool: %w", ErrInvalidStringSyntax, value, err)
}

// BoolFromNumber will convert any number to bool
// 0 is false, any other number is true.
func BoolFromNumber[V types.Number](value V) bool {
	return value != 0
}
