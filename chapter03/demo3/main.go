package main

import "fmt"

func main() {
	//// 如何一次性声明多个变量
	//var n1, n2 int
	//fmt.Println("n1 = ", n1, " n2 = ", n2)
	//
	//// 一次性生命多个变量的格式2
	//var n4, name, n3 = 100, "tom", 888
	//fmt.Println("n1 = ", n4, "name = ", name, "n3 = ", n3)

	// 一次性声明多个变量的方式3，同样可以使用类型推导
	n5, name, n3 := 100, "tom", 888
	fmt.Println("n5 = ", n5, "name = ", name, "n3 = ", n3)
}
