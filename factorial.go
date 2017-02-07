// factorial
package main

import (
	"fmt"
)

func factorial(number int) int {
	if number == 0 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

func main() {
	x := factorial(10)
	fmt.Println(x)
}
