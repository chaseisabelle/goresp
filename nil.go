package goresp

import "errors"

type Nil struct {
}

func NewNil() *Nil {
	return &Nil{}
}

func (n *Nil) Encode() ([]byte, error) {
	return []byte{'!', '\r', '\n'}, nil
}

func (n *Nil) Decode(bs []byte) error {
	_, v, e := decodeNil(0, bs)

	if e != nil {
		return e
	}

	if v == nil {
		return errors.New("failed to decode nil")
	}

	return nil
}

func (n *Nil) Bool() (bool, error) {
	return false, errors.New("cannot convert nil to bool")
}

func (n *Nil) Bytes() ([]byte, error) {
	return nil, errors.New("cannot convert nil to bytes")
}

func (n *Nil) String() (string, error) {
	return "", errors.New("cannot convert nil to string")
}

func (n *Nil) Integer() (int, error) {
	return 0, errors.New("cannot convert nil to integer")
}

func (n *Nil) Float() (float64, error) {
	return 0, errors.New("cannot convert nil to float")
}

func (n *Nil) Error() (error, error) {
	return nil, errors.New("cannot convert nil to error")
}

func (n *Nil) Array() ([]Value, error) {
	return nil, errors.New("cannot convert nil to array")
}

func (n *Nil) Null() error {
	return nil
}

