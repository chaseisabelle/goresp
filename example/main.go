package main

import "github.com/chaseisabelle/goresp"

func main() {
	input := "+foo\r\n"
	input += "$3\r\nbar\r\n"
	input += "$-1\r\n"
	input += ":123\r\n"
	input += "-error\r\n"
	input += "*3\r\n+a\r\n+b\r\n+c\r\n"
	input += "!\r\n"
	input += ".3.14\r\n"

	vals, err := goresp.Decode([]byte(input))

	if err != nil {
		panic(err)
	}

	simpleString := vals[0]
	bulkString := vals[1]
	null := vals[2]
	integer := vals[3]
	error1 := vals[4]
	array := vals[5]
	null2 := vals[6]
	decimal := vals[7]

	str, err := simpleString.String()

	if err != nil {
		panic(err)
	}

	println(str) //<< should print foo

	str, err = bulkString.String()

	if err != nil {
		panic(err)
	}

	println(str) //<< should print bar

	err = null.Null()

	if err != nil {
		panic(err)
	}

	i, err := integer.Int()

	if err != nil {
		panic(err)
	}

	println(i) //<< should print 3.14

	e, err := error1.Error()

	if err != nil {
		panic(err)
	}

	println(e.Error()) //<< should print error

	a, err := array.Array()

	if err != nil {
		panic(err)
	}

	ss, err := goresp.Strings(a)

	if err != nil {
		panic(err)
	}

	for _, s := range ss {
		println(s) //<< should print a, b, c
	}

	err = null2.Null()

	if err != nil {
		panic(err)
	}

	f, err := decimal.Float64()

	if err != nil {
		panic(err)
	}

	println(f) //<< should print 3.14

	output, err := goresp.Encode(vals)

	if err != nil {
		panic(err)
	}

	println(string(output)) //<< should be identical to the input
}
