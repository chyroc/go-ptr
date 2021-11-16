package ptr

import (
	"time"
)

// ValueString returns the value pointed to by the pointer
func ValueString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// ValueTime returns the value pointed to by the pointer
func ValueTime(s *time.Time) time.Time {
	if s != nil {
		return *s
	}
	return time.Unix(0, 0)
}

// ValueBool returns the value pointed to by the pointer
func ValueBool(s *bool) bool {
	if s != nil {
		return *s
	}
	return false
}

// ValueByte returns the value pointed to by the pointer
func ValueByte(s *byte) byte {
	if s != nil {
		return *s
	}
	return 0
}

// ValueRune returns the value pointed to by the pointer
func ValueRune(s *rune) rune {
	if s != nil {
		return *s
	}
	return 0
}

// ValueInt returns the value pointed to by the pointer
func ValueInt(s *int) int {
	if s != nil {
		return *s
	}
	return 0
}

// ValueInt8 returns the value pointed to by the pointer
func ValueInt8(s *int8) int8 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueInt16 returns the value pointed to by the pointer
func ValueInt16(s *int16) int16 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueInt32 returns the value pointed to by the pointer
func ValueInt32(s *int32) int32 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueInt64 returns the value pointed to by the pointer
func ValueInt64(s *int64) int64 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueUInt returns the value pointed to by the pointer
func ValueUInt(s *uint) uint {
	if s != nil {
		return *s
	}
	return 0
}

// ValueUInt8 returns the value pointed to by the pointer
func ValueUInt8(s *uint8) uint8 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueUInt16 returns the value pointed to by the pointer
func ValueUInt16(s *uint16) uint16 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueUInt32 returns the value pointed to by the pointer
func ValueUInt32(s *uint32) uint32 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueUInt64 returns the value pointed to by the pointer
func ValueUInt64(s *uint64) uint64 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueUIntptr returns the value pointed to by the pointer
func ValueUIntptr(s *uintptr) uintptr {
	if s != nil {
		return *s
	}
	return 0
}

// ValueFloat32 returns the value pointed to by the pointer
func ValueFloat32(s *float32) float32 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueFloat64 returns the value pointed to by the pointer
func ValueFloat64(s *float64) float64 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueComplex64 returns the value pointed to by the pointer
func ValueComplex64(s *complex64) complex64 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueDuration returns the value pointed to by the pointer
func ValueDuration(s *time.Duration) time.Duration {
	if s != nil {
		return *s
	}
	return 0
}

// ValueComplex128 returns the value pointed to by the pointer
func ValueComplex128(s *complex128) complex128 {
	if s != nil {
		return *s
	}
	return 0
}

// ValueError returns the value pointed to by the pointer
func ValueError(s *error) error {
	if s != nil {
		return *s
	}
	return nil
}
