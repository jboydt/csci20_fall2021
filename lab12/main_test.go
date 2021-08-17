package main

import (
	"fmt"
	"testing"
)

// LowestThree tests

func TestLowestThree(t *testing.T) {
	numbers := [10]int{5, 19, 84, 51, 21, 49, 40, 38, 10, 96}

	result := LowestThree(numbers[:])
	expected := [3]int{5, 10, 19}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestLowestThreeEmptySlice(t *testing.T) {
	var numbers []int

	result := LowestThree(numbers)
	var expected []int

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestLowestThreeOnlyOne(t *testing.T) {
	numbers := [1]int{10}

	result := LowestThree(numbers[:])
	expected := [1]int{10}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestLowestThreeOnlyTwo(t *testing.T) {
	numbers := [2]int{20, 10}

	result := LowestThree(numbers[:])
	expected := [2]int{10, 20}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestLowestThreeOnlyThree(t *testing.T) {
	numbers := [3]int{20, 10, 30}

	result := LowestThree(numbers[:])
	expected := [3]int{10, 20, 30}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// HighestThree tests

func TestHighestThree(t *testing.T) {
	numbers := [10]int{5, 19, 84, 51, 21, 49, 40, 38, 10, 96}

	result := HighestThree(numbers[:])
	expected := [3]int{51, 84, 96}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestHighestThreeEmptySlice(t *testing.T) {
	var numbers []int

	result := HighestThree(numbers)
	var expected []int

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestHighestThreeOnlyOne(t *testing.T) {
	numbers := [1]int{10}

	result := HighestThree(numbers[:])
	expected := [1]int{10}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestHighestThreeOnlyTwo(t *testing.T) {
	numbers := [2]int{10, 20}

	result := HighestThree(numbers[:])
	expected := [2]int{10, 20}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestHighestThreeOnlyThree(t *testing.T) {
	numbers := [3]int{10, 20, 30}

	result := HighestThree(numbers[:])
	expected := [3]int{10, 20, 30}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// DropHighestAndLowest tests

func TestDropHighestAndLowest(t *testing.T) {
	values := [10]float64{32.8, 42.1, 18.7, 30.1, 92.9, 88.4, 40.4, 63.7, 27.7, 10.3}

	result := DropHighestAndLowest(values[:])
	expected := [8]float64{18.7, 27.7, 30.1, 32.8, 40.4, 42.1, 63.7, 88.4}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestDropHighestAndLowestEmptySlice(t *testing.T) {
	var values []float64

	result := DropHighestAndLowest(values)
	var expected []float64

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestDropHighestAndLowestOnlyOne(t *testing.T) {
	values := [1]float64{1.1}

	result := DropHighestAndLowest(values[:])
	expected := [1]float64{1.1}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestDropHighestAndLowestOnlyTwo(t *testing.T) {
	values := [2]float64{2.2, 1.1}

	result := DropHighestAndLowest(values[:])
	expected := [2]float64{1.1, 2.2}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestDropHighestAndLowestOnlyThree(t *testing.T) {
	values := [2]float64{2.2, 1.1}

	result := DropHighestAndLowest(values[:])
	expected := [2]float64{1.1, 2.2}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// FirstHalf tests

func TestFirstHalf(t *testing.T) {
	colors := [5]string{"red", "green", "blue", "yellow", "purple"}

	result := FirstHalf(colors[:])
	expected := [2]string{"red", "green"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestFirstHalfEmptySlice(t *testing.T) {
	var colors []string

	result := FirstHalf(colors)
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestFirstHalfOnlyOne(t *testing.T) {
	colors := [1]string{"red"}

	result := FirstHalf(colors[:])
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestFirstHalfOnlyTwo(t *testing.T) {
	colors := [2]string{"red", "green"}

	result := FirstHalf(colors[:])
	expected := [1]string{"red"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// SecondHalf tests

func TestSecondHalf(t *testing.T) {
	colors := [5]string{"red", "green", "blue", "yellow", "purple"}

	result := SecondHalf(colors[:])
	expected := [3]string{"blue", "yellow", "purple"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestSecondHalfEmptySlice(t *testing.T) {
	var colors []string

	result := SecondHalf(colors)
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestSecondHalfOnlyOne(t *testing.T) {
	colors := [1]string{"red"}

	result := SecondHalf(colors[:])
	expected := [1]string{"red"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestSecondHalfOnlyTwo(t *testing.T) {
	colors := [2]string{"red", "green"}

	result := SecondHalf(colors[:])
	expected := [1]string{"green"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

// StringInCenter tests

func TestStringInCenter(t *testing.T) {
	colors := [5]string{"red", "green", "blue", "yellow", "purple"}

	result := StringInCenter(colors[:])
	expected := [1]string{"blue"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringInCenterEmptySlice(t *testing.T) {
	var colors []string

	result := StringInCenter(colors)
	var expected []string

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringInCenterOnlyOne(t *testing.T) {
	colors := [1]string{"red"}

	result := StringInCenter(colors[:])
	expected := [1]string{"red"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}

func TestStringInCenterOnlyTwo(t *testing.T) {
	colors := [2]string{"red", "green"}

	result := StringInCenter(colors[:])
	expected := [1]string{"green"}

	resultString := fmt.Sprintf("%v", result)
	expectedString := fmt.Sprintf("%v", expected)

	if resultString != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, resultString)
	}
}
