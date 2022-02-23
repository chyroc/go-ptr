//go:build go1.18
// +build go1.18

package ptr

func Ptr[T any](s T) *T {
	return &s
}
