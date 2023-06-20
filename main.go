package main

import (
	"fmt"
	"testing_practice_in_golang/check"
)

func main() {
	a := check.Multiply(5, 4)
	b := check.Add(1, 2)
	fmt.Println("Product=", a, "Sum=", b)
}
