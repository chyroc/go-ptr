package ptr

import (
	"reflect"
	"time"
)

// String returns the pointer of string value
func String(s string) *string {
	return &s
}

// Bool ean returns the pointer of boolean value
func Bool(s bool) *bool {
	return &s
}

// Byte returns the pointer of byte value
func Byte(s byte) *byte {
	return &s
}

// Rune returns the pointer of rune value
func Rune(s rune) *rune {
	return &s
}

// Int returns the pointer of int value
func Int(s int) *int {
	return &s
}

// Int8 returns the pointer of int8 value
func Int8(s int8) *int8 {
	return &s
}

// Int16 returns the pointer of int16 value
func Int16(s int16) *int16 {
	return &s
}

// Int32 returns the pointer of int32 value
func Int32(s int32) *int32 {
	return &s
}

// Int64 returns the pointer of int64 value
func Int64(s int64) *int64 {
	return &s
}

// UInt returns the pointer of uint value
func UInt(s uint) *uint {
	return &s
}

// UInt8 returns the pointer of uint8 value
func UInt8(s uint8) *uint8 {
	return &s
}

// UInt16 returns the pointer of uint16 value
func UInt16(s uint16) *uint16 {
	return &s
}

// UInt32 returns the pointer of uint32 value
func UInt32(s uint32) *uint32 {
	return &s
}

// UInt64 returns the pointer of uint64 value
func UInt64(s uint64) *uint64 {
	return &s
}

// UIntptr returns the pointer of uintptr value
func UIntptr(s uintptr) *uintptr {
	return &s
}

// Float32 returns the pointer of float32 value
func Float32(s float32) *float32 {
	return &s
}

// Float64 returns the pointer of float64 value
func Float64(s float64) *float64 {
	return &s
}

// Complex64 returns the pointer of complex64 value
func Complex64(s complex64) *complex64 {
	return &s
}

// Complex128 returns the pointer of complex128 value
func Complex128(s complex128) *complex128 {
	return &s
}

// Error returns the pointer of error value
func Error(s error) *error {
	if s == nil {
		return nil
	}
	return &s
}

// Time returns the pointer of time.Time value
func Time(s time.Time) *time.Time {
	return &s
}

// Duration returns the pointer of time.Duration value
func Duration(s time.Duration) *time.Duration {
	return &s
}

// ReflectValue returns the pointer of reflect.Value value
func ReflectValue(s reflect.Value) *reflect.Value {
	return &s
}

// ReflectType returns the pointer of reflect.Type value
func ReflectType(s reflect.Type) *reflect.Type {
	return &s
}
