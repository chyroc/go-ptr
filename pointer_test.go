package ptr

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	as := assert.New(t)

	v0 := false
	v1 := true

	as.Equal(&v0, Bool(v0))
	as.Equal(&v1, Bool(v1))
	as.Equal((*bool)(nil), BoolNoNonePtr(v0))
	as.Equal(&v1, BoolNoNonePtr(v1))
	as.True(nil == BoolNoNonePtr(v0))
	as.Equal(v0, ValueBool(Bool(v0)))
	as.Equal(v0, ValueBool(nil))
	as.Equal(v1, ValueBool(Bool(v1)))
	as.Equal(v1, ValueBoolWithDefault(BoolNoNonePtr(v0), v1))
	as.Equal(v1, ValueBoolWithDefault(BoolNoNonePtr(v1), v1))
}

func TestByte(t *testing.T) {
	as := assert.New(t)

	v0 := byte(0)
	v1 := byte(1)

	as.Equal(&v0, Byte(v0))
	as.Equal(&v1, Byte(v1))
	as.Equal((*byte)(nil), ByteNoNonePtr(v0))
	as.Equal(&v1, ByteNoNonePtr(v1))
	as.True(nil == ByteNoNonePtr(v0))
	as.Equal(v0, ValueByte(Byte(v0)))
	as.Equal(v0, ValueByte(nil))
	as.Equal(v1, ValueByte(Byte(v1)))
	as.Equal(v1, ValueByteWithDefault(ByteNoNonePtr(v0), v1))
	as.Equal(v1, ValueByteWithDefault(ByteNoNonePtr(v1), v1))
}

func TestComplex128(t *testing.T) {
	as := assert.New(t)

	v0 := complex(0, 0)
	v1 := complex(1, 1)

	as.Equal(&v0, Complex128(v0))
	as.Equal(&v1, Complex128(v1))
	as.Equal((*complex128)(nil), Complex128NoNonePtr(v0))
	as.Equal(&v1, Complex128NoNonePtr(v1))
	as.True(nil == Complex128NoNonePtr(v0))
	as.Equal(v0, ValueComplex128(Complex128(v0)))
	as.Equal(v0, ValueComplex128(nil))
	as.Equal(v1, ValueComplex128(Complex128(v1)))
	as.Equal(v1, ValueComplex128WithDefault(Complex128NoNonePtr(v0), v1))
	as.Equal(v1, ValueComplex128WithDefault(Complex128NoNonePtr(v1), v1))
}

func TestComplex64(t *testing.T) {
	as := assert.New(t)

	v0 := complex(float32(0), float32(0))
	v1 := complex(float32(1), float32(1))

	as.Equal(&v0, Complex64(v0))
	as.Equal(&v1, Complex64(v1))
	as.Equal((*complex64)(nil), Complex64NoNonePtr(v0))
	as.Equal(&v1, Complex64NoNonePtr(v1))
	as.True(nil == Complex64NoNonePtr(v0))
	as.Equal(v0, ValueComplex64(Complex64(v0)))
	as.Equal(v0, ValueComplex64(nil))
	as.Equal(v1, ValueComplex64(Complex64(v1)))
	as.Equal(v1, ValueComplex64WithDefault(Complex64NoNonePtr(v0), v1))
	as.Equal(v1, ValueComplex64WithDefault(Complex64NoNonePtr(v1), v1))
}

func TestDuration(t *testing.T) {
	as := assert.New(t)

	v0 := time.Duration(0)
	v1 := time.Duration(1)

	as.Equal(&v0, Duration(v0))
	as.Equal(&v1, Duration(v1))
	as.Equal((*time.Duration)(nil), DurationNoNonePtr(v0))
	as.Equal(&v1, DurationNoNonePtr(v1))
	as.True(nil == DurationNoNonePtr(v0))
	as.Equal(v0, ValueDuration(Duration(v0)))
	as.Equal(v0, ValueDuration(nil))
	as.Equal(v1, ValueDuration(Duration(v1)))
	as.Equal(v1, ValueDurationWithDefault(DurationNoNonePtr(v0), v1))
	as.Equal(v1, ValueDurationWithDefault(DurationNoNonePtr(v1), v1))
}

