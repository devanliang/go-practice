package main

import "fmt"

func main() {
	// 基本語法練習
	/*
		if false {
			fmt.Println("Go")
		} else {
			fmt.Println("Not Go")
		}
	*/
	// 簡易情境: ATM 檢測輸入金額是否正確
	var money int
	fmt.Println("請問領多少錢")
	fmt.Scanln(&money)
	if money < 100 {
		fmt.Println("Too few")
	} else if money <= 10000 {
		fmt.Println("OK")
	} else {
		fmt.Println("Too Much")
	}
	fmt.Println("執行完畢!!")
}
