// Lab 11 -- slices
//
// Programmer name: __________
// Date completed:  __________

package main

import (
  "fmt"
  "strings"
)

// IntSlice explores usage of an integer slice.
func IntSlice() {
  // 1. Create an empty slice of integers named "numbers"

  // 2. Use %v to print the slice in its default format

  // 3. Print the length of the slice

	// 4. Use the append function to populate
	//    the slice with the values 10, 20, 30, 40, 50
	//    (can be done in 1 line of code, but OK to use more if desired)

  // 5. Use %v to print the slice in its default format

  // 6. Print the length of the slice

}

// FloatSlice explores usage of a float slice.
func FloatSlice() {
	// 7. Create an empty slice of float64 named "temperatures"

  // 8. Use %v to print the slice in its default format

	// 9. Declare a float64 variable named "input"

  // 10. Use a "counting" loop to prompt for and read in 5 values for "temperatures",
	//    appending each to the "temperatures" slice
	//    (NOTE: you will need your "input" variable)

  // 11. Use a "range" loop to compute the sum of the temperatures
	var sum float64

  // 12. Print the sum of temperatures to two decimal places

  // 13. Compute the average of temperatures

  // 14. Print the average of temperatures

}

// StringSlice explores usage of a string slice.
func StringSlice() {
	// 15. Declare an ARRAY of 3 string named "greetings" -- pre-populate
	//    the array with the values "hi", "hello", and "howdy"

  // 16. Use the make function to create a slice of strings, of equal length to
  //    "greetings", name the slice "newGreetings"


	fmt.Printf("greetings: %v\n", greetings)
	fmt.Printf("new greetings (before update): %v\n", newGreetings)

  // 17. Use a "range" loop to do the following:
	//     - convert each string in "greetings" to all upper-case (use strings.ToUpper)
	//     - assign the upper-case greeting to an index in the newGreetings slice
	//
	//     NOTE: greetings[0] is "hi", newGreetings[0] should be "HI" at the end
	//           greetings[1] is "hello", newGreetings[1] should be "HELLO" at the end
	//           greetings[2] is "howdy", newGreetings[2] should be "HOWDY" at the end


  fmt.Printf("new greetings (after update): %v\n", newGreetings)
}

func main() {

	// 18. Call IntSlice

  fmt.Printf("\n")

	// 19. Call FloatSlice

  fmt.Printf("\n")

	// 20. Call StringSlice

  fmt.Printf("\n")
}
