// Lab 6 -- functions + testing
//
// Programmer name: __________
// Date completed:  __________

package main

import "fmt"

// HelloWorld returns the message "Hello world!".
// parameter(s): none
// return value(s): a formatted string containing "Hello world!"
func HelloWorld() string {
  // 1a. Delete the "stubbed" return statement
  // 1b. Return "Hello world!" as a formatted string
  return ""
}

// GreetUser returns the message "Hello PERSON!"
// parameter(s): a person's name (string)
// return value(s): a formatted string containing "Hello PERSON!"
func GreetUser(person string) string {
  // 2a. Delete the "stubbed" return statement
  // 2b. Return "Hello PERSON!" as a formatted string (PERSON replaced by parameter)
  return ""
}

// EstimateAge returns an estimated age based on a person's birth year
// parameter(s): a person's birth year (int)
// return value(s): a person's estimated age (int)
func EstimateAge(birthYear int) int {
  // 3a. Delete the "stubbed" return statement
  // 3b. Return a person's estimated age, based on birthYear (2021 - birthYear)
  return 0
}

// ComputeAverage returns an average (out of 100%) based on points earned
// and points possible.
// parameter(s): points earned (float64), points possible (float64)
// return value(s): a computed average (float64)
func ComputeAverage(pointsEarned, pointsPossible float64) float64 {
  // 4a. Delete the "stubbed" return statement
  // 4b. Return the average (pointsEarned / pointsPossible * 100.0)
  return 0.0
}

func main() {
  var name string
  fmt.Printf("Please enter your name: ")
  fmt.Scanf("%s\r\n", &name)
  fmt.Printf("%s\n", GreetUser(name))

  // 7. Declare a variable (int), prompt for, and read in a person's birth year
  // 8. Invoke and print the result of EstimateAge --
  //    pass in the birth year variable as its argument

  // 9. Declare a variable (float64), prompt for, and read in points earned
  // 10. Declare a variable (float64), prompt for, and read in points possible
  // 11. Invoke and print the result of ComputeAverage --
  //     pass in points earned/points possible variables as arguments
}
