package ptr

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	as := assert.New(t)

	vFalse := false
	vTrue := true

	as.Equal(&vTrue, Bool(true))
	as.Equal(&vFalse, Bool(false))
	as.Equal((*bool)(nil), BoolNoNonePtr(false))
	as.True(nil == BoolNoNonePtr(false))
	as.Equal(&vTrue, BoolNoNonePtr(true))
}

func TestByte(t *testing.T) {
	as := assert.New(t)

	v0 := byte(0)
	v1 := byte(1)

	as.Equal(&v1, Byte(v1))
	as.Equal((*byte)(nil), ByteNoNonePtr(v0))
	as.Equal(&v1, ByteNoNonePtr(v1))
}

func TestComplex128(t *testing.T) {
	as := assert.New(t)

	v0 := complex(0, 0)
	v1 := complex(1, 1)

	as.Equal(&v0, Complex128(v0))
	as.Equal(&v1, Complex128(v1))
	as.Equal((*complex128)(nil), Complex128NoNonePtr(v0))
	as.Equal(&v1, Complex128NoNonePtr(v1))
}

func TestComplex64(t *testing.T) {
	as := assert.New(t)

	v0 := complex(float32(0), float32(0))
	v1 := complex(float32(1), float32(1))

	as.Equal(&v0, Complex64(v0))
	as.Equal(&v1, Complex64(v1))
	as.Equal((*complex64)(nil), Complex64NoNonePtr(v0))
	as.Equal(&v1, Complex64NoNonePtr(v1))
}

func TestDuration(t *testing.T) {
	as := assert.New(t)

	v0 := time.Duration(0)
	v1 := time.Duration(1)

	as.Equal(&v0, Duration(v0))
	as.Equal(&v1, Duration(v1))
	as.Equal((*time.Duration)(nil), DurationNoNonePtr(v0))
	as.Equal(&v1, DurationNoNonePtr(v1))
}

func TestError(t *testing.T) {
	as := assert.New(t)

	v0 := (error)(nil)
	v1 := errors.New("1")

	as.Equal((*error)(nil), Error(v0))
	as.Equal(&v1, Error(v1))
	as.Equal((*error)(nil), ErrorNoNonePtr(v0))
	as.Equal(&v1, ErrorNoNonePtr(v1))
}

func TestFloat32(t *testing.T) {
	as := assert.New(t)

	v0 := float32(0)
	v1 := float32(1)

	as.Equal(&v0, Float32(v0))
	as.Equal(&v1, Float32(v1))
	as.Equal((*float32)(nil), Float32NoNonePtr(v0))
	as.Equal(&v1, Float32NoNonePtr(v1))
}

func TestFloat64(t *testing.T) {
	as := assert.New(t)

	v0 := float64(0)
	v1 := float64(1)

	as.Equal(&v0, Float64(v0))
	as.Equal(&v1, Float64(v1))
	as.Equal((*float64)(nil), Float64NoNonePtr(v0))
	as.Equal(&v1, Float64NoNonePtr(v1))
}

func TestInt(t *testing.T) {
	as := assert.New(t)

	v0 := int(0)
	v1 := int(1)

	as.Equal(&v0, Int(v0))
	as.Equal(&v1, Int(v1))
	as.Equal((*int)(nil), IntNoNonePtr(v0))
	as.Equal(&v1, IntNoNonePtr(v1))
}

func TestInt8(t *testing.T) {
	as := assert.New(t)

	v0 := int8(0)
	v1 := int8(1)

	as.Equal(&v0, Int8(v0))
	as.Equal(&v1, Int8(v1))
	as.Equal((*int8)(nil), Int8NoNonePtr(v0))
	as.Equal(&v1, Int8NoNonePtr(v1))
}

func TestInt16(t *testing.T) {
	as := assert.New(t)

	v0 := int16(0)
	v1 := int16(1)

	as.Equal(&v0, Int16(v0))
	as.Equal(&v1, Int16(v1))
	as.Equal((*int16)(nil), Int16NoNonePtr(v0))
	as.Equal(&v1, Int16NoNonePtr(v1))
}

