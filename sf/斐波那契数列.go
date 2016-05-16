package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	s1 := 0
	s2 := 1

	//
	return func() int {
		s3 := s1 + s2
		s1 = s2
		s2 = s3
		return s3

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
