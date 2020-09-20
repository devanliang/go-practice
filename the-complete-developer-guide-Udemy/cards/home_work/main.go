package main

import "fmt"

func main() {

	number := []int{}
	i := 0
	for i < 11 {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
		number = append(number, i)
		i++
	}

}
