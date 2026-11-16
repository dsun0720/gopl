package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("===============")
	fmt.Println(s)
	fmt.Println("===============")
	
}

func echo2() {
	fmt.Println("===============")
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("===============")
}

func echo3() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("===============")
	fmt.Println(s)
	fmt.Println("===============")
}

// echo4 使用 while 形式的 for（只有条件）
func echo4() {
	var s, sep string
	i := 1
	for i < len(os.Args) {
		s += sep + os.Args[i]
		sep = " "
		i++
	}
	fmt.Println("===============")
	fmt.Println(s)
	fmt.Println("===============")
}

// echo5 使用无限循环形式的 for（配合 break 退出）
func echo5() {
	var s, sep string
	i := 1
	for {
		if i >= len(os.Args) {
			break
		}
		s += sep + os.Args[i]
		sep = " "
		i++
	}
	fmt.Println("===============")
	fmt.Println(s)
	fmt.Println("===============")
}

func main() {
	echo1()
	echo2()
	echo3()
	echo4()
	echo5()
	fmt.Println("END")
}
