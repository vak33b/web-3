package main

import "fmt"

func main() {
	var str string
	var newstr string
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		newstr += string(str[i])
		if i != len(str)-1 {
			newstr += "*"
		}
	}
	fmt.Println(newstr)
}
