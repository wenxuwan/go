package main

import "fmt"

const (
	a = iota
	b
	c = 20
	e
	d = iota
)

func StringFunc(s string) {
	fmt.Printf("%c", s[0])
	var f string = ",wenxuwan"
	s = s + f
	fmt.Println(s[0:4])
	fmt.Println(len(s))
}

func main() {
	var f float32 = 16777216 // 1 << 24
	var str string = "Hello,world 世界"
	var f1 float32 = f + 3
	r := []int32(str)
	fmt.Println(string(r[12]))
	fmt.Println(string(r[13]))
	StringFunc(str)
	fmt.Println(f, f1)
	fmt.Println(a, b, c, e, d)
}
