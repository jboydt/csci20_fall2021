package main

import (
	"math"
	"testing"
)

const floatTolerance = 0.001

func TestCartItemStruct(t *testing.T) {
	item := CartItem{name: "lemon", price: 0.25, quantity: 2}

	if item.name != "lemon" {
		t.Errorf("Expected name to be 'lemon', got '%s'", item.name)
	}

	if math.Abs(item.price - 0.25) > floatTolerance {
		t.Errorf("Expected price to be 0.25, got %.2f", item.price)
	}

	if item.quantity != 2 {
		t.Errorf("Expected quantity to be 2, got %d", item.quantity)
	}
}

func TestCartItemString(t *testing.T) {
	item := CartItem{name: "lemon", price: 0.25, quantity: 2}

	itemString := CartItemString(item)
	expectedString := "lemon, 2 @ $0.25"

	if itemString != expectedString {
		t.Errorf("Expected '%s', got '%s'", expectedString, itemString)
	}
}

func TestCartItemCost(t *testing.T) {
	item := CartItem{name: "lemon", price: 0.25, quantity: 2}

	total := CartItemCost(item)
	expectedTotal := 0.50

	if math.Abs(expectedTotal - total) > floatTolerance {
		t.Errorf("Expected %.2f, got %.2f", expectedTotal, total)
	}
}
