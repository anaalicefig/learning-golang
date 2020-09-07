package main

import "fmt"

type amount int

var x amount

func main() {
	fmt.Println("X:", x)
	fmt.Printf("Type of X: %T\n", x)

	x = 42
	fmt.Println("X:", x)
}
