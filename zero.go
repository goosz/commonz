package commonz

// Zero returns the zero value for any type T.
// This is useful for eliminating boilerplate like 'var zero T; return zero'.
func Zero[T any]() T {
	var zero T
	return zero
}
