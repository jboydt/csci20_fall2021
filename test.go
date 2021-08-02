package main

import (
	"fmt"
	"math/rand"
)

func Riddle() {
	fmt.Printf("riddle me ")
	return
	fmt.Printf("this")
}

func main() {
	// for i := 1.0; i < 2.0; i += 0.1 {
	// 	fmt.Printf("i -> %.1f, rounded -> %.1f\n", i, math.Round(i))
	// }

	rand.Seed(1)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rand.Intn(6) + 1)
	}
	fmt.Printf("\n\n")
}
