package seqerr

// ConsumerWithoutError is a function that takes an element and returns nothing.
type ConsumerWithoutError[E any] = func(E)

// ConsumerWithError is a function that takes an element and returns an error.
type ConsumerWithError[E any] = func(E) error

// Consumer is a function that is consuming the sequence.
type Consumer[E any] interface {
	func(E) error | func(E)
}

func toConsumerWithError[E any, C Consumer[E]](consumer C) ConsumerWithError[E] {
	switch cons := any(consumer).(type) {
	case ConsumerWithoutError[E]:
		return func(e E) error {
			cons(e)
			return nil
		}
	case ConsumerWithError[E]:
		return cons
	default:
		panic("unknown type")
	}
}

// MapperWithError is a function that takes an element and returns a result and an error.
type MapperWithError[E any, R any] = func(E) (R, error)

// MapperWithoutError is a function that takes an element and returns a result.
type MapperWithoutError[E any, R any] = func(E) R

// Mapper is a function that takes an element and returns a result.
// It can return an error.
type Mapper[E any, R any] interface {
	func(E) R | func(E) (R, error)
}

// PredicateWithError is a function that takes an element and returns a boolean, it can fail with error.
type PredicateWithError[E any] = func(E) (bool, error)

// PredicateWithoutError is a function that takes an element and returns a boolean.
type PredicateWithoutError[E any] = func(E) bool

// Validator is a function that takes an element and returns an error.
type Validator[E any] = func(E) error

// Predicate is a function that takes an element and returns a boolean.
// It can return an error.
type Predicate[E any] interface {
	func(E) (bool, error) | func(E) bool | func(E) error
}

func toPredicateWithError[E any, P Predicate[E]](predicate P) PredicateWithError[E] {
	switch pred := any(predicate).(type) {
	case PredicateWithoutError[E]:
		return func(e E) (bool, error) {
			return pred(e), nil
		}
	case PredicateWithError[E]:
		return pred
	case Validator[E]:
		return func(e E) (bool, error) {
			err := pred(e)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	default:
		panic("unknown type")
	}
}
