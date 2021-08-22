// Lab 14 -- structs
//
// Programmer name: __________
// Date completed:  __________

package main

import (
	"fmt"
)

// 1. Declare a struct with the following:
//
// struct name: CartItem
// fields: name(string), price(float64), quantity(int)


// 2. Complete the implementation of this function
//
// CartItemString returns a string representation of a CartItem.
//
// TEMPLATE: NAME, QUANTITY @ $PRICE
// EXAMPLE:  strawberries, 24 @ $0.35
func CartItemString(item CartItem) string {
	return ""
}

// 3. Complete the implementation of this function
//
// CartItemCost returns the cost of an item (item quantity * item price).
func CartItemCost(item CartItem) float64 {
	return 0.0
}

func main() {
	var itemOne CartItem
	fmt.Printf("%s\n", CartItemString(itemOne))
	fmt.Printf("Total cost of %s: $%.2f\n", itemOne.name, CartItemCost(itemOne))

	fmt.Printf("\n")

	// 4. Create a CartItem instance named itemTwo, without using field labels
	//    name: "Apples", price: 0.25, quantity: 3


	// 5. Print itemTwo, using the CartItemString function


	// 6. Print the cost of itemTwo, using the CartItemCost function


	fmt.Printf("\n")

	// 7. Create a CartItem instance named itemThree, using field labels
	//    name: "Eggs", price: 0.75, quantity: 12


	// 8. Print itemThree, using the CartItemString function


	// 9. Print the cost of itemThree, using the CartItemCost function


	fmt.Printf("\n")

	// 10. Create a CartItem instance named itemFour, using field labels
	//    name: "Milk" (leave other fields as zero-valued)


	// 11. Print itemFour, using the CartItemString function


	// 12. Print the cost of itemFour, using the CartItemCost function

}
