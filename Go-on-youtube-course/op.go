package main

import "fmt"

func main() {
	// 算數運算 +, -, *, /
	var x int
	x = 3*3 + 10
	fmt.Println(x)
	// 指定運算 =, +=, -=, *=, /=
	x = 5
	x = x + 1
	fmt.Println(x)
	x = 5
	x += 2
	fmt.Println(x)
	x = 5
	x /= 2
	fmt.Println(x)
	x = 5
	x %= 2
	fmt.Println(x)
	// 單元運算 ++, --
	x = 5
	x++ // x + 1
	fmt.Println(x)
	x = 5
	x-- // x - 1
	fmt.Println(x)
	// 比較運算: >, <, >=, <=, ==
	var y bool = 4 == 3
	fmt.Println(y)
	//邏輯運算 !, &&, ||
	var result bool = 4 == 3
	fmt.Println(result)
	result = !false
	fmt.Println(result)
	result = true && false // and: 兩邊都是 true, 結果就是true
	fmt.Println(result)
	result = true || false // or: 只要有一邊是 true, 結果就是true
	fmt.Println(result)
}
