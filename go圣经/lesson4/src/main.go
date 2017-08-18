package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func ArrIsArgs(arr *[4]int) {
	arr[0] = 100
}

func TestMap() {
	var mymap map[string]int
	mymap = make(map[string]int)
	mymap1 := map[string]int{"hello": 1, "world": 2}
	mymap["wangwenxue"] = 27
	mymap["yechase"] = 30

	fmt.Println(mymap["wangwenxue"], mymap1["hello"])

}

func TestStruct() {
	fmt.Println("-----这个段代码主要用来测试结构体的定义和使用-----")
	type myStruct struct {
		name string
		age  int
		male string
	}
	var mystest myStruct
	var mystruct1 = myStruct{"wenxuwan", 12, "man"}
	var mystruct2 = myStruct{name: "mashijie", age: 12, male: "man"}
	fmt.Println(mystruct1, mystruct2)
	mystest.name = "wenxuwan"
	mystest.age = 28
	mystest.male = "man"
	var myStructPoint *myStruct = &mystest
	memberPoint := &mystest.age
	*memberPoint = 200
	fmt.Println(myStructPoint.name, myStructPoint.age, myStructPoint.male)
	fmt.Println("-----这个段代码主要用来测试结构体的嵌入和匿名函数的使用-----")
	type Animal struct {
		method  string
		species string
		male    string
	}

	type Bird struct {
		Animal
		fly bool
	}

	type Tiger struct {
		Animal
		yanchi int
		male   string
	}
	var tigers Tiger
	tigers.male = "nan"
	tigers.method = "pao"
	fmt.Println(tigers.Animal.method, tigers.male)
}

func testBasic() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(values[:0])
	values = append(values[:0], 10)
	fmt.Println(cap(values), len(values))
}

func TestSLice(myslice []int) {
	//fmt.Println(cap(myslice), len(myslice))
	for i := 0; i < 8; i++ {
		myslice = append(myslice, i)
	}
	//fmt.Println(myslice)
}

func TestSLice2(myslice []int) {
	//fmt.Println(cap(myslice), len(myslice))
	myslice[0] = 10
	//fmt.Println(myslice)
}

func TestSLice3(myslice *[]int) {
	for i := 0; i < 8; i++ {
		*myslice = append(*myslice, i)
	}
}

func TestJson() {
	type Book struct {
		Name   string
		Year   int `json:"released"`
		Author []string
	}

	var books = []Book{
		{Name: "C++ Primer", Year: 2006, Author: []string{"LippmanBarbara"}},
		{Name: "Thinking in Java", Year: 2006, Author: []string{"Bruce Eckel"}},
	}
	data, err := json.MarshalIndent(books, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
func Sum(a, b int) int {
	return a + b
}

func main() {
	testBasic()
	fmt.Printf("%T\n", Sum)
	//myslice1 := make([]int, 0, 3)
	//TestSLice3(&myslice1)
	//fmt.Println(myslice1)
	//fmt.Println("----------------")
	//TestMap()
	//TestStruct()
	//TestJson()
	//myslice := []int{1, 4, 6, 7, 2, 8, 9, 1}
	//TestSLice(myslice)
	//structTree.Sort(myslice)
	//fmt.Printf("%v", myslice)
	//poster.GetPicUrl("Transformers")
}
