package goresp

import "time"

type Value interface {
	Encode() ([]byte, error)
	Decode([]byte) error
	Bool() (bool, error)
	Bytes() ([]byte, error)
	String() (string, error)
	Uint() (uint, error)
	Uint8() (uint8, error)
	Uint16() (uint16, error)
	Uint32() (uint32, error)
	Uint64() (uint64, error)
	Int() (int, error)
	Int32() (int32, error)
	Int64() (int64, error)
	Float32() (float32, error)
	Float64() (float64, error)
	Error() (error, error)
	Array() ([]Value, error)
	Time(string) (time.Time, error)
	Duration(time.Duration) (time.Duration, error)
	Null() error
}
