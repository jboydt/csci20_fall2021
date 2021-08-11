// Lab 10 -- arrays
//
// Programmer name: __________
// Date completed:  __________

package main

import "fmt"

// IntArray explores usage of an integer array.
func IntArray() {
	// 1. Declare an array of 10 integers named "numbers" in its zero state

	// 2. Use %v to print the array in its default format

	// 3. Print the length of the array

	// 4. Use array indexing to individually populate
	//    the array with the values 10 (index 0), 20, 30, 40
	//    50, 60, 70, 80, 90, 100 (requires 10 separate lines of code)


	// 5. Use %v to print the updated array in its default format

}

// FloatArray explores usage of a float array.
func FloatArray() {
	// 6. Declare an array of 5 float64 named "scores" in its zero state

	// 7. Use a "counting" loop to prompt for and read in values for the array

	// 8. Use %v to print the array in its default format

	// 9. Use a "range" loop (ignoring index) to compute the sum of the scores

	// 10. Print the sum of scores to two decimal places

	// 11. Compute the average of scores

	// 12. Print the average of scores to two decimal places

}

// StringArray explores usage of a string array.
func StringArray() {
	// 13. Declare an array of 3 string named "greetings" -- pre-populate
	//    the array with the values "hi", "hello", and "howdy"

	// 14. Use a "range" loop (including index) to print each greeting, preceded
	//    by its position/index in the array
	//
	//    EXAMPLE:     [0] Howdy
	//                 [1] Hello
	//                 [2] Hi
	//
	//     BETTER:     [1] Howdy
	//                 [2] Hello
	//                 [3] Hi

}

// ParallelArrays explores usage of arrays "in parallel"
// (i.e., the arrays have related data).
func ParallelArrays() {
	// 15. Declare an array of 5 string named "products" -- pre-populate
	//    the array with "bread", "milk", "eggs", "butter", "sugar"

	// 16. Declare an array of 5 float64 named "prices" -- pre-populate
	//    the array with reasonable prices for each item; prices should
	//    "parallel" the products (i.e., prices[0] is the price for products[0])
	//
	//    (Sample prices: http://www.thepeoplehistory.com/pricebasket.html)

	// 17. Use a "counting loop" to print the products with their prices
	//
	//    EXAMPLE (prices not from data source):
	//                  bread: $1.11
	//                  milk: $2.22
	//                  eggs: $3.33
	//                  butter: $4.44
	//                  sugar: $5.55
	//
	//    BETTER:
	//                  bread: $1.11
	//                   milk: $2.22
	//                   eggs: $3.33
	//                 butter: $4.44
	//                  sugar: $5.55

}

func main() {

	IntArray()

	fmt.Printf("\n")

	// 18. Call FloatArray

	fmt.Printf("\n")

	//  19. Call StringArray

	fmt.Printf("\n")

	// 20. Call ParallelArrays

	fmt.Printf("\n")
}
