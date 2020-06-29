package goresp

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDecode_SimpleString_Success(t *testing.T) {
	s := "this is a simple string"
	bs := []byte("+" + s + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)

		return
	}

	ss, e := vs[0].String()

	if e != nil {
		t.Error(e)

		return
	}

	if ss != s {
		t.Errorf("expected %s, got %s", s, ss)
	}
}

func TestDecode_SimpleString_Empty(t *testing.T) {
	s := ""
	bs := []byte("+" + s + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)

		return
	}

	ss, e := vs[0].String()

	if e != nil {
		t.Error(e)

		return
	}

	if ss != s {
		t.Errorf("expected %s, got %s", s, ss)
	}
}

func TestDecode_BulkString_Success(t *testing.T) {
	s := "this is a bulk string"
	bs := []byte("$" + strconv.Itoa(len(s)) + "\r\n" + s + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)

		return
	}

	x, e := vs[0].String()

	if e != nil {
		t.Error(e)

		return
	}

	if x != s {
		t.Errorf("expected %s, got %s", s, x)
	}
}

func TestDecode_BulkString_Empty(t *testing.T) {
	s := ""
	bs := []byte("$" + strconv.Itoa(len(s)) + "\r\n" + s + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)

		return
	}

	x, e := vs[0].String()

	if e != nil {
		t.Error(e)

		return
	}

	if x != s {
		t.Errorf("expected %s, got %s", s, x)
	}
}

func TestDecode_Integer_Success(t *testing.T) {
	i := 123
	bs := []byte(":" + strconv.Itoa(i) + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)

		return
	}

	n, e := vs[0].Integer()

	if e != nil {
		t.Error(e)

		return
	}

	if n != i {
		t.Errorf("expected %d, got %d", i, n)
	}
}

func TestDecode_Integer_Negative(t *testing.T) {
	i := -123
	bs := []byte(":" + strconv.Itoa(i) + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)

		return
	}

	n, e := vs[0].Integer()

	if e != nil {
		t.Error(e)

		return
	}

	if n != i {
		t.Errorf("expected %d, got %d", i, n)
	}
}

func TestDecode_Float_Success(t *testing.T) {
	f := 420.69
	bs := []byte("." + fmt.Sprintf("%f", f) + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)

		return
	}

	n, e := vs[0].Float()

	if e != nil {
		t.Error(e)

		return
	}

	if n != f {
		t.Errorf("expected %f, got %f", f, n)
	}
}

func TestDecode_Float_Negative(t *testing.T) {
	f := -420.69
	bs := []byte("." + fmt.Sprintf("%f", f) + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)

		return
	}

	n, e := vs[0].Float()

	if e != nil {
		t.Error(e)

		return
	}

	if n != f {
		t.Errorf("expected %f, got %f", f, n)
	}
}

func TestDecode_Error_Success(t *testing.T) {
	s := "this is an error"
	bs := []byte("-" + s + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)
	}

	e, err := vs[0].Error()

	if err != nil {
		t.Error(err)

		return
	}

	if e.Error() != s {
		t.Errorf("expected %+v, got %+v", s, e)
	}
}

func TestDecode_Error_Empty(t *testing.T) {
	s := ""
	bs := []byte("-" + s + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)
	}

	e, err := vs[0].Error()

	if err != nil {
		t.Error(err)

		return
	}

	if e.Error() != s {
		t.Errorf("expected %+v, got %+v", s, e)
	}
}

func TestDecode_Null_Success(t *testing.T) {
	bs := []byte{'$', '-', '1', '\r', '\n'}
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)

		return
	}

	e = vs[0].Null()

	if e != nil {
		t.Error(e)
	}
}

func TestDecode_Nil_Success(t *testing.T) {
	bs := []byte{'!', '\r', '\n'}
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)

		return
	}

	e = vs[0].Null()

	if e != nil {
		t.Error(e)
	}
}

func TestDecode_Array_Success(t *testing.T) {
	i := 123
	s := "test"
	bs := []byte("*2\r\n+" + s + "\r\n:" + strconv.Itoa(i) + "\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)
	}

	vs, e = vs[0].Array()

	if e != nil {
		t.Error(e)

		return
	}

	l = len(vs)

	if l != 2 {
		t.Errorf("expect 2, got %d", l)

		return
	}

	s2, e := vs[0].String()

	if e != nil {
		t.Error(e)

		return
	}

	i2, e := vs[1].Integer()

	if e != nil {
		t.Error(e)

		return
	}

	if s2 != s {
		t.Errorf("expected %s, got %s", s, s2)
	}

	if i2 != i {
		t.Errorf("expected %d, got %d", i, i2)
	}
}

func TestDecode_Array_Empty(t *testing.T) {
	bs := []byte("*0\r\n")
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expect 1, got %d", l)
	}

	vs, e = vs[0].Array()

	if e != nil {
		t.Error(e)

		return
	}

	l = len(vs)

	if l != 0 {
		t.Errorf("expect 0, got %d", l)
	}
}
