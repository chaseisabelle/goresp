package goresp

import (
	"errors"
	"strconv"
	"time"
)

type Integer struct {
	internal int
}

func NewInteger(i int) *Integer {
	return &Integer{
		internal: i,
	}
}

func (i *Integer) Encode() ([]byte, error) {
	b, err := i.Bytes()

	if err != nil {
		return nil, err
	}

	buffer := []byte{':'}

	buffer = append(buffer, b...)
	buffer = append(buffer, '\r', '\n')

	return buffer, nil
}

func (i *Integer) Decode(bs []byte) error {
	_, v, e := decodeInteger(0, bs)

	if e != nil {
		return e
	}

	i.internal = v.internal

	return nil
}

func (i *Integer) Bool() (bool, error) {
	b := false

	if i.internal != 0 {
		b = true
	}

	return b, nil
}

func (i *Integer) Bytes() ([]byte, error) {
	s, e := i.String()

	return []byte(s), e
}

func (i *Integer) String() (string, error) {
	n, e := i.Int()

	return strconv.Itoa(n), e
}

func (i *Integer) Uint() (uint, error) {
	x, e := i.Int()

	return uint(x), e
}

func (i *Integer) Uint8() (uint8, error) {
	x, e := i.Int()

	return uint8(x), e
}

func (i *Integer) Uint16() (uint16, error) {
	x, e := i.Int()

	return uint16(x), e
}

func (i *Integer) Uint32() (uint32, error) {
	x, e := i.Int()

	return uint32(x), e
}

func (i *Integer) Uint64() (uint64, error) {
	x, e := i.Int()

	return uint64(x), e
}

func (i *Integer) Int() (int, error) {
	return i.internal, nil
}

func (i *Integer) Int32() (int32, error) {
	x, e := i.Int()

	return int32(x), e
}

func (i *Integer) Int64() (int64, error) {
	x, e := i.Int()

	return int64(x), e
}

func (i *Integer) Float32() (float32, error) {
	f, e := i.Float64()

	return float32(f), e
}

func (i *Integer) Float64() (float64, error) {
	n, e := i.Int()

	return float64(n), e
}

func (i *Integer) Error() (error, error) {
	str, err := i.String()

	return errors.New(str), err
}

func (i *Integer) Array() ([]Value, error) {
	return nil, errors.New("cannot convert integer to array")
}

func (i *Integer) Time(string) (time.Time, error) {
	i64, e := i.Int64()

	if e != nil {
		return time.Now(), e
	}

	return time.Unix(i64, 0), nil
}

func (i *Integer) Duration(d time.Duration) (time.Duration, error) {
	i64, e := i.Int64()

	if e != nil {
		return d, e
	}

	return time.Duration(i64) * d, nil
}

func (i *Integer) Null() error {
	return errors.New("integer is not null")
}
