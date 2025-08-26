package slogx

import (
	"fmt"
	"log/slog"
	"reflect"

	"github.com/go-softwarelab/common/pkg/types"
)

const (

	// ServiceKey is a predefined constant used as a key for identifying services or components in structured logging.
	ServiceKey = "service"
	// ComponentKey is a predefined constant used as a key for identifying components in structured logging.
	ComponentKey = "component"
	// ErrorKey is a predefined constant used as a key for identifying errors in structured logging.
	ErrorKey = "error"
	// UserIDKey is a predefined constant used as a key for identifying user IDs in structured logging.
	UserIDKey = "userId"
)

// Error returns a slog.Attr containing the provided error message under the "error" key.
func Error(err error) slog.Attr {
	return slog.String(ErrorKey, err.Error())
}

// Service creates a slog.Attr with the predefined ServiceKey and the given serviceName.
// This is a conventional attr for marking loggers for services/components in application.
func Service(serviceName string) slog.Attr {
	return slog.String(ServiceKey, serviceName)
}

// Component creates a slog.Attr with the predefined ComponentKey and the given componentName.
// This is a conventional attribute for marking loggers for components in an application.
// It is strongly recommended to use slogx.Service instead of this function.
// However, if you need to distinguish components (such as library tools) from services,
// this function can be useful.
func Component(componentName string) slog.Attr {
	return slog.String(ComponentKey, componentName)
}

// Number makes easier creation slog.Attr based on any number (int, float, uint) or the custom type over the number type.
func Number[T types.Number](key string, value T) slog.Attr {
	v := slog.AnyValue(value)
	if v.Kind() == slog.KindAny {
		v = valueForTypeOverPrimitive(value)
	}

	return slog.Attr{Key: key, Value: v}
}

// OptionalNumber makes easier creation slog.Attr based on pointer to any number (int, float, uint) or the custom type over the number type.
func OptionalNumber[T types.Number](key string, value *T) slog.Attr {
	if value == nil {
		return slog.String(key, "<nil>")
	}

	return Number(key, *value)
}

// String returns a slog.Attr with the provided key and string value.
// It's almost the same as a slog.String function, but allows for any custom type based on string to be passed as value.
func String[S ~string](key string, value S) slog.Attr {
	return slog.String(key, string(value))
}

// OptionalString returns a slog.Attr for a string pointer. If the pointer is nil, it sets the value to "<nil>".
func OptionalString[S ~string](key string, value *S) slog.Attr {
	if value == nil {
		return slog.String(key, "<nil>")
	}
	return String(key, *value)
}

// UserID creates a slog.Attr representing the "userId".
func UserID[ID types.Number | ~string](userID ID) slog.Attr {
	v := slog.AnyValue(userID)
	if v.Kind() == slog.KindAny {
		v = valueForTypeOverPrimitive(userID)
	}

	return slog.Attr{Key: UserIDKey, Value: v}
}

// OptionalUserID returns a slog.Attr representing the "userId" or a default value of "<unknown>" if userID is nil.
func OptionalUserID[ID types.Number | ~string](userID *ID) slog.Attr {
	if userID == nil {
		return slog.String(UserIDKey, "<unknown>")
	}
	return UserID(*userID)
}

// valueForTypeOverPrimitive - when the value is a primitive type wrapped as a custom type,
// the simple switch case won't handle it, so we need to reach for reflection :(
func valueForTypeOverPrimitive(value any) slog.Value {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return slog.Uint64Value(v.Uint())
	case reflect.Float32, reflect.Float64:
		return slog.Float64Value(v.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return slog.Int64Value(v.Int())
	case reflect.String:
		return slog.StringValue(v.String())
	case reflect.Bool:
		return slog.BoolValue(v.Bool())
	default:
		panic(fmt.Sprintf("unsupported type %T", value))
	}
}
