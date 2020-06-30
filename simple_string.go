package goresp

import (
	"errors"
	"strconv"
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

func (ss *SimpleString) Integer() (int, error) {
	str, err := ss.String()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(str)
}

func (ss *SimpleString) Float() (float64, error) {
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

func (ss *SimpleString) Null() error {
	return errors.New("simple string is not null")
}
