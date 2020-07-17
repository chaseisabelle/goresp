package goresp

import (
	"errors"
	"strconv"
	"time"
)

type Array struct {
	values []Value
}

func NewArray(vs []Value) *Array {
	return &Array{
		values: vs,
	}
}

func (a *Array) Encode() ([]byte, error) {
	values, err := a.Array()

	if err != nil {
		return nil, err
	}

	buffer := []byte{'*'}

	buffer = append(buffer, []byte(strconv.Itoa(len(values)))...)
	buffer = append(buffer, '\r', '\n')

	for _, value := range values {
		b, err := value.Encode()

		if err != nil {
			return nil, err
		}

		buffer = append(buffer, b...)
	}

	return buffer, nil
}

func (a *Array) Decode(bs []byte) error {
	_, v, e := decodeArray(0, bs)

	if e != nil {
		return e
	}

	vs, e := v.Array()

	if e != nil {
		return e
	}

	a.values = vs

	return nil
}

func (a *Array) Bool() (bool, error) {
	return false, errors.New("cannot convert array to bool")
}

func (a *Array) Bytes() ([]byte, error) {
	return nil, errors.New("cannot convert array to bytes")
}

func (a *Array) String() (string, error) {
	return "", errors.New("cannot convert array to string")
}

func (a *Array) Uint() (uint, error) {
	return 0, errors.New("cannot convert array to uint")
}

func (a *Array) Uint8() (uint8, error) {
	return 0, errors.New("cannot convert array to uint8")
}

func (a *Array) Uint16() (uint16, error) {
	return 0, errors.New("cannot convert array to uint16")
}

func (a *Array) Uint32() (uint32, error) {
	return 0, errors.New("cannot convert array to uint32")
}

func (a *Array) Uint64() (uint64, error) {
	return 0, errors.New("cannot convert array to uint64")
}

func (a *Array) Int() (int, error) {
	return 0, errors.New("cannot convert array to int")
}

func (a *Array) Int32() (int32, error) {
	return 0, errors.New("cannot convert array to int32")
}

func (a *Array) Int64() (int64, error) {
	return 0, errors.New("cannot convert array to int64")
}

func (a *Array) Float32() (float32, error) {
	return 0, errors.New("cannot convert array to float32")
}

func (a *Array) Float64() (float64, error) {
	return 0, errors.New("cannot convert array to float64")
}

func (a *Array) Error() (error, error) {
	return nil, errors.New("cannot convert array to error")
}

func (a *Array) Array() ([]Value, error) {
	return a.values, nil
}

func (a *Array) Time(string) (time.Time, error) {
	return time.Now(), errors.New("cannot convert array to time")
}

func (a *Array) Duration(d time.Duration) (time.Duration, error) {
	return d, errors.New("cannot convert array to duration")
}

func (a *Array) Null() error {
	return errors.New("array is not null")
}

func Strings(vs []Value) ([]string, error) {
	ss := make([]string, len(vs))

	for i, v := range vs {
		s, e := v.String()

		if e != nil {
			return nil, e
		}

		ss[i] = s
	}

	return ss, nil
}

func Uints(vs []Value) ([]uint, error) {
	is := make([]uint, len(vs))

	for i, v := range vs {
		j, e := v.Uint()

		if e != nil {
			return nil, e
		}

		is[i] = j
	}

	return is, nil
}

func Uint8s(vs []Value) ([]uint8, error) {
	is := make([]uint8, len(vs))

	for i, v := range vs {
		j, e := v.Uint8()

		if e != nil {
			return nil, e
		}

		is[i] = j
	}

	return is, nil
}

func Uint16s(vs []Value) ([]uint16, error) {
	is := make([]uint16, len(vs))

	for i, v := range vs {
		j, e := v.Uint16()

		if e != nil {
			return nil, e
		}

		is[i] = j
	}

	return is, nil
}

func Uint32s(vs []Value) ([]uint32, error) {
	is := make([]uint32, len(vs))

	for i, v := range vs {
		j, e := v.Uint32()

		if e != nil {
			return nil, e
		}

		is[i] = j
	}

	return is, nil
}

func Uint64s(vs []Value) ([]uint64, error) {
	is := make([]uint64, len(vs))

	for i, v := range vs {
		j, e := v.Uint64()

		if e != nil {
			return nil, e
		}

		is[i] = j
	}

	return is, nil
}

func Ints(vs []Value) ([]int, error) {
	is := make([]int, len(vs))

	for i, v := range vs {
		j, e := v.Int()

		if e != nil {
			return nil, e
		}

		is[i] = j
	}

	return is, nil
}

func Int32s(vs []Value) ([]int32, error) {
	is := make([]int32, len(vs))

	for i, v := range vs {
		j, e := v.Int32()

		if e != nil {
			return nil, e
		}

		is[i] = j
	}

	return is, nil
}

func Int64s(vs []Value) ([]int64, error) {
	is := make([]int64, len(vs))

	for i, v := range vs {
		j, e := v.Int64()

		if e != nil {
			return nil, e
		}

		is[i] = j
	}

	return is, nil
}

func Float32s(vs []Value) ([]float32, error) {
	fs := make([]float32, len(vs))

	for i, v := range vs {
		f, e := v.Float32()

		if e != nil {
			return nil, e
		}

		fs[i] = f
	}

	return fs, nil
}

func Float64s(vs []Value) ([]float64, error) {
	fs := make([]float64, len(vs))

	for i, v := range vs {
		f, e := v.Float64()

		if e != nil {
			return nil, e
		}

		fs[i] = f
	}

	return fs, nil
}

func Errors(vs []Value) ([]error, error) {
	es := make([]error, len(vs))

	for i, v := range vs {
		g, e := v.Error()

		if e != nil {
			return nil, e
		}

		es[i] = g
	}

	return es, nil
}

func Bools(vs []Value) ([]bool, error) {
	bs := make([]bool, len(vs))

	for i, v := range vs {
		b, e := v.Bool()

		if e != nil {
			return nil, e
		}

		bs[i] = b
	}

	return bs, nil
}

func Times(vs []Value, s string) ([]time.Time, error) {
	ts := make([]time.Time, len(vs))

	for _, v := range vs {
		t, e := v.Time(s)

		if e != nil {
			return nil, e
		}

		ts = append(ts, t)
	}

	return ts, nil
}

func Durations(vs []Value, d time.Duration) ([]time.Duration, error) {
	ds := make([]time.Duration, len(vs))

	for _, v := range vs {
		t, e := v.Duration(d)

		if e != nil {
			return nil, e
		}

		ds = append(ds, t)
	}

	return ds, nil
}