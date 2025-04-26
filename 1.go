package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 70
	var str string = " Hello, World"
	fmt.Println(str + string(a))
	fmt.Println(str + strconv.Itoa(a))
}
