package main

import "fmt"

func main() {
	// 基本迴圈
	/*
		var x int = 0
		for x < 3 {
			fmt.Println(x)
			x++
		}
	*/
	/*
		var x int
		for x = 0; x < 3; x++ {
			fmt.Println(x)
		}
	*/
	// 實例: 計算 1+2+3+...+50 的結果
	/*
		var x int = 1
		var result int = 0
		for x <= 50 {
			//fmt.Println("x:", x)
			result = result + x
			x++
		}
		fmt.Println(result)
	*/
	var x int
	var result int = 0
	for x = 1; x < 50; x++ {
		result += x
		//fmt.Println(x)
	}
	fmt.Println(result)
}
