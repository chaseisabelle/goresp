package goresp

import (
	"errors"
	"time"
)

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

func (n *Null) Bool() (bool, error) {
	return false, errors.New("cannot convert null to bool")
}

func (n *Null) Bytes() ([]byte, error) {
	return nil, errors.New("cannot convert null to bytes")
}

func (n *Null) String() (string, error) {
	return "", errors.New("cannot convert null to string")
}

func (n *Null) Uint() (uint, error) {
	return 0, errors.New("cannot convert null to uint")
}

func (n *Null) Uint8() (uint8, error) {
	return 0, errors.New("cannot convert null to uint8")
}

func (n *Null) Uint16() (uint16, error) {
	return 0, errors.New("cannot convert null to uint16")
}

func (n *Null) Uint32() (uint32, error) {
	return 0, errors.New("cannot convert null to uint32")
}

func (n *Null) Uint64() (uint64, error) {
	return 0, errors.New("cannot convert null to uint64")
}

func (n *Null) Int() (int, error) {
	return 0, errors.New("cannot convert null to int")
}

func (n *Null) Int32() (int32, error) {
	return 0, errors.New("cannot convert null to int32")
}

func (n *Null) Int64() (int64, error) {
	return 0, errors.New("cannot convert null to int64")
}

func (n *Null) Float32() (float32, error) {
	return 0, errors.New("cannot convert null to float32")
}

func (n *Null) Float64() (float64, error) {
	return 0, errors.New("cannot convert null to float64")
}

func (n *Null) Error() (error, error) {
	return nil, errors.New("cannot convert null to error")
}

func (n *Null) Array() ([]Value, error) {
	return nil, errors.New("cannot convert null to array")
}

func (n *Null) Time(string) (time.Time, error) {
	return time.Now(), errors.New("cannot convert null to time")
}

func (n *Null) Duration(d time.Duration) (time.Duration, error) {
	return d, errors.New("cannot convert null to duration")
}

func (n *Null) Null() error {
	return nil
}
