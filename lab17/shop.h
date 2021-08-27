// Item and ShoppingList (C++)
//
// Programmer name: KEY
// Date completed:  26 August 2021

#pragma once

#include <iomanip>
#include <sstream>
#include <string>
#include <vector>
using namespace std;

// Item represents an item in a shopping list.
//
// fields: description(string), quantity(int), pricePerItem(float)
struct Item {

};

// ShoppingList represents a shopping list with items. A Label
// can be used to personalize the list or associate the list with
// a specific store or occasion.
//
// fields: label(string), allItems(vector<Item>)
class ShoppingList {
private:
  string label;
  vector<Item> allItems;

public:
  // Construct a shopping list with a given label.
  ShoppingList(string label) {

  }

  // addItem adds an item to a shopping list. Checks first to see if the item
  // is already present in the shopping list -- if the item is not present,
  // the item is appended to the list. If the item is present, then the
  // quantity of the item in the list is increased by the quantity of the
  // item passed as an argument.
  void addItem(Item i) {

  }

  // hasItem checks to see (by Description) if an item is already in the
  // shopping list. Returns true if the item is present, otherwise returns
  // false.
  bool hasItem(Item i) {
      return false;
  }

  // increaseItemQuantity locates (in the shopping list) an item matching the
  // given item's Description, and increases the quantity of the item in the
  // shopping list by the quantity in the given item.
  void increaseItemQuantity(Item i) {

  }

  // toString returns a string representation of a shopping list.
  //
  // TEMPLATE:  label:
  //            (1) ITEM_NAME, QUANTITY @ $PRICE_PER_ITEM each
  //            (2) ITEM_NAME, QUANTITY @ $PRICE_PER_ITEM each
  //
  // EXAMPLE:   Boyd's Grocery List
  //            (1) apples, 3 @ $0.75 each
  //            (2) oranges, 2 @ $1.50 each
  string toString() {
    stringstream output;
    output << this->label << ":\n";



    return output.str();
  }

  // totalCost returns the total cost of the items in the shopping list, by
  // calculating the sum of quantity * price_per_item of all items in the
  // shopping list.
  float totalCost() {
    return 0.0;
  }
};