func TestError(t *testing.T) {
	as := assert.New(t)

	v0 := (error)(nil)
	v1 := errors.New("1")

	as.Equal((*error)(nil), Error(v0))
	as.Equal(&v1, Error(v1))
	as.Equal((*error)(nil), ErrorNoNonePtr(v0))
	as.Equal(&v1, ErrorNoNonePtr(v1))
	as.True(nil == ErrorNoNonePtr(v0))
	as.Equal(v0, ValueError(Error(v0)))
	as.Equal(v0, ValueError(nil))
	as.Equal(v1, ValueError(Error(v1)))
	as.Equal(v1, ValueErrorWithDefault(ErrorNoNonePtr(v0), v1))
	as.Equal(v1, ValueErrorWithDefault(ErrorNoNonePtr(v1), v1))
}

func TestFloat32(t *testing.T) {
	as := assert.New(t)

	v0 := float32(0)
	v1 := float32(1)

	as.Equal(&v0, Float32(v0))
	as.Equal(&v1, Float32(v1))
	as.Equal((*float32)(nil), Float32NoNonePtr(v0))
	as.Equal(&v1, Float32NoNonePtr(v1))
	as.True(nil == Float32NoNonePtr(v0))
	as.Equal(v0, ValueFloat32(Float32(v0)))
	as.Equal(v0, ValueFloat32(nil))
	as.Equal(v1, ValueFloat32(Float32(v1)))
	as.Equal(v1, ValueFloat32WithDefault(Float32NoNonePtr(v0), v1))
	as.Equal(v1, ValueFloat32WithDefault(Float32NoNonePtr(v1), v1))
}

func TestFloat64(t *testing.T) {
	as := assert.New(t)

	v0 := float64(0)
	v1 := float64(1)

	as.Equal(&v0, Float64(v0))
	as.Equal(&v1, Float64(v1))
	as.Equal((*float64)(nil), Float64NoNonePtr(v0))
	as.Equal(&v1, Float64NoNonePtr(v1))
	as.True(nil == Float64NoNonePtr(v0))
	as.Equal(v0, ValueFloat64(Float64(v0)))
	as.Equal(v0, ValueFloat64(nil))
	as.Equal(v1, ValueFloat64(Float64(v1)))
	as.Equal(v1, ValueFloat64WithDefault(Float64NoNonePtr(v0), v1))
	as.Equal(v1, ValueFloat64WithDefault(Float64NoNonePtr(v1), v1))
}

func TestInt(t *testing.T) {
	as := assert.New(t)

	v0 := int(0)
	v1 := int(1)

	as.Equal(&v0, Int(v0))
	as.Equal(&v1, Int(v1))
	as.Equal((*int)(nil), IntNoNonePtr(v0))
	as.Equal(&v1, IntNoNonePtr(v1))
	as.True(nil == IntNoNonePtr(v0))
	as.Equal(v0, ValueInt(Int(v0)))
	as.Equal(v0, ValueInt(nil))
	as.Equal(v1, ValueInt(Int(v1)))
	as.Equal(v1, ValueIntWithDefault(IntNoNonePtr(v0), v1))
	as.Equal(v1, ValueIntWithDefault(IntNoNonePtr(v1), v1))
}

func TestInt8(t *testing.T) {
	as := assert.New(t)

	v0 := int8(0)
	v1 := int8(1)

	as.Equal(&v0, Int8(v0))
	as.Equal(&v1, Int8(v1))
	as.Equal((*int8)(nil), Int8NoNonePtr(v0))
	as.Equal(&v1, Int8NoNonePtr(v1))
	as.True(nil == Int8NoNonePtr(v0))
	as.Equal(v0, ValueInt8(Int8(v0)))
	as.Equal(v0, ValueInt8(nil))
	as.Equal(v1, ValueInt8(Int8(v1)))
	as.Equal(v1, ValueInt8WithDefault(Int8NoNonePtr(v0), v1))
	as.Equal(v1, ValueInt8WithDefault(Int8NoNonePtr(v1), v1))
}

func TestInt16(t *testing.T) {
	as := assert.New(t)

	v0 := int16(0)
	v1 := int16(1)

	as.Equal(&v0, Int16(v0))
	as.Equal(&v1, Int16(v1))
	as.Equal((*int16)(nil), Int16NoNonePtr(v0))
	as.Equal(&v1, Int16NoNonePtr(v1))
	as.True(nil == Int16NoNonePtr(v0))
	as.Equal(v0, ValueInt16(Int16(v0)))
	as.Equal(v0, ValueInt16(nil))
	as.Equal(v1, ValueInt16(Int16(v1)))
	as.Equal(v1, ValueInt16WithDefault(Int16NoNonePtr(v0), v1))
	as.Equal(v1, ValueInt16WithDefault(Int16NoNonePtr(v1), v1))
}

