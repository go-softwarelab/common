package to

// Options helper function to handle options pattern.
func Options[T any](opts ...func(*T)) T {
	var options T
	for _, opt := range opts {
		opt(&options)
	}
	return options
}

// OptionsWithDefault helper function to handle options pattern with default value for options.
func OptionsWithDefault[T any](defaultOptions T, opts ...func(*T)) T {
	options := defaultOptions
	for _, opt := range opts {
		opt(&options)
	}
	return options
}
