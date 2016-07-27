package main

import (
	"fmt"
)
func main(){
	// fmt.Println(^2)
	// fmt.Println(1^2)
	// fmt.Println(1<< 10 << 10 >>10)
	// fmt.Println(6 & 11)
	// fmt.Println(6 | 11)
	// fmt.Println(6 ^ 11)
	// fmt.Println(6 &^ 11)
	// fmt.Println()
	a:=[10]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(a)
	sl:=a[5:10]
	fmt.Println(sl)
	//slice 切片
	slice1:=make([]int,10,20)
	fmt.Println(len(slice1),cap(slice1))
	slice1 = append(slice1,3,1,5,6)
	fmt.Printf("%p",slice1)
	for	i:=0;i<20;i++{
		v:=i
		fmt.Println(&v)
	}
	
	slice2:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	for index,value:=range slice2{
		fmt.Printf("Index:%d  IndexAddr:%X\t Value:%d  ValueAddr:%X  ElmentAddr:%X\n",index,&index,value,&value,&slice2[index])
	}
	for	index:=2;index<len(slice2);index++{
		fmt.Printf("Index:%d  Value:%d\n",index,slice2[index])
	}
	
	//二维切片
	slice3:=[][]int{{12,13},{20,30,40}}
	slice3[0] = append(slice3[0],40)
}