package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func main() {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	for i, v := range a {
		fmt.Printf("[%d] %d\n", i, v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)    // "[1 2 3]"
	fmt.Println(r[2]) // "0"

	q1 := [...]int{1, 2, 3, 5, 6}
	fmt.Printf("%T\t%[1]v\n", q1)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	r1 := [...]int{0: -1, 3: 1, 10: 10}
	fmt.Printf("%T\t%[1]v\n", r1) // "[11]int"

	c1 := sha256.Sum256([]byte("sun"))
	c2 := sha256.Sum256([]byte("SUN"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	s := arr[2:5]                                                // len=3, cap=5
	fmt.Printf("指针:%p, len:%d, cap:%d\n", &s[0], len(s), cap(s)) // 指针地址，len=3，cap=5
	s = append(s, 10)
	fmt.Println(arr, s)
}
