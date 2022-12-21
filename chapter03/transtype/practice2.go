package main

import (
	"fmt"
)

func main() {
	var n1 int32 = 12
	var n3 int8
	var n4 int8
	n4 = int8(n1) + 127 // 编译通过，但是 结果不是127+12，溢出问题
	n3 = int8(n1) + 128 // 编译不过
	fmt.Println(n3)
}
