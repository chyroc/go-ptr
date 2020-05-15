package ptr

import "time"

func StringNoNonePtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func BoolNoNonePtr(s bool) *bool {
	if s == false {
		return nil
	}
	return &s
}

func ByteNoNonePtr(s byte) *byte {
	if s == 0 {
		return nil
	}
	return &s
}

func RuneNoNonePtr(s rune) *rune {
	if s == 0 {
		return nil
	}
	return &s
}

func IntNoNonePtr(s int) *int {
	if s == 0 {
		return nil
	}
	return &s
}

func Int8NoNonePtr(s int8) *int8 {
	if s == 0 {
		return nil
	}
	return &s
}

func Int16NoNonePtr(s int16) *int16 {
	if s == 0 {
		return nil
	}
	return &s
}

func Int32NoNonePtr(s int32) *int32 {
	if s == 0 {
		return nil
	}
	return &s
}

func Int64NoNonePtr(s int64) *int64 {
	if s == 0 {
		return nil
	}
	return &s
}

func UIntNoNonePtr(s uint) *uint {
	if s == 0 {
		return nil
	}
	return &s
}

func UInt8NoNonePtr(s uint8) *uint8 {
	if s == 0 {
		return nil
	}
	return &s
}

func UInt16NoNonePtr(s uint16) *uint16 {
	if s == 0 {
		return nil
	}
	return &s
}

func UInt32NoNonePtr(s uint32) *uint32 {
	if s == 0 {
		return nil
	}
	return &s
}

func UInt64NoNonePtr(s uint64) *uint64 {
	if s == 0 {
		return nil
	}
	return &s
}

func UIntptrNoNonePtr(s uintptr) *uintptr {
	if s == 0 {
		return nil
	}
	return &s
}

func Float32NoNonePtr(s float32) *float32 {
	if s == 0 {
		return nil
	}
	return &s
}

func Float64NoNonePtr(s float64) *float64 {
	if s == 0 {
		return nil
	}
	return &s
}

func Complex64NoNonePtr(s complex64) *complex64 {
	if s == 0 {
		return nil
	}
	return &s
}

func Complex128NoNonePtr(s complex128) *complex128 {
	if s == 0 {
		return nil
	}
	return &s
}

func ErrorNoNonePtr(s error) *error {
	if s == nil {
		return nil
	}
	return &s
}

func TimeNoNonePtr(s time.Time) *time.Time {
	if s.UnixNano() == 0 {
		return nil
	}
	return &s
}

func DurationNoNonePtr(s time.Duration) *time.Duration {
	if s == 0 {
		return nil
	}
	return &s
}
