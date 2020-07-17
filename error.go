package goresp

import (
	"errors"
	"time"
)

type Error struct {
	internal error
}

func NewError(e error) *Error {
	return &Error{
		internal: e,
	}
}

func (e *Error) Encode() ([]byte, error) {
	b, err := e.Bytes()

	if err != nil {
		return nil, err
	}

	buffer := []byte{'-'}

	buffer = append(buffer, b...)
	buffer = append(buffer, '\r', '\n')

	return buffer, nil
}

func (e *Error) Decode(bs []byte) error {
	_, v, err := decodeError(0, bs)

	if err != nil {
		return err
	}

	e.internal = v.internal

	return nil
}

func (e *Error) Bool() (bool, error) {
	return false, errors.New("cannot convert error to bool")
}

func (e *Error) Bytes() ([]byte, error) {
	s, err := e.String()

	return[]byte(s), err
}

func (e *Error) String() (string, error) {
	tmp, err := e.Error()

	return tmp.Error(), err
}

func (e *Error) Uint() (uint, error) {
	return 0, errors.New("cannot convert error to uint")
}

func (e *Error) Uint8() (uint8, error) {
	return 0, errors.New("cannot convert error to uint8")
}

func (e *Error) Uint16() (uint16, error) {
	return 0, errors.New("cannot convert error to uint16")
}

func (e *Error) Uint32() (uint32, error) {
	return 0, errors.New("cannot convert error to uint32")
}

func (e *Error) Uint64() (uint64, error) {
	return 0, errors.New("cannot convert error to uint64")
}

func (e *Error) Int() (int, error) {
	return 0, errors.New("cannot convert error to int")
}

func (e *Error) Int32() (int32, error) {
	return 0, errors.New("cannot convert error to int32")
}

func (e *Error) Int64() (int64, error) {
	return 0, errors.New("cannot convert error to int64")
}

func (e *Error) Float32() (float32, error) {
	return 0, errors.New("cannot convert error to float32")
}

func (e *Error) Float64() (float64, error) {
	return 0, errors.New("cannot convert error to float64")
}

func (e *Error) Error() (error, error) {
	return e.internal, nil
}

func (e *Error) Array() ([]Value, error) {
	return nil, errors.New("cannot convert error to array")
}

func (e *Error) Time(string) (time.Time, error) {
	return time.Now(), errors.New("cannot convert error to time")
}

func (e *Error) Duration(d time.Duration) (time.Duration, error) {
	return d, errors.New("cannot convert error to duration")
}

func (e *Error) Null() error {
	return errors.New("error is not null")
}
