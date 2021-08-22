package main

import "fmt"

type Treasure struct {
	description string
	value int
}

func main() {
	//var t1 Treasure
	//var t2 = Treasure{"Gold Coins", 100}
	var t3 = Treasure{value: 500, description: "Rubies"}
	//t4 := Treasure{description: "Wool Scarf"}
	//t5 := new(Treasure)

	fmt.Printf("%v\n", t3)

	fmt.Printf("Done\n\n")
}
