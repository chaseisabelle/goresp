package goresp

import (
	"errors"
	"fmt"
	"strconv"
)

func Decode(bs []byte) ([]Value, error) {
	_, vs, e := decode(0, bs, -1)

	return vs, e
}

func decode(i int, bs []byte, m int) (int, []Value, error) {
	if m < -1 {
		return i, nil, fmt.Errorf("invalid max values: %d", m)
	}

	vs := make([]Value, 0)
	l := len(bs)
	s := 0

	for i := i; i < l; i++ {
		t := len(vs)

		if m != -1 {
			if t > m {
				return i - 1, nil, fmt.Errorf("exceeded max values: %d", m)
			}

			if t == m {
				return i - 1, vs, nil
			}
		}

		b := bs[i]

		var v Value
		var e error
		var n *Null

		n = nil

		switch s {
		case 0:
			switch b {
			case '*':
				i, v, e = decodeArray(i, bs)
			case '+':
				i, v, e = decodeSimpleString(i, bs)
			case '$':
				i, v, n, e = decodeBulkString(i, bs)
			case '-':
				i, v, e = decodeError(i, bs)
			case ':':
				i, v, e = decodeInteger(i, bs)
			case '!':
				i, v, e = decodeNil(i, bs)
			case '.':
				i, v, e = decodeFloat(i, bs)
			default:
				return i, nil, badByte("failed to parse input", b, i)
			}

			if e != nil {
				return i, vs, e
			}

			if n != nil {
				v = n
			}

			vs = append(vs, v)
		}
	}

	return l, vs, nil
}

func decodeInteger(i int, bs []byte) (int, *Integer, error) {
	s := 0
	t := make([]byte, 0)

	for i := i; i < len(bs); i++ {
		b := bs[i]

		switch s {
		case 0:
			if b != ':' {
				return i, nil, badByte("failed to parse integer identifier", b, i)
			}

			s = 1
		case 1:
			if b == '\r' {
				s = 2
			} else {
				t = append(t, b)
			}
		case 2:
			if b == '\n' {
				n, e := strconv.Atoi(string(t))

				if e != nil {
					return i, nil, fmt.Errorf("failed to parse integer: %s", e.Error())
				}

				return i, NewInteger(n), nil
			}

			s = 1

			t = append(t, '\r', b)
		default:
			return i, nil, fmt.Errorf("failed to parse integer: invalid state %d", s)
		}
	}

	return len(bs), nil, badEOI("failed to parse integer")
}

func decodeFloat(i int, bs []byte) (int, *Float, error) {
	s := 0
	t := make([]byte, 0)

	for i := i; i < len(bs); i++ {
		b := bs[i]

		switch s {
		case 0:
			if b != '.' {
				return i, nil, badByte("failed to parse float identifier", b, i)
			}

			s = 1
		case 1:
			if b == '\r' {
				s = 2
			} else {
				t = append(t, b)
			}
		case 2:
			if b == '\n' {
				n, e := strconv.ParseFloat(string(t), 64)

				if e != nil {
					return i, nil, fmt.Errorf("failed to parse float: %s", e.Error())
				}

				return i, NewFloat(n), nil
			}

			s = 1

			t = append(t, '\r', b)
		default:
			return i, nil, fmt.Errorf("failed to parse float: invalid state %d", s)
		}
	}

	return len(bs), nil, badEOI("failed to parse float")
}

func decodeError(i int, bs []byte) (int, *Error, error) {
	s := 0
	t := make([]byte, 0)

	for i := i; i < len(bs); i++ {
		b := bs[i]

		switch s {
		case 0:
			if b != '-' {
				return i, nil, badByte("failed to parse error identifier", b, i)
			}

			s = 1
		case 1:
			if b == '\r' {
				s = 2
			} else {
				t = append(t, b)
			}
		case 2:
			if b == '\n' {
				return i, NewError(errors.New(string(t))), nil
			}

			s = 1

			t = append(t, '\r', b)
		default:
			return i, nil, fmt.Errorf("failed to parse error: invalid state %d", s)
		}
	}

	return len(bs), nil, badEOI("failed to parse error")
}

