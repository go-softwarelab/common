package to

import (
	"fmt"
	"strings"
)

// Enum converts the provided value to the available enum value in a case-insensitive manner.
// In case when the provided value doesn't match any of the enums, it will return an empty string as the enum and error.
// This is a case-insensitive version if you prefer to be more strict (value matches exactly one of the enums), use to.EnumStrict.
func Enum[V ~string, T ~string](value V, enumValues ...T) (T, error) {
	for _, e := range enumValues {
		if strings.EqualFold(string(value), string(e)) {
			return e, nil
		}
	}
	return "", fmt.Errorf("invalid value: %s doesn't match enum values %v", value, enumValues)
}

// EnumStrict converts the provided value to the available enum value in a strict manner (value matches exactly one of the enums).
// In case when the provided value doesn't match any of the enums, it will return an empty string as the enum and error.
// This is a strict version if you prefer to be more flexible about value (value matches one of the enums in case-insensitive manner), use to.Enum.
func EnumStrict[V ~string, T ~string](value V, enumValues ...T) (T, error) {
	for _, e := range enumValues {
		if string(value) == string(e) {
			return e, nil
		}
	}
	return "", fmt.Errorf("invalid value: %s doesn't match enum values %v", value, enumValues)
}
