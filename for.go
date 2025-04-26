package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}
	for i := range 3 {
		fmt.Println(i)
	}
	for i := range 100 {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		}
	}
}
