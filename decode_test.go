package goresp

import (
	"errors"
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

	n, e := vs[0].Int()

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

	n, e := vs[0].Int()

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

	n, e := vs[0].Float64()

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

	n, e := vs[0].Float64()

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

func TestDecode_Bool_Success(t *testing.T) {
	bs := []byte{'<', '\r', '\n', '>', '\r', '\n'}
	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 2 {
		t.Errorf("expect 2, got %d", l)

		return
	}

	b1, e := vs[0].Bool()

	if e != nil {
		t.Error(e)
	}

	b2, e := vs[1].Bool()

	if e != nil {
		t.Error(e)
	}

	if b1 != false {
		t.Error(errors.New("expected false, got true"))
	}

	if b2 != true {
		t.Error(errors.New("expected true, got false"))
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

	i2, e := vs[1].Int()

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

func TestDecode_Array_Nested(t *testing.T) {
	bs := []byte("*2\r\n*1\r\n+foo\r\n*1\r\n+bar\r\n")
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

	a1, e := vs[0].Array()

	if e != nil {
		t.Error(e)

		return
	}

	a2, e := vs[1].Array()

	if e != nil {
		t.Error(e)

		return
	}

	s1, e := a1[0].String()

	if e != nil {
		t.Error(e)

		return
	}

	s2, e := a2[0].String()

	if e != nil {
		t.Error(e)

		return
	}

	if s1 != "foo" {
		t.Errorf("expected foo, got %s", s1)
	}

	if s2 != "bar" {
		t.Errorf("expected bar, go %s", s2)
	}
}
