package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf strings.Builder
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%c\t%d\n", i, r, r)
	}

	// "program" in Japanese katakana
	s = "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
	fmt.Println(string(r))

	fmt.Println(string(65))      // "A", not "65"
	fmt.Println(string(0x4eac))  // "京"
	fmt.Println(string(1234567)) // \uFFFD
	fmt.Println(basename("a/b.c.go"))
	fmt.Println(comma("1234567890"))
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
	fmt.Println(math.Pi)
}

// Preserve everything before last '.'.
