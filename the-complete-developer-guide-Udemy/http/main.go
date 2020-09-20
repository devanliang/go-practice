package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err) // print error message
		os.Exit(1)                 // make sure exit the program
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body) // os.Stdout: value of type file, resp.Body: 是取得的body資料
}

func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("JSUT WROTE THIS MANY BYPE:", len(bs))
	return len(bs), nil
}
