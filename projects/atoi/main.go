package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ch string
	fmt.Println("enter number: ")
	fmt.Scan(&ch)
	for _, a := range ch {
		y, e := strconv.Atoi(string(a))
		if e == nil {
			fmt.Print(y * y)
		}
	}
}
