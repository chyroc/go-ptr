//go:build go1.18
// +build go1.18

package ptr

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	as := assert.New(t)

	v0 := false
	v1 := true

	as.Equal(&v0, Ptr(v0))
	as.Equal(&v1, Bool(v1))
}

func Test_NoNonePtr(t *testing.T) {
	as := assert.New(t)

	{
		v0 := complex(0, 0)
		v1 := complex(1, 1)

		as.Equal((*complex128)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
	}

	{
		v0 := complex(float32(0), float32(0))
		v1 := complex(float32(1), float32(1))

		as.Equal((*complex64)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := float32(0)
		v1 := float32(1)

		as.Equal((*float32)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := float64(0)
		v1 := float64(1)

		as.Equal((*float64)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := int(0)
		v1 := int(1)

		as.Equal(&v0, Ptr(v0))
		as.Equal(&v1, Ptr(v1))
		as.Equal((*int)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(Ptr(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(Ptr(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := int8(0)
		v1 := int8(1)

		as.Equal(&v0, Int8(v0))
		as.Equal(&v1, Int8(v1))
		as.Equal((*int8)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(Int8(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(Int8(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := int16(0)
		v1 := int16(1)

		as.Equal(&v0, Int16(v0))
		as.Equal(&v1, Int16(v1))
		as.Equal((*int16)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(Int16(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(Int16(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := int32(0)
		v1 := int32(1)

		as.Equal(&v0, Int32(v0))
		as.Equal(&v1, Int32(v1))
		as.Equal((*int32)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(Int32(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(Int32(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := int64(0)
		v1 := int64(1)

		as.Equal(&v0, Int64(v0))
		as.Equal(&v1, Int64(v1))
		as.Equal((*int64)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(Int64(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(Int64(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := rune(0)
		v1 := rune(1)

		as.Equal(&v0, Rune(v0))
		as.Equal(&v1, Rune(v1))
		as.Equal((*rune)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, ValueRune(Rune(v0)))
		as.Equal(v0, ValueRune(nil))
		as.Equal(v1, ValueRune(Rune(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := uint(0)
		v1 := uint(1)

		as.Equal(&v0, UInt(v0))
		as.Equal(&v1, UInt(v1))
		as.Equal((*uint)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(UInt(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(UInt(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := uint8(0)
		v1 := uint8(1)

		as.Equal(&v0, UInt8(v0))
		as.Equal(&v1, UInt8(v1))
		as.Equal((*uint8)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(UInt8(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(UInt8(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := uint16(0)
		v1 := uint16(1)

		as.Equal(&v0, UInt16(v0))
		as.Equal(&v1, UInt16(v1))
		as.Equal((*uint16)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(UInt16(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(UInt16(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := uint32(0)
		v1 := uint32(1)

		as.Equal(&v0, UInt32(v0))
		as.Equal(&v1, UInt32(v1))
		as.Equal((*uint32)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(UInt32(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(UInt32(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := uint64(0)
		v1 := uint64(1)

		as.Equal(&v0, UInt64(v0))
		as.Equal(&v1, UInt64(v1))
		as.Equal((*uint64)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(UInt64(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(UInt64(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := uintptr(0)
		v1 := uintptr(1)

		as.Equal(&v0, UIntptr(v0))
		as.Equal(&v1, UIntptr(v1))
		as.Equal((*uintptr)(nil), NoNonePtr(v0))
		as.Equal(&v1, NoNonePtr(v1))
		as.True(nil == NoNonePtr(v0))
		as.Equal(v0, Value(UIntptr(v0)))
		// as.Equal(v0, Value(nil))
		as.Equal(v1, Value(UIntptr(v1)))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}

	{
		v0 := reflect.ValueOf(1)
		v1 := reflect.TypeOf(1)

		as.Equal(&v0, ReflectValue(v0))
		as.Equal(&v1, ReflectType(v1))
		// as.Equal((*uintptr)(nil), NoNonePtr(v0))
		// as.Equal(&v1, NoNonePtr(v1))
		// as.True(nil == NoNonePtr(v0))
		// as.Equal(v0, Value(UIntptr(v0)))
		// as.Equal(v0, Value(nil))
		// as.Equal(v1, Value(UIntptr(v1)))
		// as.Equal(v1, ValueWithDefault(NoNonePtr(v0), v1))
		// as.Equal(v1, ValueWithDefault(NoNonePtr(v1), v1))
	}
}
