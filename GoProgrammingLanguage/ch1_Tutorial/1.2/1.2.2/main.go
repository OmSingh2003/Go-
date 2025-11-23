// use strings.Join because strings are immutable in golang so when we iterate and add contniously add something to a string
// lets say n times so each time a new memory is allocated and our variable starts to point over there creating
// n + 1 spaces which is cleaned by garbage collector but in order to that program stops
// which makes it bad practice and ineffecient
// [java also the same problem]
// [solution : joins or string builder]
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
