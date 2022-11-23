//go:build go1.18
// +build go1.18

package ptr

func Ptr[T any](s T) *T {
	return &s
}

// NoNonePtr returns the pointer of int value, if the value is zero-value, it returns nil.
func NoNonePtr[T comparable](s T) *T {
	var empty T
	if s == empty {
		return nil
	}
	return &s
}

// Value returns the value pointed to by the pointer
func Value[T any](s *T) T {
	if s != nil {
		return *s
	}
	var empty T
	return empty
}

func ValueWithDefault[T any](s *T, defaultValue T) T {
	if s != nil {
		return *s
	}
	return defaultValue
}
