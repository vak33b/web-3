package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)
	maxim := -10000
	for _, a := range str {
		y, e := strconv.Atoi(string(a))
		if e == nil {
			if y > maxim {
				maxim = y
			}
		}
	}
	fmt.Println(maxim)
}
