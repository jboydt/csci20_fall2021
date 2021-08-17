package main

import "fmt"

func main() {
	yourLuckyNumbers := [4]int{37, 41, 43, 47}
	myLuckyNumbers := append(yourLuckyNumbers[:2], 17, 19)

	fmt.Printf("%v\n", myLuckyNumbers)
}