func TestInt32(t *testing.T) {
	as := assert.New(t)

	v0 := int32(0)
	v1 := int32(1)

	as.Equal(&v0, Int32(v0))
	as.Equal(&v1, Int32(v1))
	as.Equal((*int32)(nil), Int32NoNonePtr(v0))
	as.Equal(&v1, Int32NoNonePtr(v1))
	as.True(nil == Int32NoNonePtr(v0))
	as.Equal(v0, ValueInt32(Int32(v0)))
	as.Equal(v0, ValueInt32(nil))
	as.Equal(v1, ValueInt32(Int32(v1)))
	as.Equal(v1, ValueInt32WithDefault(Int32NoNonePtr(v0), v1))
	as.Equal(v1, ValueInt32WithDefault(Int32NoNonePtr(v1), v1))
}

func TestInt64(t *testing.T) {
	as := assert.New(t)

	v0 := int64(0)
	v1 := int64(1)

	as.Equal(&v0, Int64(v0))
	as.Equal(&v1, Int64(v1))
	as.Equal((*int64)(nil), Int64NoNonePtr(v0))
	as.Equal(&v1, Int64NoNonePtr(v1))
	as.True(nil == Int64NoNonePtr(v0))
	as.Equal(v0, ValueInt64(Int64(v0)))
	as.Equal(v0, ValueInt64(nil))
	as.Equal(v1, ValueInt64(Int64(v1)))
	as.Equal(v1, ValueInt64WithDefault(Int64NoNonePtr(v0), v1))
	as.Equal(v1, ValueInt64WithDefault(Int64NoNonePtr(v1), v1))
}

func TestRune(t *testing.T) {
	as := assert.New(t)

	v0 := rune(0)
	v1 := rune(1)

	as.Equal(&v0, Rune(v0))
	as.Equal(&v1, Rune(v1))
	as.Equal((*rune)(nil), RuneNoNonePtr(v0))
	as.Equal(&v1, RuneNoNonePtr(v1))
	as.True(nil == RuneNoNonePtr(v0))
	as.Equal(v0, ValueRune(Rune(v0)))
	as.Equal(v0, ValueRune(nil))
	as.Equal(v1, ValueRune(Rune(v1)))
	as.Equal(v1, ValueRuneWithDefault(RuneNoNonePtr(v0), v1))
	as.Equal(v1, ValueRuneWithDefault(RuneNoNonePtr(v1), v1))
}

func TestString(t *testing.T) {
	as := assert.New(t)

	v0 := string("")
	v1 := string("1")

	as.Equal(&v0, String(v0))
	as.Equal(&v1, String(v1))
	as.Equal((*string)(nil), StringNoNonePtr(v0))
	as.Equal(&v1, StringNoNonePtr(v1))
	as.True(nil == StringNoNonePtr(v0))
	as.Equal(v0, ValueString(String(v0)))
	as.Equal(v0, ValueString(nil))
	as.Equal(v1, ValueString(String(v1)))
	as.Equal(v1, ValueStringWithDefault(StringNoNonePtr(v0), v1))
	as.Equal(v1, ValueStringWithDefault(StringNoNonePtr(v1), v1))
}

func TestTime(t *testing.T) {
	as := assert.New(t)

	v0 := time.Unix(0, 0)
	v1 := time.Unix(1, 0)

	as.Equal(&v0, Time(v0))
	as.Equal(&v1, Time(v1))
	as.Equal((*time.Time)(nil), TimeNoNonePtr(v0))
	as.Equal(&v1, TimeNoNonePtr(v1))
	as.True(nil == TimeNoNonePtr(v0))
	as.Equal(v0, ValueTime(Time(v0)))
	as.Equal(v0, ValueTime(nil))
	as.Equal(v1, ValueTime(Time(v1)))
	as.Equal(v1, ValueTimeWithDefault(TimeNoNonePtr(v0), v1))
	as.Equal(v1, ValueTimeWithDefault(TimeNoNonePtr(v1), v1))
}

func TestUInt(t *testing.T) {
	as := assert.New(t)

	v0 := uint(0)
	v1 := uint(1)

	as.Equal(&v0, UInt(v0))
	as.Equal(&v1, UInt(v1))
	as.Equal((*uint)(nil), UIntNoNonePtr(v0))
	as.Equal(&v1, UIntNoNonePtr(v1))
	as.True(nil == UIntNoNonePtr(v0))
	as.Equal(v0, ValueUInt(UInt(v0)))
	as.Equal(v0, ValueUInt(nil))
	as.Equal(v1, ValueUInt(UInt(v1)))
	as.Equal(v1, ValueUIntWithDefault(UIntNoNonePtr(v0), v1))
	as.Equal(v1, ValueUIntWithDefault(UIntNoNonePtr(v1), v1))
}

