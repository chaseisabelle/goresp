package goresp

import (
	"errors"
	"fmt"
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
	n, e := f.Float()

	return fmt.Sprintf("%f", n), e
}

func (f *Float) Integer() (int, error) {
	return 0, errors.New("cannot convert float to integer")
}

func (f *Float) Float() (float64, error) {
	return f.internal, nil
}

func (f *Float) Error() (error, error) {
	str, err := f.String()

	return errors.New(str), err
}

func (f *Float) Array() ([]Value, error) {
	return nil, errors.New("cannot convert float to array")
}

func (f *Float) Null() error {
	return errors.New("float is not null")
}
