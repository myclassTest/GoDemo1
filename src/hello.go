package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	if i := 0; i <= 0 {
		fmt.Println(i)
	}
	/*fmt.Println("一Byte:" + B + "字节")
	fmt.Println("一KB:" + KB + "字节")
	fmt.Println("一MB:" + MB + "字节")
	fmt.Println("一GB:" + GB + "字节")
	fmt.Println("一TB:" + TB + "字节")
	fmt.Println("一PB:" + PB + "字节")
	fmt.Println("一EB:" + EB + "字节")
	fmt.Println("一ZB:" + ZB + "字节")
	fmt.Println("一YB:" + YB + "字节")*/
}
