package goresp

func Encode(vs []Value) ([]byte, error) {
	bs := make([]byte, 0)

	for _, v := range vs {
		t, e := v.Encode()

		if e != nil {
			return nil, e
		}

		bs = append(bs, t...)
	}

	return bs, nil
}
