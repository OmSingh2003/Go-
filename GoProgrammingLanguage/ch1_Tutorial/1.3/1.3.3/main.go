package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		var num1, num2 int
		n, err := fmt.Sscanf(line, "%d,%d", &num1, &num2)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in parsing string: %v.Check format.\n", err)
			continue
		}
		if n != 2 {
			fmt.Fprintf(os.Stderr, "Input number need to be 2:%d,%s\n", n, line)
			continue
		}
		sum := num1 + num2
		fmt.Printf("Sum of 2 number %d and %d  is : %d \n", num1, num2, sum)
	}
}
