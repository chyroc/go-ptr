package ptr

func ValueStringWithDefault(s *string, defaultValue string) string {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueBoolWithDefault(s *bool, defaultValue bool) bool {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueByteWithDefault(s *byte, defaultValue byte) byte {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueRuneWithDefault(s *rune, defaultValue rune) rune {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueIntWithDefault(s *int, defaultValue int) int {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueInt8WithDefault(s *int8, defaultValue int8) int8 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueInt16WithDefault(s *int16, defaultValue int16) int16 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueInt32WithDefault(s *int32, defaultValue int32) int32 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueInt64WithDefault(s *int64, defaultValue int64) int64 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueUIntWithDefault(s *uint, defaultValue uint) uint {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueUInt8WithDefault(s *uint8, defaultValue uint8) uint8 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueUInt16WithDefault(s *uint16, defaultValue uint16) uint16 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueUInt32WithDefault(s *uint32, defaultValue uint32) uint32 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueUInt64WithDefault(s *uint64, defaultValue uint64) uint64 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueUIntptrWithDefault(s *uintptr, defaultValue uintptr) uintptr {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueFloat32WithDefault(s *float32, defaultValue float32) float32 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueFloat64WithDefault(s *float64, defaultValue float64) float64 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueComplex64WithDefault(s *complex64, defaultValue complex64) complex64 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueComplex128WithDefault(s *complex128, defaultValue complex128) complex128 {
	if s != nil {
		return *s
	}
	return defaultValue
}

func ValueErrorWithDefault(s *error, defaultValue error) error {
	if s != nil {
		return *s
	}
	return defaultValue
}
