package main

import (
	"fmt"
	"sort"
)

func main() {
	// var m map[int]map[int]string
	// m = make(map[int]map[int]string)
	// a, ok := m[2][1]
	// if !ok {
	//  m[2] = make(map[int]string)
	// }
	// m[2][1] = "Goolge"
	// a, ok = m[2][1]
	// fmt.Println(a, ok)

	// sm := make([]map[int]string, 5)
	// for i := range sm {
	//  sm[i] = make(map[int]string, 1)
	//  sm[i][1] = "OK"
	//  fmt.Println(sm[i])
	// }
	// fmt.Println(sm)
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m1)
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)

	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}
