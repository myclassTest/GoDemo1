package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s1[:]
	fmt.Printf("%p", s2)
}
