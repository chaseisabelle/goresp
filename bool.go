package goresp

import "errors"

type Bool struct {
	internal bool
}

func NewBool(b bool) *Bool {
	return &Bool{
		internal: b,
	}
}

func (b *Bool) Encode() ([]byte, error) {
	raw := byte('<')

	if b.internal {
		raw = '>'
	}

	return []byte{raw, '\r', '\n'}, nil
}

func (b *Bool) Decode(bs []byte) error {
	_, v, e := decodeBool(0, bs)

	if e != nil {
		return e
	}

	b.internal = v.internal

	return nil
}

func (b *Bool) Bool() (bool, error) {
	return b.internal, nil
}

func (b *Bool) Bytes() ([]byte, error) {
	return nil, errors.New("cannot convert bool to bytes")
}

func (b *Bool) String() (string, error) {
	return "", errors.New("cannot convert bool to string")
}

func (b *Bool) Integer() (int, error) {
	i := 0

	if b.internal {
		i = 1
	}

	return i, nil
}

func (b *Bool) Float() (float64, error) {
	return 0, errors.New("cannot convert bool to float")
}

func (b *Bool) Error() (error, error) {
	return nil, errors.New("cannot convert bool to error")
}

func (b *Bool) Array() ([]Value, error) {
	return nil, errors.New("cannot convert bool to array")
}

func (b *Bool) Null() error {
	return errors.New("bool is not null")
}
