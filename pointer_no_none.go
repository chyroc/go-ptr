package ptr

import "time"

// StringNoNonePtr returns the pointer of string value, if the value is zero-value, it returns nil
func StringNoNonePtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// BoolNoNonePtr returns the pointer of bool value, if the value is zero-value, it returns nil
func BoolNoNonePtr(s bool) *bool {
	if s == false {
		return nil
	}
	return &s
}

// ByteNoNonePtr returns the pointer of byte value, if the value is zero-value, it returns nil
func ByteNoNonePtr(s byte) *byte {
	if s == 0 {
		return nil
	}
	return &s
}

// RuneNoNonePtr returns the pointer of rune value, if the value is zero-value, it returns nil
func RuneNoNonePtr(s rune) *rune {
	if s == 0 {
		return nil
	}
	return &s
}

// IntNoNonePtr returns the pointer of int value, if the value is zero-value, it returns nil
func IntNoNonePtr(s int) *int {
	if s == 0 {
		return nil
	}
	return &s
}

// Int8NoNonePtr returns the pointer of int8 value, if the value is zero-value, it returns nil
func Int8NoNonePtr(s int8) *int8 {
	if s == 0 {
		return nil
	}
	return &s
}

// Int16NoNonePtr returns the pointer of int16 value, if the value is zero-value, it returns nil
func Int16NoNonePtr(s int16) *int16 {
	if s == 0 {
		return nil
	}
	return &s
}

// Int32NoNonePtr returns the pointer of int32 value, if the value is zero-value, it returns nil
func Int32NoNonePtr(s int32) *int32 {
	if s == 0 {
		return nil
	}
	return &s
}

// Int64NoNonePtr returns the pointer of int64 value, if the value is zero-value, it returns nil
func Int64NoNonePtr(s int64) *int64 {
	if s == 0 {
		return nil
	}
	return &s
}

// UIntNoNonePtr returns the pointer of uint value, if the value is zero-value, it returns nil
func UIntNoNonePtr(s uint) *uint {
	if s == 0 {
		return nil
	}
	return &s
}

// UInt8NoNonePtr returns the pointer of uint8 value, if the value is zero-value, it returns nil
func UInt8NoNonePtr(s uint8) *uint8 {
	if s == 0 {
		return nil
	}
	return &s
}

// UInt16NoNonePtr returns the pointer of uint16 value, if the value is zero-value, it returns nil
func UInt16NoNonePtr(s uint16) *uint16 {
	if s == 0 {
		return nil
	}
	return &s
}

// UInt32NoNonePtr returns the pointer of uint32 value, if the value is zero-value, it returns nil
func UInt32NoNonePtr(s uint32) *uint32 {
	if s == 0 {
		return nil
	}
	return &s
}

// UInt64NoNonePtr returns the pointer of uint64 value, if the value is zero-value, it returns nil
func UInt64NoNonePtr(s uint64) *uint64 {
	if s == 0 {
		return nil
	}
	return &s
}

// UIntptrNoNonePtr returns the pointer of uintptr value, if the value is zero-value, it returns nil
func UIntptrNoNonePtr(s uintptr) *uintptr {
	if s == 0 {
		return nil
	}
	return &s
}

// Float32NoNonePtr returns the pointer of float32 value, if the value is zero-value, it returns nil
func Float32NoNonePtr(s float32) *float32 {
	if s == 0 {
		return nil
	}
	return &s
}

// Float64NoNonePtr returns the pointer of float64 value, if the value is zero-value, it returns nil
func Float64NoNonePtr(s float64) *float64 {
	if s == 0 {
		return nil
	}
	return &s
}

// Complex64NoNonePtr returns the pointer of complex64 value, if the value is zero-value, it returns nil
func Complex64NoNonePtr(s complex64) *complex64 {
	if s == 0 {
		return nil
	}
	return &s
}

// Complex128NoNonePtr returns the pointer of complex128 value, if the value is zero-value, it returns nil
func Complex128NoNonePtr(s complex128) *complex128 {
	if s == 0 {
		return nil
	}
	return &s
}

// ErrorNoNonePtr returns the pointer of error value, if the value is zero-value, it returns nil
func ErrorNoNonePtr(s error) *error {
	if s == nil {
		return nil
	}
	return &s
}

// TimeNoNonePtr returns the pointer of time.Time value, if the value is zero-value, it returns nil
func TimeNoNonePtr(s time.Time) *time.Time {
	if s.UnixNano() == 0 {
		return nil
	}
	return &s
}

// DurationNoNonePtr returns the pointer of time.Duration value, if the value is zero-value, it returns nil
func DurationNoNonePtr(s time.Duration) *time.Duration {
	if s == 0 {
		return nil
	}
	return &s
}
