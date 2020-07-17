package goresp

import (
	"errors"
	"strconv"
	"time"
)

type BulkString struct {
	internal string
}

func NewBulkString(s string) *BulkString {
	return &BulkString{
		internal: s,
	}
}

func (bs *BulkString) Encode() ([]byte, error) {
	b, err := bs.Bytes()

	if err != nil {
		return nil, err
	}

	buffer := []byte{'$'}

	buffer = append(buffer, []byte(strconv.Itoa(len(b)))...)
	buffer = append(buffer, '\r', '\n')
	buffer = append(buffer, b...)
	buffer = append(buffer, '\r', '\n')

	return buffer, nil
}

func (bs *BulkString) Decode(b []byte) error {
	_, v, _, e := decodeBulkString(0, b)

	if e != nil {
		return e
	}

	if v == nil {
		return errors.New("failed to decode bulk string")
	}

	bs.internal = v.internal

	return nil
}

func (bs *BulkString) Bytes() ([]byte, error) {
	s, e := bs.String()

	return []byte(s), e
}

func (bs *BulkString) Bool() (bool, error) {
	return false, errors.New("cannot convert bulk string to bool")
}

func (bs *BulkString) String() (string, error) {
	return bs.internal, nil
}

func (bs *BulkString) Uint() (uint, error) {
	i, e := bs.Int()

	return uint(i), e
}

func (bs *BulkString) Uint8() (uint8, error) {
	i, e := bs.Int()

	return uint8(i), e
}

func (bs *BulkString) Uint16() (uint16, error) {
	i, e := bs.Int()

	return uint16(i), e
}

func (bs *BulkString) Uint32() (uint32, error) {
	i, e := bs.Int()

	return uint32(i), e
}

func (bs *BulkString) Uint64() (uint64, error) {
	i, e := bs.Int()

	return uint64(i), e
}

func (bs *BulkString) Int() (int, error) {
	str, err := bs.String()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(str)
}

func (bs *BulkString) Int32() (int32, error) {
	x, e := bs.Int()

	return int32(x), e
}

func (bs *BulkString) Int64() (int64, error) {
	x, e := bs.Int()

	return int64(x), e
}

func (bs *BulkString) Float32() (float32, error) {
	f64, e := bs.Float64()

	return float32(f64), e
}

func (bs *BulkString) Float64() (float64, error) {
	str, err := bs.String()

	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(str, 64)
}

func (bs *BulkString) Error() (error, error) {
	str, err := bs.String()

	return errors.New(str), err
}

func (bs *BulkString) Array() ([]Value, error) {
	return nil, errors.New("cannot convert bulk string to array")
}

func (bs *BulkString) Time(l string) (time.Time, error) {
	s, e := bs.String()

	if e != nil {
		return time.Now(), e
	}

	return time.Parse(l, s)
}

func (bs *BulkString) Duration(d time.Duration) (time.Duration, error) {
	s, e := bs.String()

	if e != nil {
		return d, e
	}

	return time.ParseDuration(s)
}

func (bs *BulkString) Null() error {
	return errors.New("bulk string is not null")
}
