package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // create channel

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c { // chan 有值後使用for range 取出來 給到 l 變數
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // (): 使用這個 func literal 的意思
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be doen")
		c <- link
		return // make sure dont't do anything inside of if function
	}

	fmt.Println(link, "is up!")
	c <- link
}
