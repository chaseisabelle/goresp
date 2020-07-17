package goresp

import (
	"errors"
	"time"
)

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

func (b *Bool) Uint() (uint, error) {
	i, e := b.Int()

	return uint(i), e
}

func (b *Bool) Uint8() (uint8, error) {
	i, e := b.Int()

	return uint8(i), e
}

func (b *Bool) Uint16() (uint16, error) {
	i, e := b.Int()

	return uint16(i), e
}

func (b *Bool) Uint32() (uint32, error) {
	i, e := b.Int()

	return uint32(i), e
}

func (b *Bool) Uint64() (uint64, error) {
	i, e := b.Int()

	return uint64(i), e
}

func (b *Bool) Int() (int, error) {
	i := 0

	if b.internal {
		i = 1
	}

	return i, nil
}

func (b *Bool) Int32() (int32, error) {
	i, e := b.Int()

	return int32(i), e
}

func (b *Bool) Int64() (int64, error) {
	i, e := b.Int()

	return int64(i), e
}

func (b *Bool) Float32() (float32, error) {
	return 0, errors.New("cannot convert bool to float32")
}

func (b *Bool) Float64() (float64, error) {
	return 0, errors.New("cannot convert bool to float64")
}

func (b *Bool) Error() (error, error) {
	return nil, errors.New("cannot convert bool to error")
}

func (b *Bool) Array() ([]Value, error) {
	return nil, errors.New("cannot convert bool to array")
}

func (b *Bool) Time(string) (time.Time, error) {
	return time.Now(), errors.New("cannot convert bool to time")
}

func (b *Bool) Duration(d time.Duration) (time.Duration, error) {
	return d, errors.New("cannot convert bool to duration")
}

func (b *Bool) Null() error {
	return errors.New("bool is not null")
}
