package main

import "fmt"

func FuncValue(a, b int) int {
	return a + b
}

//匿名函数
func AnonymousFunc() func() int {
	var sum int
	childFunc := func() int {
		sum++
		decrease := 0
		fmt.Println("The sum is and decrease is", sum, decrease)
		return decrease
	}
	return childFunc

}

//可变参数
func UnstableParam(vals ...int) {
	for i, val := range vals {
		fmt.Println(i, val)
	}
}

func TestFunc() {

}

func ErrProcess() {

	defer func() {
		if p := recover(); p != nil {
			fmt.Println("internal error: %v", p)
		}
	}()
	panic("wwx")

	fmt.Println("mashijie")

}

func TteratePara() {
	for i := 0; i < 3; i++ {
		a := i
		defer func() {
			fmt.Printf("%d", a)
		}()
	}
}
func main() {
	ErrProcess()
	TteratePara()
	fmt.Println("haha")
	funcvalue := FuncValue
	funcvalue2 := func(a, b int) int {
		return a * b
	}
	myslice := []int{10, 11, 12, 13, 14}
	UnstableParam(myslice...)
	UnstableParam(1, 2, 3, 4, 5)
	var j int = 10
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}
	b := a()
	b()
	fmt.Println(funcvalue(1, 2))
	fmt.Println(funcvalue2(3, 4))
	//匿名函数
	func1 := AnonymousFunc()
	a1 := func1()
	b1 := func1()
	c1 := func1()
	fmt.Println(a1, b1, c1)
}
