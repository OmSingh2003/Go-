// Command Line Arguments
// echo program
// Print command Line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, space string
	for i := 1; i < len(os.Args); i++ {
		s += space + os.Args[i]
		space = " "
	}
	var st, sp string
	for _, str := range os.Args[1:] {
		st += sp + str
		sp = " "
	}
	fmt.Println(s)
	fmt.Println(st)
}
