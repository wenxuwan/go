package main

import (
	"fmt"
	"os"
	"popcount"
	"strconv"
	"time"
	"weightconv"
)

func main() {
	//2.1的习题
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		//2.2的习题
		p := weightconv.Pound(t)
		k := weightconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n", p, weightconv.PToK(p), k, weightconv.KToP(k))

	}
	//2.3的习题
	fmt.Println("----------------2.3-----------------")
	t1 := time.Now()
	fmt.Println(t1)
	sum := popcount.PopCountFor(6562342)
	t2 := time.Now()
	fmt.Println(t2)
	fmt.Println("The PopCountFor function costs ", t2.Sub(t1), "to complete", "and the number of 1 is", sum)

	t3 := time.Now()
	fmt.Println(t3)
	sum2 := popcount.PopCount(6562342)
	t4 := time.Now()
	fmt.Println(t4)

	fmt.Println("The PopCount function costs ", t4.Sub(t3), "to complete", " and the number of 1 is", sum2)

	//2.4的代码
	fmt.Println("----------------2.4-----------------")
	t5 := time.Now()
	fmt.Println(t5)
	sum3 := popcount.PopCountShift(6562342)
	t6 := time.Now()
	fmt.Println(t6)
	fmt.Println("The PopCountShift function costs ", t6.Sub(t5), "to complete", " and the number of 1 is", sum3)

	fmt.Println("----------------2.5-----------------")
	t7 := time.Now()
	fmt.Println(t7)
	sum4 := popcount.PopCountUseBit(8)
	t8 := time.Now()
	fmt.Println(t8)
	fmt.Println("The PopCountShift function costs ", t6.Sub(t5), "to complete", " and the number of 1 is", sum4)
}
