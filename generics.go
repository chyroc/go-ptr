//go:build go1.18
// +build go1.18

package ptr

func Ptr[T any](s T) *T {
	return &s
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~complex64 | ~complex128
}

// NumberNoNonePtr returns the pointer of int value, if the value is zero-value, it returns nil.
func NumberNoNonePtr[T Number](s T) *T {
	if s == 0 {
		return nil
	}
	return &s
}

// ValueNumber returns the value pointed to by the pointer
func ValueNumber[T Number](s *T) T {
	if s != nil {
		return *s
	}
	return 0
}

func ValueWithDefault[T any](s *T, defaultValue T) T {
	if s != nil {
		return *s
	}
	return defaultValue
}
