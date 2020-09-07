package main

import "fmt"

type amount int

var x amount
var y int

func main() {
	fmt.Println("X:", x)
	fmt.Printf("Type of X: %T\n", x)

	x = 42
	fmt.Println("X:", x)

	fmt.Println()

	y = int(x)

	fmt.Println("Y:", y)
	fmt.Printf("Type of Y: %T\n", y)
}
