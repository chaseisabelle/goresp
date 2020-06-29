package goresp

import (
	"errors"
	"testing"
)

func TestEncode_SimpleString_Success(t *testing.T) {
	ss1 := NewSimpleString("foo")
	bs, e := Encode([]Value{ss1})

	if e != nil {
		t.Error(e)

		return
	}

	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expected 1, got %d", l)

		return
	}

	ss2 := vs[0]
	_, s1 := ss1.String()
	_, s2 := ss2.String()

	if s1 != s2 {
		t.Errorf("expected %s, got %s", s1, s2)
	}
}

func TestEncode_BulkString_Success(t *testing.T) {
	bs1 := NewBulkString("foo")
	bs, e := Encode([]Value{bs1})

	if e != nil {
		t.Error(e)

		return
	}

	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expected 1, got %d", l)

		return
	}

	bs2 := vs[0]
	_, s1 := bs1.String()
	_, s2 := bs2.String()

	if s1 != s2 {
		t.Errorf("expected %s, got %s", s1, s2)
	}
}

func TestEncode_Integer_Success(t *testing.T) {
	i1 := NewInteger(420)
	bs, e := Encode([]Value{i1})

	if e != nil {
		t.Error(e)

		return
	}

	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expected 1, got %d", l)

		return
	}

	i2 := vs[0]
	_, n1 := i1.Integer()
	_, n2 := i2.Integer()

	if n1 != n1 {
		t.Errorf("expected %d, got %d", n1, n2)
	}
}

func TestEncode_Error_Success(t *testing.T) {
	e1 := NewError(errors.New("test"))
	bs, e := Encode([]Value{e1})

	if e != nil {
		t.Error(e)

		return
	}

	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expected 1, got %d", l)

		return
	}

	e2 := vs[0]
	_, s1 := e1.Error()
	_, s2 := e2.Error()

	if s1 != s2 {
		t.Errorf("expected %+v, got %+v", s1, s2)
	}
}

func TestEncode_Null_Success(t *testing.T) {
	n1 := NewNull()
	bs, e := Encode([]Value{n1})

	if e != nil {
		t.Error(e)

		return
	}

	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expected 1, got %d", l)

		return
	}

	n2 := vs[0]
	e = n2.Null()

	if e != nil {
		t.Errorf("expected nil, got %+v", e)
	}
}

func TestEncode_Array_Success(t *testing.T) {
	n1 := NewNull()
	a1 := NewArray([]Value{n1})
	bs, e := Encode([]Value{a1})

	if e != nil {
		t.Error(e)

		return
	}

	vs, e := Decode(bs)

	if e != nil {
		t.Error(e)

		return
	}

	l := len(vs)

	if l != 1 {
		t.Errorf("expected 1, got %d", l)

		return
	}

	a2 := vs[0]
	vs, e = a2.Array()

	if e != nil {
		t.Error(e)

		return
	}

	l = len(vs)

	if l != 1 {
		t.Errorf("expected 1, got %d", l)

		return
	}

	n2 := vs[0]
	e = n2.Null()

	if e != nil {
		t.Errorf("expected nil, got %+v", e)
	}
}