func TestInt32(t *testing.T) {
	as := assert.New(t)

	v0 := int32(0)
	v1 := int32(1)

	as.Equal(&v0, Int32(v0))
	as.Equal(&v1, Int32(v1))
	as.Equal((*int32)(nil), Int32NoNonePtr(v0))
	as.Equal(&v1, Int32NoNonePtr(v1))
}

func TestInt64(t *testing.T) {
	as := assert.New(t)

	v0 := int64(0)
	v1 := int64(1)

	as.Equal(&v0, Int64(v0))
	as.Equal(&v1, Int64(v1))
	as.Equal((*int64)(nil), Int64NoNonePtr(v0))
	as.Equal(&v1, Int64NoNonePtr(v1))
}

func TestRune(t *testing.T) {
	as := assert.New(t)

	v0 := rune(0)
	v1 := rune(1)

	as.Equal(&v0, Rune(v0))
	as.Equal(&v1, Rune(v1))
	as.Equal((*rune)(nil), RuneNoNonePtr(v0))
	as.Equal(&v1, RuneNoNonePtr(v1))
}

func TestString(t *testing.T) {
	as := assert.New(t)

	v0 := string("")
	v1 := string("1")

	as.Equal(&v0, String(v0))
	as.Equal(&v1, String(v1))
	as.Equal((*string)(nil), StringNoNonePtr(v0))
	as.Equal(&v1, StringNoNonePtr(v1))
}

func TestTime(t *testing.T) {
	as := assert.New(t)

	v0 := time.Unix(0, 0)
	v1 := time.Unix(1, 0)

	as.Equal(&v0, Time(v0))
	as.Equal(&v1, Time(v1))
	as.Equal((*time.Time)(nil), TimeNoNonePtr(v0))
	as.Equal(&v1, TimeNoNonePtr(v1))
}

func TestUInt(t *testing.T) {
	as := assert.New(t)

	v0 := uint(0)
	v1 := uint(1)

	as.Equal(&v0, UInt(v0))
	as.Equal(&v1, UInt(v1))
	as.Equal((*uint)(nil), UIntNoNonePtr(v0))
	as.Equal(&v1, UIntNoNonePtr(v1))
}


func TestUInt8(t *testing.T) {
	as := assert.New(t)

	v0 := uint8(0)
	v1 := uint8(1)

	as.Equal(&v0, UInt8(v0))
	as.Equal(&v1, UInt8(v1))
	as.Equal((*uint8)(nil), UInt8NoNonePtr(v0))
	as.Equal(&v1, UInt8NoNonePtr(v1))
}


func TestUInt16(t *testing.T) {
	as := assert.New(t)

	v0 := uint16(0)
	v1 := uint16(1)

	as.Equal(&v0, UInt16(v0))
	as.Equal(&v1, UInt16(v1))
	as.Equal((*uint16)(nil), UInt16NoNonePtr(v0))
	as.Equal(&v1, UInt16NoNonePtr(v1))
}

func TestUInt32(t *testing.T) {
	as := assert.New(t)

	v0 := uint32(0)
	v1 := uint32(1)

	as.Equal(&v0, UInt32(v0))
	as.Equal(&v1, UInt32(v1))
	as.Equal((*uint32)(nil), UInt32NoNonePtr(v0))
	as.Equal(&v1, UInt32NoNonePtr(v1))
}

func TestUInt64(t *testing.T) {
	as := assert.New(t)

	v0 := uint64(0)
	v1 := uint64(1)

	as.Equal(&v0, UInt64(v0))
	as.Equal(&v1, UInt64(v1))
	as.Equal((*uint64)(nil), UInt64NoNonePtr(v0))
	as.Equal(&v1, UInt64NoNonePtr(v1))
}

func TestUIntptr(t *testing.T) {
	as := assert.New(t)

	v0 := uintptr(0)
	v1 := uintptr(1)

	as.Equal(&v0, UIntptr(v0))
	as.Equal(&v1, UIntptr(v1))
	as.Equal((*uintptr)(nil), UIntptrNoNonePtr(v0))
	as.Equal(&v1, UIntptrNoNonePtr(v1))
}
