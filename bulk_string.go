package goresp

import (
	"errors"
	"strconv"
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

func (bs *BulkString) String() (string, error) {
	return bs.internal, nil
}

func (bs *BulkString) Integer() (int, error) {
	str, err := bs.String()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(str)
}

func (bs *BulkString) Float() (float64, error) {
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

func (bs *BulkString) Null() error {
	return errors.New("bulk string is not null")
}
