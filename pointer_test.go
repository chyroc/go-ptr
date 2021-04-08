package ptr

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	vFalse := false
	vTrue := true
	tests := []struct {
		name string
		args bool
		want *bool
	}{
		{
			name: "bool-false",
			args: false,
			want: &vFalse,
		},
		{
			name: "bool-true",
			args: true,
			want: &vTrue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte(t *testing.T) {
	v1 := byte(1)
	vS := byte('s')
	tests := []struct {
		name string
		args byte
		want *byte
	}{
		{
			name: "byte-1",
			args: byte(1),
			want: &v1,
		},
		{
			name: "byte-s",
			args: 's',
			want: &vS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Byte(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128(t *testing.T) {
	v1 := complex(1, 1)
	v2 := complex(200, 200)
	tests := []struct {
		name string
		args complex128
		want *complex128
	}{
		{
			name: "complex-1",
			args: complex(1, 1),
			want: &v1,
		},
		{
			name: "complex-2",
			args: complex(200, 200),
			want: &v2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Complex128(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64(t *testing.T) {
	v1 := complex(float32(1), float32(1))
	tests := []struct {
		name string
		args complex64
		want *complex64
	}{
		{
			name: "complex-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Complex64(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	v1 := time.Duration(1)
	tests := []struct {
		name string
		args time.Duration
		want *time.Duration
	}{
		{
			name: "duration-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Duration(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError(t *testing.T) {
	v1 := fmt.Errorf("1")
	tests := []struct {
		name string
		args error
		want *error
	}{
		{
			name: "err-nil",
			args: nil,
			want: nil,
		},
		{
			name: "err-nil",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Error(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	v1 := float32(1)
	tests := []struct {
		name string
		args float32
		want *float32
	}{
		{
			name: "float-32",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	v1 := float64(1)
	tests := []struct {
		name string
		args float64
		want *float64
	}{
		{
			name: "float-64",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	v1 := int(1)
	tests := []struct {
		name string
		args int
		want *int
	}{
		{
			name: "int-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	v1 := int16(1)
	tests := []struct {
		name string
		args int16
		want *int16
	}{
		{
			name: "int16-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	v1 := int32(1)
	tests := []struct {
		name string
		args int32
		want *int32
	}{
		{
			name: "int32-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	v1 := int64(1)
	tests := []struct {
		name string
		args int64
		want *int64
	}{
		{
			name: "int64-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	v1 := int8(1)
	tests := []struct {
		name string
		args int8
		want *int8
	}{
		{
			name: "int8-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRune(t *testing.T) {
	v1 := rune(1)
	tests := []struct {
		name string
		args rune
		want *rune
	}{
		{
			name: "rune-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rune(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	v1 := string("1")
	tests := []struct {
		name string
		args string
		want *string
	}{
		{
			name: "string-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime(t *testing.T) {
	v1 := time.Now()
	tests := []struct {
		name string
		args time.Time
		want *time.Time
	}{
		{
			name: "uint-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt(t *testing.T) {
	v1 := uint(1)
	tests := []struct {
		name string
		args uint
		want *uint
	}{
		{
			name: "uint-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt16(t *testing.T) {
	v1 := uint16(1)
	tests := []struct {
		name string
		args uint16
		want *uint16
	}{
		{
			name: "uint16-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt16(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt32(t *testing.T) {
	v1 := uint32(1)
	tests := []struct {
		name string
		args uint32
		want *uint32
	}{
		{
			name: "uint32-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt32(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt64(t *testing.T) {
	v1 := uint64(1)
	tests := []struct {
		name string
		args uint64
		want *uint64
	}{
		{
			name: "uint64-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt64(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt8(t *testing.T) {
	v1 := uint8(1)
	tests := []struct {
		name string
		args uint8
		want *uint8
	}{
		{
			name: "uint8-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt8(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUIntptr(t *testing.T) {
	v1 := uintptr(1)
	tests := []struct {
		name string
		args uintptr
		want *uintptr
	}{
		{
			name: "uintptr-1",
			args: v1,
			want: &v1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UIntptr(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UIntptr() = %v, want %v", got, tt.want)
			}
		})
	}
}
