package goresp

import (
	"errors"
	"strconv"
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

func (i *Integer) Bytes() ([]byte, error) {
	s, e := i.String()

	return []byte(s), e
}

func (i *Integer) String() (string, error) {
	n, e := i.Integer()

	return strconv.Itoa(n), e
}

func (i *Integer) Integer() (int, error) {
	return i.internal, nil
}

func (i *Integer) Float() (float64, error) {
	n, e := i.Integer()

	return float64(n), e
}

func (i *Integer) Error() (error, error) {
	str, err := i.String()

	return errors.New(str), err
}

func (i *Integer) Array() ([]Value, error) {
	return nil, errors.New("cannot convert integer to array")
}

func (i *Integer) Null() error {
	return errors.New("integer is not null")
}
