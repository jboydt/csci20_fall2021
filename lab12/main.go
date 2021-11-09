// Lab 12 -- slice expressions
//
// Programmer name: __________
// Date completed:  __________

package main

import (
	"fmt"
	"sort"
)

// LowestThree sorts a slice of integers (ascending) and returns a slice
// containing the lowest three (i.e., first three) integers.
// If the slice contains 3 or fewer items, returns the sorted slice without
// removing any elements.
//
// uses sort.Ints
func LowestThree(numbers []int) []int {
	return []int{}
}

// HighestThree sorts a slice of integers (ascending) and returns a slice
// containing the highest three (i.e., last three) integers.
// If the slice contains 3 or fewer items, returns the sorted slice without
// removing any elements.
//
// uses sort.Ints
func HighestThree(numbers []int) []int {
	return []int{}
}

// DropHighestAndLowest sorts a slice of float64 (ascending) and returns
// a slice containing all of the float64 except the first and last.
// If the slice contains 2 or fewer items, returns the sorted slice without
// removing any elements.
//
// uses sort.Float64s
func DropHighestAndLowest(values []float64) []float64 {
	return []float64{}
}

// FirstHalf returns the first half of a slice of strings. The first
// half is computed as the strings from the beginning of the slice up to
// (but not including) the index computed as the length of the slice divided
// by two.
// If the list is empty, returns the list without doing any work.
func FirstHalf(list []string) []string {
	return []string{}
}

// SecondHalf returns the second half of a slice of strings. The second
// half is computed as the strings from the index at the halfway point
// (computed as the length of the slice divided by two) through the end
// of the slice.
// If the list is empty, returns the list without doing any work.
func SecondHalf(list []string) []string {
	return []string{}
}

// StringInCenter returns a slice containing the single string that is at the
// center of the slice. The center is computed as the length of the slice
// divided by two.
// If the list is empty, returns the list without doing any work.
func StringInCenter(list []string) []string {
	return []string{}
}

// You can use main to experiment as you implement your functions. An
// example below. Then, instead of "go test *.go" simply run the program.
// Do not forget that you still must pass the unit tests.
func main() {
	// EXAMPLE EXPERIMENT
	numbers := [10]int{5, 4, 3, 2, 1}

	// Make a copy and sort it so we can check the results against the
	// sorted original
	sortedNumbers := numbers
	sort.Ints(sortedNumbers[:])
	fmt.Printf("sorted:        %v\n", sortedNumbers)

	// Display results
	fmt.Printf("lowest three:  %v\n", LowestThree(numbers[:]))
	fmt.Printf("highest three: %v\n", HighestThree(numbers[:]))
}
