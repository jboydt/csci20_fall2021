// Lab 8 -- decision making
//
// Programmer name: __________
// Date completed:  __________

package main

//import "fmt"

// EncodeProduct returns the code number associated with a given product name.
// Returns -1 for unrecognized product names.
//
// Codes:
//
// apples:       1
// bananas:      2
// grapes:       3
// oranges:      4
// strawberries: 5
// all others:  -1
func EncodeProduct(productName string) int {
  return 0
}

// DecodeProduct returns the product name associated with a given code number.
// Returns "invalid" for unrecognized code numbers.
//
// Codes:
//
// 1:          apples
// 2:          bananas
// 3:          grapes
// 4:          oranges
// 5:          strawberries
// all others: invalid
func DecodeProduct(codeNumber int) string {
  return ""
}

// ComputeShippingCharge returns a shipping charge of 5.0 if the totalOwed
// is greater than or equal to the threshold, otherwise it returns a shipping
// charge of 0.0.
func ComputeShippingCharge(totalOwed, threshold float64) float64 {
  return 0.0
}

// TakeWalk implements a simple ID3 decision tree.
// Returns true if the conditions allow for a walk, otherwise returns false.
// See: https://bit.ly/3rRSF8p
//
// outlook - humidity - wind -  take walk?
// --------------------------------------
// sunny     high               no
// sunny     normal             yes
// overcast                     yes
// rainy                 strong no
// rainy                 weak   yes
// all others combinations      no
func TakeWalk(outlook, humidity, wind string) bool {
  return false
}

// PrettyPlayingCard returns a string representation of a playing card, given
// a value (1-13) and a suit (c/C, d/D, h/H, s/S).
//
// Values:
// 1:    Ace
// 2-10: 2, 3, 4, 5, 6, 7, 8, 9, 10
// 11:   Jack
// 12:   Queen
// 13:   King
// invalid value: ERROR
//
// Suits:
// c or C: Clubs
// d or D: Diamonds
// h or H: Hearts
// s or S: Spades
// invalid suit: ERROR
//
// EXAMPLES:
//
// PrettyPlayingCard(1, "c") returns "Ace of Clubs"
// PrettyPlayingCard(2, "D") returns "2 of Diamonds"
// PrettyPlayingCard(10, "h") returns "10 of Hearts"
// PrettyPlayingCard(13, "H") returns "King of Hearts"
// PrettyPlayingCard(0, "x") returns "ERROR of ERROR"
func PrettyPlayingCard(value int, suit string) string {
  return ""
}
