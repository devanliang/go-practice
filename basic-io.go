package main

import "fmt"

func main() {
	// fmt.Println(3, "Hello", true)
	// &變數名稱 : 取得變數的指標(Pointer)
	/*
		var x int
		fmt.Print("請輸入一個數字")
		fmt.Scanln(&x)
		fmt.Println(x)
	*/
	/*
		var x int
		var y int
		fmt.Print("輸入第一個數字:")
		fmt.Scanln(&x)
		fmt.Print("輸入第二個數字:")
		fmt.Scanln(&y)
		var result int = x * y
		fmt.Println(result)
	*/
	var x int
	var y int
	fmt.Println("輸入兩個數字, 用空格隔開:")
	fmt.Scanln(&x, &y)
	var result int = x * y
	fmt.Println(result)
}
