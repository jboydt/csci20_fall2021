package main

import (
	"fmt"
)

func main() {
	x := 42

	switch {
	case x > 42:
		fmt.Printf("Too high\n")
	case x < 42:
		fmt.Printf("Too low\n")
	case x == 42:
		fmt.Printf("That is the answer!\n")
	default:
		fmt.Printf("Impossible?\n")
	}
}
