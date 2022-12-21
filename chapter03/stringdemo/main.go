package main

import "fmt"

func main() {
	// 字符串一旦被赋值就不能再修改了
	// str := "test"
	// str[0] = 'a'

	str2 := "abc\nabc"
	fmt.Println(str2)

	// 使用的反引号
	str3 := `package main

	import "fmt"
	
	func main() {
		// 该区域的数据值可以在同一类型范围内不断变化
		var i int = 10
		i = 30
		i = 50
		fmt.Println("i = ", i)
		// i = 1.2 // int，原因是不能改变数据类型
	
		// 变量再同一个作用域（在一个函数或者在代码块）内不能重名
		// var i int = 59
		// i := 99 // 不能再写下面这个了
	
		// 变量 = 变量名 + 值 + 数据类型，变量的三要素
	
		// golang的变量如果没有赋初值，编译器会使用默认值
	}`
	fmt.Println(str3)

	// 字符串拼接方式
	var str = "hello world"
	str += " hello"
	fmt.Println(str)

	// 当一行字符串太长时，这样来写
	var str4 = "hello " + "world " +
		" hello " + " world"
	fmt.Println(str4)

}
