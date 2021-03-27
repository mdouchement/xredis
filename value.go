package xredis

import (
	"errors"
	"fmt"
	"strconv"
)

type (
	// A Value is an agnostic data converter.
	// The conversion is based on your software knowledge.
	Value interface {
		// IsNil returns true if the value is nil.
		IsNil() bool
		// Float32 tries to convert into float32.
		Float32() (float32, error)
		// MustFloat32 converts into float32 or panic.
		MustFloat32() float32
		// Float64 tries to convert into float64.
		Float64() (float64, error)
		// MustFloat64 converts into float64 or panic.
		MustFloat64() float64
		// Int tries to convert into int.
		Int() (int, error)
		// MustInt converts into int or panic.
		MustInt() int
		// Int64 tries to convert into int64.
		Int64() (int64, error)
		// MustInt64 converts into int64 or panic.
		MustInt64() int64
		// Uint64 tries to convert into uint64.
		Uint64() (uint64, error)
		// MustUint64 converts into uint64 or panic.
		MustUint64() uint64
		// Bytes tries to convert into []bytes.
		Bytes() []byte
		// String tries to convert into string.
		String() string
		// Values tries to convert into []Value.
		Values() ([]Value, error)
		// MustValues converts into []Value or panic.
		MustValues() []Value
	}

	value struct {
		value interface{}
	}
)

var (
	// ErrNotASlice is used when a value is not castable into a slice.
	ErrNotASlice = errors.New("value is not a slice")
)

// IsNil returns true if the value is nil.
func (v *value) IsNil() bool {
	return v.value == nil
}

func (v *value) Float32() (float32, error) {
	x, err := v.Int64()
	return float32(x), err
}

func (v *value) MustFloat32() float32 {
	n, err := v.Float32()
	if err != nil {
		panic(err)
	}
	return n
}

func (v *value) Float64() (float64, error) {
	switch x := v.value.(type) {
	case int64:
		return float64(x), nil
	case string:
		return strconv.ParseFloat(x, 64)
	}
	return 0, fmt.Errorf("unsupported value type: %T", v.value)
}

func (v *value) MustFloat64() float64 {
	n, err := v.Float64()
	if err != nil {
		panic(err)
	}
	return n
}

func (v *value) Int() (int, error) {
	x, err := v.Int64()
	return int(x), err
}

func (v *value) MustInt() int {
	n, err := v.Int()
	if err != nil {
		panic(err)
	}
	return n
}

func (v *value) Int64() (int64, error) {
	switch x := v.value.(type) {
	case int64:
		return x, nil
	case string:
		return strconv.ParseInt(x, 10, 64)
	}
	return 0, fmt.Errorf("unsupported value type: %T", v.value)
}

func (v *value) MustInt64() int64 {
	n, err := v.Int64()
	if err != nil {
		panic(err)
	}
	return n
}

func (v *value) Uint64() (uint64, error) {
	switch x := v.value.(type) {
	case int64:
		return uint64(x), nil
	case string:
		return strconv.ParseUint(x, 10, 64)
	}
	return 0, fmt.Errorf("unsupported value type: %T", v.value)
}

func (v *value) MustUint64() uint64 {
	n, err := v.Uint64()
	if err != nil {
		panic(err)
	}
	return n
}

func (v *value) Bytes() []byte {
	return []byte(v.String())
}

func (v *value) String() string {
	switch x := v.value.(type) {
	case int64:
		return strconv.FormatInt(x, 10)
	case string:
		return x
	}
	panic(fmt.Errorf("unsupported value type: %T", v.value))
}

func (v *value) Values() ([]Value, error) {
	sl, ok := v.value.([]interface{})
	if !ok {
		return nil, ErrNotASlice
	}

	result := make([]Value, len(sl))
	for i, v := range sl {
		result[i] = &value{value: v}
	}
	return result, nil
}

func (v *value) MustValues() []Value {
	val, err := v.Values()
	if err != nil {
		panic(err)
	}
	return val
}

// MarshalJSON implements json.Marshaler.
func (v *value) MarshalJSON() ([]byte, error) {
	if v.IsNil() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", v.value)), nil
}
