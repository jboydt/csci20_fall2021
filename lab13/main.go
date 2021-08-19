// Lab 13 -- slices and functions
//
// Programmer name: __________
// Date completed:  __________

package main

import (
	"strings"
)

// StringsOfLength returns a slice containing all strings
// in a given list matching a given length.
//
// uses len(string)
func StringsOfLength(list []string, length int) []string {
	return nil
}

// StringsBeginningWith returns a slice containing all strings
// in a given list that begin with a given prefix.
// Case insensitive.
//
// uses strings.ToUpper or strings.ToLower and strings.HasPrefix
func StringsBeginningWith(list []string, prefix string) []string {
	return nil
}

// StringsEndingWith returns a slice containing all strings
// in a given list that end with a given suffix.
// Case insensitive.
//
// uses strings.ToUpper or strings.ToLower and strings.HasSuffix
func StringsEndingWith(list []string, suffix string) []string {
	return nil
}

// StringsContaining returns a slice containing all strings
// in a given list that contain a given substring. Case insensitive.
//
// uses strings.ToUpper or strings.ToLower and strings.Contains
func StringsContaining(list []string, substring string) []string {
	return nil
}

// FieldsToLowercase separates a string into its fields as defined by
// the given separator. Then, it converts all of these fields to lowercase
// and returns the result.
//
// uses strings.Split and strings.ToLower
func FieldsToLowercase(s string, separator string) []string {
	return nil
}

// JoinedToUppercase makes a new slice the same length as the given list,
// converts all of the elements of the given slice into uppercase -- assigning
// the converted string into the new slice, then joins the elements of the new
// slice together into a single string, with the given separator
// between the elements.
//
// uses strings.ToUpper and strings.Join
func JoinedToUppercase(fields []string, separator string) string {
	return ""
}

// SpacesTrimmed makes a new slice the same length as the given list,
// trims the leading and trailing spaces of each element of the given slice --
// assigning the trimmed string into the new slice, then returns the new slice.
//
// uses strings.TrimSpace
func SpacesTrimmed(list []string) []string {
 	return nil
}

// SuffixesTrimmed makes a new slice the same length as the given list,
// trims the given suffix from each element of the given slice --
// assigning the trimmed string into the new slice, then returns the new slice.
// Case sensitive.
//
// uses strings.TrimSuffix
func SuffixesTrimmed(list []string, suffix string) []string {
	return nil
}

// You can use main to experiment as you implement your functions. Then,
// instead of "go test *.go" simply run the program.
// Do not forget that you still must pass the unit tests.
func main() {
}
