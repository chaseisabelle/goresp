package goresp

import (
	"errors"
	"strconv"
	"time"
)

type SimpleString struct {
	internal string
}

func NewSimpleString(s string) *SimpleString {
	return &SimpleString{
		internal: s,
	}
}

func (ss *SimpleString) Encode() ([]byte, error) {
	b, err := ss.Bytes()

	if err != nil {
		return nil, err
	}

	return append(append([]byte{'+'}, b...), '\r', '\n'), nil
}

func (ss *SimpleString) Decode(bs []byte) error {
	_, v, e := decodeSimpleString(0, bs)

	if e != nil {
		return e
	}

	ss.internal = v.internal

	return nil
}

func (ss *SimpleString) Bool() (bool, error) {
	return false, errors.New("cannot convert simple string to bool")
}

func (ss *SimpleString) Bytes() ([]byte, error) {
	str, err := ss.String()

	return []byte(str), err
}

func (ss *SimpleString) String() (string, error) {
	return ss.internal, nil
}

func (ss *SimpleString) Uint() (uint, error) {
	i, e := ss.Int()

	return uint(i), e
}

func (ss *SimpleString) Uint8() (uint8, error) {
	i, e := ss.Int()

	return uint8(i), e
}

func (ss *SimpleString) Uint16() (uint16, error) {
	i, e := ss.Int()

	return uint16(i), e
}

func (ss *SimpleString) Uint32() (uint32, error) {
	i, e := ss.Int()

	return uint32(i), e
}

func (ss *SimpleString) Uint64() (uint64, error) {
	i, e := ss.Int()

	return uint64(i), e
}

func (ss *SimpleString) Int() (int, error) {
	str, err := ss.String()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(str)
}

func (ss *SimpleString) Int32() (int32, error) {
	x, e := ss.Int()

	return int32(x), e
}

func (ss *SimpleString) Int64() (int64, error) {
	x, e := ss.Int()

	return int64(x), e
}

func (ss *SimpleString) Float32() (float32, error) {
	f64, e := ss.Float64()

	return float32(f64), e
}

func (ss *SimpleString) Float64() (float64, error) {
	str, err := ss.String()

	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(str, 64)
}

func (ss *SimpleString) Error() (error, error) {
	str, err := ss.String()

	return errors.New(str), err
}

func (ss *SimpleString) Array() ([]Value, error) {
	return nil, errors.New("cannot convert simple string to array")
}

func (ss *SimpleString) Time(l string) (time.Time, error) {
	s, e := ss.String()

	if e != nil {
		return time.Now(), e
	}

	return time.Parse(l, s)
}

func (ss *SimpleString) Duration(d time.Duration) (time.Duration, error) {
	s, e := ss.String()

	if e != nil {
		return d, e
	}

	return time.ParseDuration(s)
}

func (ss *SimpleString) Null() error {
	return errors.New("simple string is not null")
}
