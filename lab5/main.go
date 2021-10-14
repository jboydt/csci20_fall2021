// Lab 5 -- functions
//
// Programmer name: __________
// Date completed:  __________

package main

import "fmt"

// HelloWorld prints the message "Hello world!" (followed by a newline)
// parameter(s): none
// return value(s): none
func HelloWorld() {
  // 1. Print "Hello world!" (followed by a newline)
}

// GreetUser prints the message "Hello PERSON!" (followed by a newline)
// parameter(s): a person's name (string)
// return value(s): none
func GreetUser(person string) {
  // 2. Print "Hello PERSON!" (followed by a newline; PERSON replaced by parameter)
}

// EstimateAge prints the message "I estimate your age to be AGE." (followed by a newline)
// parameter(s): a person's birth year (int)
// return value(s): none
func EstimateAge(birthYear int) {
  // 3a. Compute a person's estimated age, based on birthYear (2021 - birthYear)
  // 3b. Print "I estimate your age to be AGE." (followed by a newline;
  //     AGE replaced by computed value)
}

// ComputeAverage prints the message "Your average is AVERAGE%." (followed by a newline)
// parameter(s): points earned (float64), points possible (float64)
// return value(s): none
func ComputeAverage(pointsEarned, pointsPossible float64) {
  // 4a. Compute the average (pointsEarned / pointsPossible * 100.0)
  // 4b. Print "Your average is AVERAGE%." (followed by a newline,
  //     AVERAGE replaced by computed value, formatted to two decimal places)
}

func main() {
  HelloWorld()

  // 5. Declare a variable (string), prompt for, and read in a person's name
  // 6. Invoke GreetUser -- pass in the name variable as its argument

  // 7. Declare a variable (int), prompt for, and read in a person's birth year
  // 8. Invoke EstimateAge -- pass in the birth year variable as its argument

  // 9. Declare a variable (float64), prompt for, and read in points earned
  // 10. Declare a variable (float64), prompt for, and read in points possible
  // 11. Invoke ComputeAverage -- pass in points earned/points possible variables
  //     as arguments
}
