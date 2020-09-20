package main

import (
	"fmt"
	"io"
	"os"
)

type reader struct{}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//bs := make([]byte, 99)
	//n, err := f.Read(bs)
	//fmt.Println(string(bs[:n]))

	rd := reader{}
	io.Copy(rd, f)
}

func (r reader) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("JSUT WROTE THIS MANY BYPE:", len(bs))
	return len(bs), nil
}
