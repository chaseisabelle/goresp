package goresp

import "errors"

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

func (e *Error) Integer() (int, error) {
	return 0, errors.New("cannot convert error to integer")
}

func (e *Error) Float() (float64, error) {
	return 0, errors.New("cannot convert error to float")
}

func (e *Error) Error() (error, error) {
	return e.internal, nil
}

func (e *Error) Array() ([]Value, error) {
	return nil, errors.New("cannot convert error to array")
}

func (e *Error) Null() error {
	return errors.New("error is not null")
}