func TestUInt8(t *testing.T) {
	as := assert.New(t)

	v0 := uint8(0)
	v1 := uint8(1)

	as.Equal(&v0, UInt8(v0))
	as.Equal(&v1, UInt8(v1))
	as.Equal((*uint8)(nil), UInt8NoNonePtr(v0))
	as.Equal(&v1, UInt8NoNonePtr(v1))
	as.True(nil == UInt8NoNonePtr(v0))
	as.Equal(v0, ValueUInt8(UInt8(v0)))
	as.Equal(v0, ValueUInt8(nil))
	as.Equal(v1, ValueUInt8(UInt8(v1)))
	as.Equal(v1, ValueUInt8WithDefault(UInt8NoNonePtr(v0), v1))
	as.Equal(v1, ValueUInt8WithDefault(UInt8NoNonePtr(v1), v1))
}

func TestUInt16(t *testing.T) {
	as := assert.New(t)

	v0 := uint16(0)
	v1 := uint16(1)

	as.Equal(&v0, UInt16(v0))
	as.Equal(&v1, UInt16(v1))
	as.Equal((*uint16)(nil), UInt16NoNonePtr(v0))
	as.Equal(&v1, UInt16NoNonePtr(v1))
	as.True(nil == UInt16NoNonePtr(v0))
	as.Equal(v0, ValueUInt16(UInt16(v0)))
	as.Equal(v0, ValueUInt16(nil))
	as.Equal(v1, ValueUInt16(UInt16(v1)))
	as.Equal(v1, ValueUInt16WithDefault(UInt16NoNonePtr(v0), v1))
	as.Equal(v1, ValueUInt16WithDefault(UInt16NoNonePtr(v1), v1))
}

func TestUInt32(t *testing.T) {
	as := assert.New(t)

	v0 := uint32(0)
	v1 := uint32(1)

	as.Equal(&v0, UInt32(v0))
	as.Equal(&v1, UInt32(v1))
	as.Equal((*uint32)(nil), UInt32NoNonePtr(v0))
	as.Equal(&v1, UInt32NoNonePtr(v1))
	as.True(nil == UInt32NoNonePtr(v0))
	as.Equal(v0, ValueUInt32(UInt32(v0)))
	as.Equal(v0, ValueUInt32(nil))
	as.Equal(v1, ValueUInt32(UInt32(v1)))
	as.Equal(v1, ValueUInt32WithDefault(UInt32NoNonePtr(v0), v1))
	as.Equal(v1, ValueUInt32WithDefault(UInt32NoNonePtr(v1), v1))
}

func TestUInt64(t *testing.T) {
	as := assert.New(t)

	v0 := uint64(0)
	v1 := uint64(1)

	as.Equal(&v0, UInt64(v0))
	as.Equal(&v1, UInt64(v1))
	as.Equal((*uint64)(nil), UInt64NoNonePtr(v0))
	as.Equal(&v1, UInt64NoNonePtr(v1))
	as.True(nil == UInt64NoNonePtr(v0))
	as.Equal(v0, ValueUInt64(UInt64(v0)))
	as.Equal(v0, ValueUInt64(nil))
	as.Equal(v1, ValueUInt64(UInt64(v1)))
	as.Equal(v1, ValueUInt64WithDefault(UInt64NoNonePtr(v0), v1))
	as.Equal(v1, ValueUInt64WithDefault(UInt64NoNonePtr(v1), v1))
}

func TestUIntptr(t *testing.T) {
	as := assert.New(t)

	v0 := uintptr(0)
	v1 := uintptr(1)

	as.Equal(&v0, UIntptr(v0))
	as.Equal(&v1, UIntptr(v1))
	as.Equal((*uintptr)(nil), UIntptrNoNonePtr(v0))
	as.Equal(&v1, UIntptrNoNonePtr(v1))
	as.True(nil == UIntptrNoNonePtr(v0))
	as.Equal(v0, ValueUIntptr(UIntptr(v0)))
	as.Equal(v0, ValueUIntptr(nil))
	as.Equal(v1, ValueUIntptr(UIntptr(v1)))
	as.Equal(v1, ValueUIntptrWithDefault(UIntptrNoNonePtr(v0), v1))
	as.Equal(v1, ValueUIntptrWithDefault(UIntptrNoNonePtr(v1), v1))
}
