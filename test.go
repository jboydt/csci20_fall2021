package main

import (
	"fmt"
)

func main() {
	weeklyTemperatures := make([]float64, 7)
	weeklyTemperatures[1] = 82.5
	weeklyTemperatures[5] = 98.7

	for i := 0; i < len(weeklyTemperatures); i++ {
		fmt.Printf("%.1f", weeklyTemperatures[i])
		if i < len(weeklyTemperatures) - 1 {
			fmt.Printf(" - ")
        	}
    	}
}

