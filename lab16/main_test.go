package main

import (
	"math"
	"testing"
)

const tolerance = 0.001

func TestBasicItem(t *testing.T) {
	item := Item{"apples", 3, 0.75}

	if item.Description != "apples" {
		t.Errorf("Expected 'apples', got '%s'", item.Description)
	}

	if item.Quantity != 3 {
		t.Errorf("Expected 3, got %d", item.Quantity)
	}

	if math.Abs(item.PricePerItem - 0.75) > tolerance {
		t.Errorf("Expected 0.75, got %.2f", item.PricePerItem)
	}
}

func TestBasicShoppingList(t *testing.T) {
	list := ShoppingList{Label: "basic"}

	if list.Label != "basic" {
		t.Errorf("Expected 'basic', got '%s'", list.Label)
	}

	if len(list.AllItems) != 0 {
		t.Errorf("Expected length 0 list, got length %d", len(list.AllItems))
	}
}

func TestNewShoppingList(t *testing.T) {
	list := NewShoppingList("basic")

	if list.Label != "basic" {
		t.Errorf("Expected 'basic', got '%s'", list.Label)
	}

	if len(list.AllItems) != 0 {
		t.Errorf("Expected length 0 list, got length %d", len(list.AllItems))
	}

	if list.HasItem(Item{"apples", 3, 0.75}) == true {
		t.Errorf("Expected HasItem(apple) to return false, got true")
	}
}

func TestShoppingListNoDuplicateItems(t *testing.T) {
	list := NewShoppingList("basic")

	if list.Label != "basic" {
		t.Errorf("Expected 'basic', got '%s'", list.Label)
	}

	if len(list.AllItems) != 0 {
		t.Errorf("Expected length 0 list, got length %d", len(list.AllItems))
	}

	list.AddItem(Item{"apples", 3, 0.75})
	list.AddItem(Item{"oranges", 2, 1.50})

	if len(list.AllItems) != 2 {
		t.Errorf("Expected length 2 list, got length %d", len(list.AllItems))
	}

	if list.HasItem(Item{"apples", 3, 0.75}) == false {
		t.Errorf("Expected HasItem(apple) to return true, got false")
	}

	if list.HasItem(Item{"oranges", 2, 1.50}) == false {
		t.Errorf("Expected HasItem(oranges) to return true, got false")
	}

	if math.Abs(list.TotalCost() - 5.25) > tolerance {
		t.Errorf("Expected total cost to be 5.25, got %.2f", list.TotalCost())
	}
}

func TestShoppingListWithDuplicateItems(t *testing.T) {
	list := NewShoppingList("basic")

	list.AddItem(Item{"apples", 3, 0.75})
	list.AddItem(Item{"apples", 3, 0.75})
	list.AddItem(Item{"oranges", 2, 1.50})
	list.AddItem(Item{"oranges", 2, 1.50})

	if len(list.AllItems) != 2 {
		t.Errorf("Expected length 2 list, got length %d", len(list.AllItems))
	}

	if list.HasItem(Item{"apples", 3, 0.75}) == false {
		t.Errorf("Expected HasItem(apple) to return true, got false")
	}

	if list.HasItem(Item{"oranges", 2, 1.50}) == false {
		t.Errorf("Expected HasItem(oranges) to return true, got false")
	}

	if math.Abs(list.TotalCost() - 10.50) > tolerance {
		t.Errorf("Expected total cost to be 10.50, got %.2f", list.TotalCost())
	}
}

func TestShoppingListEmptyToString(t *testing.T) {
	list := NewShoppingList("test")

	expected := "test:\n"

	if list.ToString() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, list.ToString())
	}
}

func TestShoppingListNotEmptyToString(t *testing.T) {
	list := NewShoppingList("test")
	list.AddItem(Item{"apples", 3, 0.75})
	list.AddItem(Item{"apples", 3, 0.75})
	list.AddItem(Item{"oranges", 2, 1.50})
	list.AddItem(Item{"oranges", 2, 1.50})

	expected := "test:\n(1) apples, 6 @ $0.75 each\n(2) oranges, 4 @ $1.50 each\n"

	if list.ToString() != expected {
		t.Errorf("Expected '%s'\n\ngot '%s'", expected, list.ToString())
	}
}
