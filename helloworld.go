package main

// importing the dependencies , we can import multiple like this:
import (
	"fmt"
	"math"
)

// making a new function which gives sum of 2 nums
func addTwo(x, y int) int {
	return x + y
}

// swapping strings
func swap(x, y string) (string, string) {
	return y, x
}

// main has no return type and is neccessary to run as main
func main() {
	fmt.Println("Hello World")
	fmt.Println("value of pi is :", math.Pi)
	fmt.Println(addTwo(12, 11))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
