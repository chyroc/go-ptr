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

func Test_NumberNoNonePtr(t *testing.T) {
	as := assert.New(t)

	{
		v0 := complex(0, 0)
		v1 := complex(1, 1)

		as.Equal((*complex128)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
	}

	{
		v0 := complex(float32(0), float32(0))
		v1 := complex(float32(1), float32(1))

		as.Equal((*complex64)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := float32(0)
		v1 := float32(1)

		as.Equal((*float32)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := float64(0)
		v1 := float64(1)

		as.Equal((*float64)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := int(0)
		v1 := int(1)

		as.Equal(&v0, Ptr(v0))
		as.Equal(&v1, Ptr(v1))
		as.Equal((*int)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(Ptr(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(Ptr(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := int8(0)
		v1 := int8(1)

		as.Equal(&v0, Int8(v0))
		as.Equal(&v1, Int8(v1))
		as.Equal((*int8)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(Int8(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(Int8(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := int16(0)
		v1 := int16(1)

		as.Equal(&v0, Int16(v0))
		as.Equal(&v1, Int16(v1))
		as.Equal((*int16)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(Int16(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(Int16(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := int32(0)
		v1 := int32(1)

		as.Equal(&v0, Int32(v0))
		as.Equal(&v1, Int32(v1))
		as.Equal((*int32)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(Int32(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(Int32(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := int64(0)
		v1 := int64(1)

		as.Equal(&v0, Int64(v0))
		as.Equal(&v1, Int64(v1))
		as.Equal((*int64)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(Int64(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(Int64(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := rune(0)
		v1 := rune(1)

		as.Equal(&v0, Rune(v0))
		as.Equal(&v1, Rune(v1))
		as.Equal((*rune)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueRune(Rune(v0)))
		as.Equal(v0, ValueRune(nil))
		as.Equal(v1, ValueRune(Rune(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := uint(0)
		v1 := uint(1)

		as.Equal(&v0, UInt(v0))
		as.Equal(&v1, UInt(v1))
		as.Equal((*uint)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(UInt(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(UInt(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := uint8(0)
		v1 := uint8(1)

		as.Equal(&v0, UInt8(v0))
		as.Equal(&v1, UInt8(v1))
		as.Equal((*uint8)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(UInt8(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(UInt8(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := uint16(0)
		v1 := uint16(1)

		as.Equal(&v0, UInt16(v0))
		as.Equal(&v1, UInt16(v1))
		as.Equal((*uint16)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(UInt16(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(UInt16(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := uint32(0)
		v1 := uint32(1)

		as.Equal(&v0, UInt32(v0))
		as.Equal(&v1, UInt32(v1))
		as.Equal((*uint32)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(UInt32(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(UInt32(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := uint64(0)
		v1 := uint64(1)

		as.Equal(&v0, UInt64(v0))
		as.Equal(&v1, UInt64(v1))
		as.Equal((*uint64)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(UInt64(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(UInt64(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := uintptr(0)
		v1 := uintptr(1)

		as.Equal(&v0, UIntptr(v0))
		as.Equal(&v1, UIntptr(v1))
		as.Equal((*uintptr)(nil), NumberNoNonePtr(v0))
		as.Equal(&v1, NumberNoNonePtr(v1))
		as.True(nil == NumberNoNonePtr(v0))
		as.Equal(v0, ValueNumber(UIntptr(v0)))
		// as.Equal(v0, ValueNumber(nil))
		as.Equal(v1, ValueNumber(UIntptr(v1)))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}

	{
		v0 := reflect.ValueOf(1)
		v1 := reflect.TypeOf(1)

		as.Equal(&v0, ReflectValue(v0))
		as.Equal(&v1, ReflectType(v1))
		// as.Equal((*uintptr)(nil), NumberNoNonePtr(v0))
		// as.Equal(&v1, NumberNoNonePtr(v1))
		// as.True(nil == NumberNoNonePtr(v0))
		// as.Equal(v0, ValueNumber(UIntptr(v0)))
		// as.Equal(v0, ValueNumber(nil))
		// as.Equal(v1, ValueNumber(UIntptr(v1)))
		// as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v0), v1))
		// as.Equal(v1, ValueWithDefault(NumberNoNonePtr(v1), v1))
	}
}
