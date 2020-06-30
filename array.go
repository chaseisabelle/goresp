package goresp

import (
	"errors"
	"strconv"
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

func (a *Array) Integer() (int, error) {
	return 0, errors.New("cannot convert array to integer")
}

func (a *Array) Float() (float64, error) {
	return 0, errors.New("cannot convert array to float")
}

func (a *Array) Error() (error, error) {
	return nil, errors.New("cannot convert array to error")
}

func (a *Array) Array() ([]Value, error) {
	return a.values, nil
}

func (a *Array) Null() error {
	return errors.New("array is not null")
}
