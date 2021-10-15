package main

import (
	"fmt"
)

func main() {
	result := simple()
	fmt.Printf("Hello, the result is %t", result)
}

func simple() bool {
	return true
}
