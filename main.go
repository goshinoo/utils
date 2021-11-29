package main

import (
	"fmt"
	"utils/streamBuf"
)

func main() {
	//Encode 编码
	//使用样例
	s := streamBuf.NewEncodeStream()
	var (
		str string = "1234"
		a   int8   = 1
		b   uint8  = 1
		c   int16  = 2
		d   int16  = 2
		e   int32  = 3
		f   int32  = 3
		g   int64  = 4
		h   int64  = 4
	)
	_ = s.Encode(&str)
	_ = s.Encode(&a)
	_ = s.Encode(&b)
	_ = s.Encode(&c)
	_ = s.Encode(&d)
	_ = s.Encode(&e)
	_ = s.Encode(&f)
	_ = s.Encode(&g)
	_ = s.Encode(&h)
	fmt.Println(s) //&{{[0 0 0 0 0 0 0 5 49 50 51 52 0 1 1 0 2 0 2 0 0 0 3 0 0 0 3 0 0 0 0 0 0 0 4 0 0 0 0 0 0 0 4] utf8 B 43}}

	//Decode 解码
	//使用样例
	s1 := streamBuf.NewDecodeStream([]byte{0, 0, 0, 0, 0, 0, 0, 5, 49, 50, 51, 52, 0, 1, 1, 0, 2, 0, 2, 0, 0, 0, 3, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 4})
	var (
		str1 string
		a1   int8
		b1   uint8
		c1   int16
		d1   int16
		e1   int32
		f1   int32
		g1   int64
		h1   int64
	)
	_ = s1.Decode(&str1)
	_ = s1.Decode(&a1)
	_ = s1.Decode(&b1)
	_ = s1.Decode(&c1)
	_ = s1.Decode(&d1)
	_ = s1.Decode(&e1)
	_ = s1.Decode(&f1)
	_ = s1.Decode(&g1)
	_ = s1.Decode(&h1)

	fmt.Println(str1, a1, b1, c1, d1, e1, f1, g1, h1) //1234 1 1 2 2 3 3 4 4
}
