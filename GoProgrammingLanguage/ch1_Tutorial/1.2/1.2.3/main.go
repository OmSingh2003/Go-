package main

import (
	"fmt"
	"os"
)

func main() {
	for ind, str := range os.Args {
		fmt.Println(ind, " ", str)
	}
}
