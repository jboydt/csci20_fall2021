package main

import (
	"testing"
)

func TestGetEvens(t *testing.T) {
	var evens, expected string

	evens = GetEvens(0, 0)
	expected = "0"
	if evens != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, evens)
	}

	evens = GetEvens(1, 1)
	expected = ""
	if evens != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, evens)
	}

	evens = GetEvens(1, 2)
	expected = "2"
	if evens != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, evens)
	}

	evens = GetEvens(2, 3)
	expected = "2"
	if evens != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, evens)
	}

	evens = GetEvens(2, 10)
	expected = "2,4,6,8,10"
	if evens != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, evens)
	}

	evens = GetEvens(1, 11)
	expected = "2,4,6,8,10"
	if evens != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, evens)
	}
}

func TestGetOdds(t *testing.T) {
	var odds, expected string

	odds = GetOdds(0, 0)
	expected = ""
	if odds != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, odds)
	}

	odds = GetOdds(1, 1)
	expected = "1"
	if odds != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, odds)
	}

	odds = GetOdds(1, 2)
	expected = "1"
	if odds != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, odds)
	}

	odds = GetOdds(2, 3)
	expected = "3"
	if odds != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, odds)
	}

	odds = GetOdds(2, 10)
	expected = "3,5,7,9"
	if odds != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, odds)
	}

	odds = GetOdds(1, 11)
	expected = "1,3,5,7,9,11"
	if odds != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, odds)
	}
}

func TestFizzBuzz(t *testing.T) {
	var fizzbuzz, expected string

	fizzbuzz = FizzBuzz(1)
	expected = "1"
	if fizzbuzz != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, fizzbuzz)
	}

	fizzbuzz = FizzBuzz(5)
	expected = "1,2,Fizz,4,Buzz"
	if fizzbuzz != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, fizzbuzz)
	}

	fizzbuzz = FizzBuzz(15)
	expected = "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz"
	if fizzbuzz != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, fizzbuzz)
	}

	fizzbuzz = FizzBuzz(31)
	expected = "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz,16,17,Fizz,19,Buzz,Fizz,22,23,Fizz,Buzz,26,Fizz,28,29,FizzBuzz,31"
	if fizzbuzz != expected {
		t.Errorf("Expected '%s', got '%s'\n", expected, fizzbuzz)
	}
}
