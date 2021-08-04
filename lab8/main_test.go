package main

import (
	"math"
	"testing"
)

func TestEncodeProduct(t *testing.T) {
	code := EncodeProduct("apples")
	if code != 1 {
		t.Errorf("Expected code 1 for apples, got %d\n", code)
	}

	code = EncodeProduct("bananas")
	if code != 2 {
		t.Errorf("Expected code 2 for bananas, got %d\n", code)
	}

	code = EncodeProduct("grapes")
	if code != 3 {
		t.Errorf("Expected code 3 for grapes, got %d\n", code)
	}

	code = EncodeProduct("oranges")
	if code != 4 {
		t.Errorf("Expected code 4 for oranges, got %d\n", code)
	}

	code = EncodeProduct("strawberries")
	if code != 5 {
		t.Errorf("Expected code 5 for strawberries, got %d\n", code)
	}

	code = EncodeProduct("okra")
	if code != -1 {
		t.Errorf("Expected code -1 for okra, got %d\n", code)
	}
}

func TestDecodeProduct(t *testing.T) {
	productName := DecodeProduct(1)
	if productName != "apples" {
		t.Errorf("Expected apples for code 1, got %s\n", productName)
	}

	productName = DecodeProduct(2)
	if productName != "bananas" {
		t.Errorf("Expected bananas for code 2, got %s\n", productName)
	}

	productName = DecodeProduct(3)
	if productName != "grapes" {
		t.Errorf("Expected grapes for code 3, got %s\n", productName)
	}

	productName = DecodeProduct(4)
	if productName != "oranges" {
		t.Errorf("Expected oranges for code 4, got %s\n", productName)
	}

	productName = DecodeProduct(5)
	if productName != "strawberries" {
		t.Errorf("Expected strawberries for code 5, got %s\n", productName)
	}

	productName = DecodeProduct(99)
	if productName != "invalid" {
		t.Errorf("Expected invalid for code 99, got %s\n", productName)
	}
}

func TestComputeShippingCharge(t *testing.T) {
	charge := ComputeShippingCharge(1.0, 2.0)
	roundedCharge := math.Floor(charge*100) / 100
	if roundedCharge != 0.0 {
		t.Errorf("Expected 0.0, got %.1f\n", roundedCharge)
	}

	charge = ComputeShippingCharge(1.0, 1.0)
	roundedCharge = math.Floor(charge*100) / 100
	if roundedCharge != 5.0 {
		t.Errorf("Expected 5.0, got %.1f\n", roundedCharge)
	}

	charge = ComputeShippingCharge(2.0, 1.0)
	roundedCharge = math.Floor(charge*100) / 100
	if roundedCharge != 5.0 {
		t.Errorf("Expected 5.0, got %.1f\n", roundedCharge)
	}
}

func TestTakeWalk(t *testing.T) {
	walk := TakeWalk("sunny", "high", "")
	if walk != false {
		t.Errorf("Expected sunny/high to be false, got true\n")
	}

	walk = TakeWalk("sunny", "normal", "")
	if walk != true {
		t.Errorf("Expected sunny/normal to be true, got false\n")
	}

	walk = TakeWalk("overcast", "", "")
	if walk != true {
		t.Errorf("Expected overcast to be true, got false\n")
	}

	walk = TakeWalk("rainy", "", "strong")
	if walk != false {
		t.Errorf("Expected rainy/strong to be false, got true\n")
	}

	walk = TakeWalk("rainy", "", "weak")
	if walk != true {
		t.Errorf("Expected rainy/weak to be true, got false\n")
	}

	walk = TakeWalk("sunny", "", "")
	if walk != false {
		t.Errorf("Expected sunny (incomplete) to be false, got true\n")
	}

	walk = TakeWalk("rainy", "", "")
	if walk != false {
		t.Errorf("Expected rainy (incomplete) to be false, got true\n")
	}

	walk = TakeWalk("unknown", "unknown", "unknown")
	if walk != false {
		t.Errorf("Expected unknown/unknown/unknown to be false, got true\n")
	}
}

func TestPrettyPlayingCard(t *testing.T) {
	suit := "c"
	value := 1
	card := PrettyPlayingCard(value, suit)
	expected := "Ace of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 2
	card = PrettyPlayingCard(value, suit)
	expected = "2 of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 3
	card = PrettyPlayingCard(value, suit)
	expected = "3 of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 4
	card = PrettyPlayingCard(value, suit)
	expected = "4 of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 5
	card = PrettyPlayingCard(value, suit)
	expected = "5 of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 6
	card = PrettyPlayingCard(value, suit)
	expected = "6 of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 7
	card = PrettyPlayingCard(value, suit)
	expected = "7 of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 8
	card = PrettyPlayingCard(value, suit)
	expected = "8 of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 9
	card = PrettyPlayingCard(value, suit)
	expected = "9 of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 10
	card = PrettyPlayingCard(value, suit)
	expected = "10 of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 11
	card = PrettyPlayingCard(value, suit)
	expected = "Jack of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 12
	card = PrettyPlayingCard(value, suit)
	expected = "Queen of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	value = 13
	card = PrettyPlayingCard(value, suit)
	expected = "King of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	suit = "C"
	value = 1
	card = PrettyPlayingCard(value, suit)
	expected = "Ace of Clubs"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	suit = "d"
	value = 10
	card = PrettyPlayingCard(value, suit)
	expected = "10 of Diamonds"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	suit = "D"
	value = 11
	card = PrettyPlayingCard(value, suit)
	expected = "Jack of Diamonds"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	suit = "h"
	value = 9
	card = PrettyPlayingCard(value, suit)
	expected = "9 of Hearts"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	suit = "H"
	value = 12
	card = PrettyPlayingCard(value, suit)
	expected = "Queen of Hearts"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	suit = "s"
	value = 8
	card = PrettyPlayingCard(value, suit)
	expected = "8 of Spades"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	suit = "S"
	value = 13
	card = PrettyPlayingCard(value, suit)
	expected = "King of Spades"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}

	suit = ""
	value = 0
	card = PrettyPlayingCard(value, suit)
	expected = "ERROR of ERROR"
	if card != expected {
		t.Errorf("Expected %d/%s to be '%s', got %s\n", value, suit, expected, card)
	}
}
