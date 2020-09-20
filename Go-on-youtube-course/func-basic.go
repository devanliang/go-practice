package main

import "fmt"

func test(msg string) {
	fmt.Println(msg)
}

func add(n1 int, n2 int) {
	var result int = n1 + n2
	fmt.Println(result)
}

func sum(end int) {
	var result int = 0
	var x int
	for x = 1; x <= end; x++ {
		result += x
	}
	fmt.Println(result)
}

func main() {
	// 基本函式語法演練
	/*
		test("a")
		test("你好")
		add(5, 7)
		add(-5, 10)
	*/
	// 計算 1+2+3+4+5+...+x 函式包裝
	sum(100)
}
