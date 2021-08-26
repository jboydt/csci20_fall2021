// Lab 16 -- nested struct and slice fields
//
// Programmer name: __________
// Date completed:  __________

package main

import (
	"fmt"
)

// You can use main to experiment as you implement your functions. Then,
// instead of "go test *.go" simply run the program.
// Do not forget that you still must pass the unit tests.
func main() {
	myList := NewShoppingList("Boyd's Groceries")
	fmt.Printf("%v\n", myList)

	testItemOne := Item{"apples", 3, 0.75}

	myList.AddItem(testItemOne)
	fmt.Printf("%v\n", myList)

	myList.AddItem(testItemOne)
	fmt.Printf("%v\n", myList)
	fmt.Printf("%s\n", myList.ToString())
}
