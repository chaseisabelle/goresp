package goresp

import "errors"

type Null struct {
}

func NewNull() *Null {
	return &Null{}
}

func (n *Null) Encode() ([]byte, error) {
	return []byte{'$', '-', '1', '\r', '\n'}, nil
}

func (n *Null) Decode(bs []byte) error {
	_, _, v, e := decodeBulkString(0, bs)

	if e != nil {
		return e
	}

	if v == nil {
		return errors.New("failed to decode null")
	}

	return nil
}

func (n *Null) Bytes() ([]byte, error) {
	return nil, errors.New("cannot convert null to bytes")
}

func (n *Null) String() (string, error) {
	return "", errors.New("cannot convert null to string")
}

func (n *Null) Integer() (int, error) {
	return 0, errors.New("cannot convert null to integer")
}

func (n *Null) Float() (float64, error) {
	return 0, errors.New("cannot convert null to float")
}

func (n *Null) Error() (error, error) {
	return nil, errors.New("cannot convert null to error")
}

func (n *Null) Array() ([]Value, error) {
	return nil, errors.New("cannot convert null to array")
}

func (n *Null) Null() error {
	return nil
}
