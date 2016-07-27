package main //包名

import "fmt" //引用

//程序入口
func main() {
   fmt.Println("Hello, World!");
   fmt.Println("我是第一个GO语言程序!");
   const one int = 10;
   const two int = 5;
   var area int;
   area = one + two;
   fmt.Printf("%d + %d = %d",one,two,area);

}

