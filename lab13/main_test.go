package main

import (
	"fmt"
	"testing"
)

// StringsOfLength tests

func TestStringsOfLengthEmptyList(t *testing.T) {
	var list []string

	result := StringsOfLength(list, 5)
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsOfLengthOneMatch(t *testing.T) {
	list := [5]string{"pug", "slug", "sheep", "turtle", "dolphin"}

	result := StringsOfLength(list[:], 5)
	expected := [1]string{"sheep"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsOfLengthAllMatch(t *testing.T) {
	list := [4]string{"cat", "cow", "pig", "pug"}

	result := StringsOfLength(list[:], 3)
	expected := [4]string{"cat", "cow", "pig", "pug"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsOfLengthNoMatch(t *testing.T) {
	list := [5]string{"horse", "slug", "sheep", "turtle", "dolphin"}

	result := StringsOfLength(list[:], 3)
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// StringsBeginningWith tests

func TestStringsBeginningWithEmptyList(t *testing.T) {
	var list []string

	result := StringsBeginningWith(list, "abc")
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsBeginningWithOneMatch(t *testing.T) {
	list := [5]string{"pug", "slug", "abcsheep", "turtle", "dolphin"}

	result := StringsBeginningWith(list[:], "abc")
	expected := [1]string{"abcsheep"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsBeginningWithAllMatch(t *testing.T) {
	list := [5]string{"abcpug", "abcslug", "abcsheep", "abcturtle", "abcdolphin"}

	result := StringsBeginningWith(list[:], "abc")
	expected := [5]string{"abcpug", "abcslug", "abcsheep", "abcturtle", "abcdolphin"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsBeginningWithNoMatch(t *testing.T) {
	list := [5]string{"abcpug", "abcslug", "abcsheep", "abcturtle", "abcdolphin"}

	result := StringsBeginningWith(list[:], "xyz")
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// StringsEndingWith tests

func TestStringsEndingWithEmptyList(t *testing.T) {
	var list []string

	result := StringsEndingWith(list, "xyz")
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsEndingWithOneMatch(t *testing.T) {
	list := [5]string{"pug", "slug", "sheepxyz", "turtle", "dolphin"}

	result := StringsEndingWith(list[:], "xyz")
	expected := [1]string{"sheepxyz"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsEndingWithAllMatch(t *testing.T) {
	list := [5]string{"pugxyz", "slugxyz", "sheepxyz", "turtlexyz", "dolphinxyz"}

	result := StringsEndingWith(list[:], "xyz")
	expected := [5]string{"pugxyz", "slugxyz", "sheepxyz", "turtlexyz", "dolphinxyz"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsEndingWithNoMatch(t *testing.T) {
	list := [5]string{"pugxyz", "slugxyz", "sheepxyz", "turtlexyz", "dolphinxyz"}

	result := StringsEndingWith(list[:], "abc")
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// StringsContaining tests

func TestStringsContainingEmptyList(t *testing.T) {
	var list []string

	result := StringsContaining(list, "zzz")
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsContainingOneMatch(t *testing.T) {
	list := [5]string{"pug", "slug", "shezzzep", "turtle", "dolphin"}

	result := StringsContaining(list[:], "zzz")
	expected := [1]string{"shezzzep"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsContainingAllMatch(t *testing.T) {
	list := [5]string{"puzzzg", "sluzzzg", "shezzzep", "turzzztle", "dolzzzphin"}

	result := StringsContaining(list[:], "zzz")
	expected := [5]string{"puzzzg", "sluzzzg", "shezzzep", "turzzztle", "dolzzzphin"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringsContainingNoMatch(t *testing.T) {
	list := [5]string{"puzg", "sluzg", "shezep", "turztle", "dolzphin"}

	result := StringsContaining(list[:], "zzz")
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// FieldsToLowercase tests

func TestFieldsToLowercaseEmptyString(t *testing.T) {
	result := FieldsToLowercase("", "")
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestFieldsToLowercaseOneFieldLowercase(t *testing.T) {
	result := FieldsToLowercase("banana", "-")
	expected := [1]string{"banana"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestFieldsToLowercaseOneFieldUppercase(t *testing.T) {
	result := FieldsToLowercase("BANANA", "-")
	expected := [1]string{"banana"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestFieldsToLowercaseMultipleFieldsLowercase(t *testing.T) {
	result := FieldsToLowercase("vanilla-milk-shake", "-")
	expected := [3]string{"vanilla", "milk", "shake"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestFieldsToLowercaseMultipleFieldsUppercase(t *testing.T) {
	result := FieldsToLowercase("VANILLA-MILK-SHAKE", "-")
	expected := [3]string{"vanilla", "milk", "shake"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestFieldsToLowercaseMultipleFieldsMixedCase(t *testing.T) {
	result := FieldsToLowercase("VaNiLlA-mIlK-sHaKe", "-")
	expected := [3]string{"vanilla", "milk", "shake"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// JoinedToUppercase tests

func TestJoinedToUppercaseEmptyList(t *testing.T) {
	result := JoinedToUppercase([]string{}, "")
	expected := ""

	if result != expected {
		t.Errorf("Expected %s, got %s\n", expected, result)
	}
}

func TestJoinedToUppercaseOneFieldLowercase(t *testing.T) {
	list := [1]string{"banana"}
	result := JoinedToUppercase(list[:], "-")
	expected := "BANANA"

	if result != expected {
		t.Errorf("Expected %s, got %s\n", expected, result)
	}
}

func TestJoinedToUppercaseOneFieldUppercase(t *testing.T) {
	list := [1]string{"BANANA"}
	result := JoinedToUppercase(list[:], "-")
	expected := "BANANA"

	if result != expected {
		t.Errorf("Expected %s, got %s\n", expected, result)
	}
}

func TestJoinedToUppercaseMultipleFieldsLowercase(t *testing.T) {
	list := [3]string{"vanilla", "milk", "shake"}
	result := JoinedToUppercase(list[:], "-")
	expected := "VANILLA-MILK-SHAKE"

	if result != expected {
		t.Errorf("Expected %s, got %s\n", expected, result)
	}
}

func TestJoinedToUppercaseMultipleFieldsUppercase(t *testing.T) {
	list := [3]string{"VANILLA", "MILK", "SHAKE"}
	result := JoinedToUppercase(list[:], "-")
	expected := "VANILLA-MILK-SHAKE"

	if result != expected {
		t.Errorf("Expected %s, got %s\n", expected, result)
	}
}

func TestJoinedToUppercaseMultipleFieldsMixedCase(t *testing.T) {
	list := [3]string{"vAnIlLa", "MiLk", "ShAkE"}
	result := JoinedToUppercase(list[:], "-")
	expected := "VANILLA-MILK-SHAKE"

	if result != expected {
		t.Errorf("Expected %s, got %s\n", expected, result)
	}
}

// SpacesTrimmed tests

func TestSpacesTrimmedEmptyList(t *testing.T) {
	var list []string
	result := SpacesTrimmed(list)
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestSpacesTrimmedNoTrimmingNeeded(t *testing.T) {
	list := [3]string{"red", "green", "blue"}
	result := SpacesTrimmed(list[:])
	expected := [3]string{"red", "green", "blue"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestSpacesTrimmedOnlyLeadingSpaces(t *testing.T) {
	list := [3]string{" \t\r\n red", " \t\r\n green", " \t\r\n blue"}
	result := SpacesTrimmed(list[:])
	expected := [3]string{"red", "green", "blue"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestSpacesTrimmedOnlyTrailingSpaces(t *testing.T) {
	list := [3]string{"red \t\r\n ", "green \t\r\n ", "blue \t\r\n "}
	result := SpacesTrimmed(list[:])
	expected := [3]string{"red", "green", "blue"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestSpacesTrimmedLeadingAndTrailingSpaces(t *testing.T) {
	list := [3]string{" \t\r\n red \t\r\n ", " \t\r\n green \t\r\n ", " \t\r\n blue \t\r\n "}
	result := SpacesTrimmed(list[:])
	expected := [3]string{"red", "green", "blue"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// SuffixesTrimmed tests

func TestSuffixesTrimmedEmptyList(t *testing.T) {
	var list []string
	result := SuffixesTrimmed(list, ".txt")
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestSuffixesTrimmedNoTrimmingNeeded(t *testing.T) {
	list := [3]string{"file1", "file2", "file3"}
	result := SuffixesTrimmed(list[:], ".txt")
	expected := [3]string{"file1", "file2", "file3"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestSuffixesTrimmed(t *testing.T) {
	list := [3]string{"file1.txt", "file2.txt", "file3.txt"}
	result := SuffixesTrimmed(list[:], ".txt")
	expected := [3]string{"file1", "file2", "file3"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}
