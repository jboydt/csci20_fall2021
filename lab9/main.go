// Lab 9 -- repetition/looping
//
// Programmer name: __________
// Date completed:  __________

package main

import "fmt"

// GetEvens returns a string containing the even numbers from start up to and
// including end. The numbers are returned as a list separated by commas
// (no comma after the last value). Zero is considered an even number.
//
// EX: GetEvens(1, 5) returns "2,4"
//     GetEvens(0, 5) returns "0,2,4"
//     GetEvens(1, 1) returns ""
//     GetEvens(0, 0) returns "0"
func GetEvens(start, end int) string {
	return ""
}

// GetOdds returns a string containing the odd numbers from start up to and
// including end. The numbers are returned as a list separated by commas
// (no comma after the last value). Zero is considered an even number.
//
// EX: GetOdds(1, 5) returns "1,3,5"
//     GetOdds(0, 5) returns "1,3,5"
//     GetOdds(1, 1) returns "1"
//     GetOdds(0, 0) returns ""
func GetOdds(start, end int) string {
  return ""
}

// FizzBuzz returns a string containing a solution to the Fizz Buzz challenge
// from 1 up to and including max. Each output is separated by a comma
// (no comma after the last output).
// https://leetcode.com/problems/fizz-buzz/description/
//
// EX: FizzBuzz(5) returns "1,2,Fizz,4,Buzz"
//     FizzBuzz(15) returns "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz"
func FizzBuzz(max int) string {
	return ""
}

// You can use main to experiment as you implement your functions. A few
// examples below. Then, instead of "go test *.go" simply run the program.
// Do not forget that you still must pass the unit tests.
func main() {
	fmt.Printf("%s\n\n", GetEvens(1, 5))
	fmt.Printf("%s\n\n", GetOdds(1, 5))
	fmt.Printf("%s\n\n", FizzBuzz(5))
}
