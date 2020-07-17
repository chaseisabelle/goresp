package goresp

import (
	"errors"
	"time"
)

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

func (n *Nil) Uint() (uint, error) {
	return 0, errors.New("cannot convert nil to uint")
}

func (n *Nil) Uint8() (uint8, error) {
	return 0, errors.New("cannot convert nil to uint8")
}

func (n *Nil) Uint16() (uint16, error) {
	return 0, errors.New("cannot convert nil to uint16")
}

func (n *Nil) Uint32() (uint32, error) {
	return 0, errors.New("cannot convert nil to uint32")
}

func (n *Nil) Uint64() (uint64, error) {
	return 0, errors.New("cannot convert nil to uint64")
}

func (n *Nil) Int() (int, error) {
	return 0, errors.New("cannot convert nil to int")
}

func (n *Nil) Int32() (int32, error) {
	return 0, errors.New("cannot convert nil to int32")
}

func (n *Nil) Int64() (int64, error) {
	return 0, errors.New("cannot convert nil to int64")
}

func (n *Nil) Float32() (float32, error) {
	return 0, errors.New("cannot convert nil to float32")
}

func (n *Nil) Float64() (float64, error) {
	return 0, errors.New("cannot convert nil to float64")
}

func (n *Nil) Error() (error, error) {
	return nil, errors.New("cannot convert nil to error")
}

func (n *Nil) Array() ([]Value, error) {
	return nil, errors.New("cannot convert nil to array")
}

func (n *Nil) Time(string) (time.Time, error) {
	return time.Now(), errors.New("cannot convert nil to time")
}

func (n *Nil) Duration(d time.Duration) (time.Duration, error) {
	return d, errors.New("cannot convert nil to duration")
}

func (n *Nil) Null() error {
	return nil
}

