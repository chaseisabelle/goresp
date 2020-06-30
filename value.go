package goresp

type Value interface {
	Encode() ([]byte, error)
	Decode([]byte) error
	Bool() (bool, error)
	Bytes() ([]byte, error)
	String() (string, error)
	Integer() (int, error)
	Float() (float64, error)
	Error() (error, error)
	Array() ([]Value, error)
	Null() error
}