func decodeBulkString(i int, bs []byte) (int, *BulkString, *Null, error) {
	s := 0
	l := -2
	t := make([]byte, 0)

	var e error

	for i := i; i < len(bs); i++ {
		b := bs[i]

		switch s {
		case 0:
			if b != '$' {
				return i, nil, nil, badByte("failed to parse bulk string identifier", b, i)
			}

			s = 1
		case 1:
			if b == '\r' {
				s = 2
			} else {
				t = append(t, b)
			}
		case 2:
			if b != '\n' {
				return i, nil, nil, badByte("failed to parse bulk string length delimiter", b, i)
			}

			l, e = strconv.Atoi(string(t))

			if e != nil {
				return i, nil, nil, badLen("failed to parse bulk string length", e)
			}

			if l == -1 {
				return i, nil, NewNull(), nil
			}

			t = make([]byte, 0)
			s = 3
		case 3:
			if b == '\r' {
				s = 4
			} else {
				t = append(t, b)
			}
		case 4:
			if b == '\n' {
				c := len(t)

				if l != c {
					return i, nil, nil, fmt.Errorf("bulk string length mismatch: expected %d, parsed %d", l, c)
				}

				return i, NewBulkString(string(t)), nil, nil
			}

			t = append(t, '\r', b)
			s = 3
		default:
			return i, nil, nil, fmt.Errorf("failed to parse bulk string: invalid state %d", s)
		}
	}

	return len(bs), nil, nil, badEOI("failed to parse bulk string")
}

func decodeSimpleString(i int, bs []byte) (int, *SimpleString, error) {
	s := 0
	t := make([]byte, 0)

	for i := i; i < len(bs); i++ {
		b := bs[i]

		switch s {
		case 0:
			if b != '+' {
				return i, nil, badByte("failed to parse simple string identifier", b, i)
			}

			s = 1
		case 1:
			if b == '\r' {
				s = 2
			} else {
				t = append(t, b)
			}
		case 2:
			if b == '\n' {
				return i, NewSimpleString(string(t)), nil
			}

			s = 1

			t = append(t, '\r', b)
		default:
			return i, nil, fmt.Errorf("failed to parse simple string: invalid state %d", s)
		}
	}

	return len(bs), nil, badEOI("failed to parse simple string")
}

func decodeArray(i int, bs []byte) (int, *Array, error) {
	s := 0
	c := make([]byte, 0)
	l := -1

	var e error

	for i := i; i < len(bs); i++ {
		b := bs[i]

		switch s {
		case 0:
			if b != '*' {
				return i, nil, badByte("failed to parse array identifier", b, i)
			}

			s = 1
		case 1:
			if b >= '0' && b <= '9' {
				c = append(c, b)

				continue
			}

			if b != '\r' {
				return i, nil, badByte("failed to parse delimiter", b, i)
			}

			l, e = strconv.Atoi(string(c))

			if e != nil {
				return i, nil, badLen("failed to parse array count", e)
			}

			if l < 0 {
				return i, nil, fmt.Errorf("invalid array count: %d", l)
			}

			s = 2
		case 2:
			if b != '\n' {
				return i, nil, badByte("failed to parse delimiter", b, i)
			}

			s = 3
		case 3:
			i, vs, e := decode(i, bs, l)

			if e != nil {
				return i, nil, e
			}

			if l != len(vs) {
				e = fmt.Errorf("array count mismatch: expected %d, parsed %d", l, len(vs))
			}

			return i, NewArray(vs), e
		default:
			return i, nil, fmt.Errorf("failed to parse array: invalid state %d", s)
		}
	}

	if s == 3 && l == 0 {
		return len(bs), NewArray(make([]Value, 0)), nil
	}

	return len(bs), nil, badEOI("failed to parse array")
}

func decodeNil(i int, bs []byte) (int, *Nil, error) {
	s := 0

	for i := i; i < len(bs); i++ {
		b := bs[i]

		switch s {
		case 0:
			if b != '!' {
				return i, nil, badByte("failed to parse nil identifier", b, i)
			}

			s = 1
		case 1:
			if b != '\r' {
				return i, nil, badByte("failed to parse nil delimiter", b, i)
			}

			s = 2
		case 2:
			if b != '\n' {
				return i, nil, badByte("failed to parse nil delimiter", b, i)
			}

			return i, NewNil(), nil
		default:
			return i, nil, fmt.Errorf("failed to parse nil: invalid state %d", s)
		}
	}

	return len(bs), nil, badEOI("failed to parse nil")
}

func badByte(s string, b byte, i int) error {
	return fmt.Errorf("%s: invalid byte %b at %d", s, b, i)
}

func badLen(s string, e error) error {
	return fmt.Errorf("%s: %s", s, e.Error())
}

func badEOI(s string) error {
	return fmt.Errorf("%s: unexpected end of input", s)
}
