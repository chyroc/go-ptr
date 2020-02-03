package ptr

func ValueString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func ValueBool(s *bool) bool {
	if s != nil {
		return *s
	}
	return false
}

func ValueByte(s *byte) byte {
	if s != nil {
		return *s
	}
	return 0
}

func ValueRune(s *rune) rune {
	if s != nil {
		return *s
	}
	return 0
}

func ValueInt(s *int) int {
	if s != nil {
		return *s
	}
	return 0
}

func ValueInt8(s *int8) int8 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueInt16(s *int16) int16 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueInt32(s *int32) int32 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueInt64(s *int64) int64 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueUInt(s *uint) uint {
	if s != nil {
		return *s
	}
	return 0
}

func ValueUInt8(s *uint8) uint8 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueUInt16(s *uint16) uint16 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueUInt32(s *uint32) uint32 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueUInt64(s *uint64) uint64 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueUIntptr(s *uintptr) uintptr {
	if s != nil {
		return *s
	}
	return 0
}

func ValueFloat32(s *float32) float32 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueFloat64(s *float64) float64 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueComplex64(s *complex64) complex64 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueComplex128(s *complex128) complex128 {
	if s != nil {
		return *s
	}
	return 0
}

func ValueError(s *error) error {
	if s != nil {
		return *s
	}
	return nil
}
