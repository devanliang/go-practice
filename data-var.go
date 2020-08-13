package main // 可執行程式必需使用main 封包
import "fmt" // 載入內建的fmt 封包, 用來做基本輸出輸入
func main(){ // 建立main函式, 程式的進入點
	// 執行程式時, 從這裡開始
	// 輸出訊息到終端: fmt.Println(資料, 資料, ...)
	/*
	fmt.Println(3) // 整數 int
	fmt.Println(3.1415) // float64
	fmt.Println("測試中文") // string
	fmt.Println(true) // bool
	fmt.Println('a') // 字符 rune
	*/
	// 變數的使用
	var x int //宣告變數
	x = 4 // 指定資料
	fmt.Println(x) // 使用變數: 用變數代替資料
	x = 10 // 指定新的資料
	fmt.Println(x)
	var f float64 = 3.1514 // 宣告順便放進資料
	fmt.Println(f)
	var s string = "一二三"
	fmt.Println(s)
	var b bool = true
	fmt.Println(b)
	var r rune = 'f'
	fmt.Println(r)
}
