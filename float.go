package goresp

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type Float struct {
	internal float64
}

func NewFloat(f float64) *Float {
	return &Float{
		internal: f,
	}
}

func (f *Float) Encode() ([]byte, error) {
	b, err := f.Bytes()

	if err != nil {
		return nil, err
	}

	buffer := []byte{'.'}

	buffer = append(buffer, b...)
	buffer = append(buffer, '\r', '\n')

	return buffer, nil
}

func (f *Float) Decode(bs []byte) error {
	_, v, e := decodeFloat(0, bs)

	if e != nil {
		return e
	}

	f.internal = v.internal

	return nil
}

func (f *Float) Bool() (bool, error) {
	return false, errors.New("cannot convert float to bool")
}

func (f *Float) Bytes() ([]byte, error) {
	s, e := f.String()

	return []byte(s), e
}

func (f *Float) String() (string, error) {
	n, e := f.Float64()

	return fmt.Sprintf("%f", n), e
}

func (f *Float) Uint() (uint, error) {
	i, e := f.Int()

	return uint(i), e
}

func (f *Float) Uint8() (uint8, error) {
	i, e := f.Int()

	return uint8(i), e
}

func (f *Float) Uint16() (uint16, error) {
	i, e := f.Int()

	return uint16(i), e
}

func (f *Float) Uint32() (uint32, error) {
	i, e := f.Int()

	return uint32(i), e
}

func (f *Float) Uint64() (uint64, error) {
	i, e := f.Int()

	return uint64(i), e
}

func (f *Float) Int() (int, error) {
	return 0, errors.New("cannot convert float to int")
}

func (f *Float) Int32() (int32, error) {
	return 0, errors.New("cannot convert float to int32")
}

func (f *Float) Int64() (int64, error) {
	return 0, errors.New("cannot convert float to int64")
}

func (f *Float) Float32() (float32, error) {
	f64, e := f.Float64()

	return float32(f64), e
}

func (f *Float) Float64() (float64, error) {
	return f.internal, nil
}

func (f *Float) Error() (error, error) {
	str, err := f.String()

	return errors.New(str), err
}

func (f *Float) Array() ([]Value, error) {
	return nil, errors.New("cannot convert float to array")
}

func (f *Float) Time(string) (time.Time, error) {
	f64, e := f.Float64()

	if e != nil {
		return time.Now(), e
	}

	s, d := math.Modf(f64)

	return time.Unix(int64(s), int64(d*1e9)), nil
}

func (f *Float) Duration(d time.Duration) (time.Duration, error) {
	f64, e := f.Float64()

	if e != nil {
		return d, e
	}

	return time.Duration(f64) * d, nil
}

func (f *Float) Null() error {
	return errors.New("float is not null")
}
