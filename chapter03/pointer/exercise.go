package main

import "fmt"

func main() {
	var num int = 9
	fmt.Printf("num address = %v", &num)

	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Println("\nnum = ", num)

	var a int = 300
	var ptr *int = a // 错误，值赋给一个指针变量了

	var a int = 300
	var ptr *float32 = &a // 错误，类型不匹配

	var a int = 300
	var b int = 400
	var ptr *int = &a                                   // ok
	*ptr = 100                                          // a = 100
	ptr = &b                                            // // ok
	*ptr = 200                                          // b = 200
	fmt.Printf("a = %d, b = %d, *ptr = %d", a, b, *ptr) // a = 100, b = 200, *ptr = 200
}
