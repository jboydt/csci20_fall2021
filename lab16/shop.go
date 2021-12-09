// Item and ShoppingList
//
// Programmer name: __________
// Date completed:  __________

package main

import "fmt"

// Item represents an item in a shopping list.
//
// fields: Description(string), Quantity(int), PricePerItem(float64)
type Item struct {

}

// ShoppingList represents a shopping list with items. A Label
// can be used to personalize the list or associate the list with
// a specific store or occasion.
//
// fields: Label(string), AllItems([]Item)
type ShoppingList struct {

}

// NewShoppingList returns a pointer to a new shopping list with
// the given label and no items.
func NewShoppingList(label string) *ShoppingList {
  return nil
}

// AddItem adds an item to a shopping list. Checks first to see if the item
// is already present in the shopping list -- if the item is not present,
// the item is appended to the list. If the item is present, then the
// quantity of the item in the list is increased by the quantity of the
// item passed as an argument.
func (s *ShoppingList) AddItem(i Item) {

}

// HasItem checks to see (by Description) if an item is already in the
// shopping list. Returns true if the item is present, otherwise returns
// false.
func (s *ShoppingList) HasItem(i Item) bool {
  return false
}

// IncreaseItemQuantity locates (in the shopping list) an item matching the
// given item's Description, and increases the quantity of the item in the
// shopping list by the quantity in the given item.
func (s *ShoppingList) IncreaseItemQuantity(i Item) {

}

// ToString returns a string representation of a shopping list.
//
// TEMPLATE:  label:
//            (1) ITEM_NAME, QUANTITY @ $PRICE_PER_ITEM each
//            (2) ITEM_NAME, QUANTITY @ $PRICE_PER_ITEM each
//
// EXAMPLE:   Boyd's Grocery List:
//            (1) apples, 3 @ $0.75 each
//            (2) oranges, 2 @ $1.50 each
func (s *ShoppingList) ToString() string {
  return ""
}

// TotalCost returns the total cost of the items in the shopping list, by
// calculating the sum of quantity * price_per_item of all items in the
// shopping list.
func (s *ShoppingList) TotalCost() float64 {
  return 0.0
}
