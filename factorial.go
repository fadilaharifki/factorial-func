package main

import "fmt"

func factorial(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func main() {
	resultFactorial := factorial(5)
	fmt.Println(resultFactorial)
}
