package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 70
	var str string = " Hello, World"
	fmt.Println(str + string(a))       // this converts a to UNICODE and add it
	fmt.Println(str + strconv.Itoa(a)) // strconv.Itoa converts a which is int to string and then Concat
}
