package main

import (
	"math"
  "math/rand"
	"testing"
)

func TestDistanceBetween(t *testing.T) {
	distance := DistanceBetween(0, 0, 5, 5)
	roundedDistance := math.Floor(distance*100) / 100
	if roundedDistance != 7.07 {
		t.Errorf("Expected 7.07, got %.2f\n", roundedDistance)
	}
}

func TestAreaOfCircle(t *testing.T) {
	area := AreaOfCircle(5.0)
	roundedArea := math.Floor(area*100) / 100
	if roundedArea != 78.53 {
		t.Errorf("Expected 78.53, got %.2f\n", roundedArea)
	}
}

func TestAreaOfRectangle(t *testing.T) {
	area := AreaOfRectangle(1.5, 2.5)
	roundedArea := math.Floor(area*100) / 100
	if roundedArea != 3.75 {
		t.Errorf("Expected 3.75, got %.2f\n", roundedArea)
	}
}

func TestVolumeOfCylinder(t *testing.T) {
	volume := VolumeOfCylinder(5.0, 7.5)
	roundedVolume := math.Floor(volume*100) / 100
	if roundedVolume != 589.04 {
		t.Errorf("Expected 589.04, got %.2f\n", roundedVolume)
	}
}

func TestVolumeOfBox(t *testing.T) {
	volume := VolumeOfBox(1.5, 2.5, 3.5)
	roundedVolume := math.Floor(volume*100) / 100
	if roundedVolume != 13.12 {
		t.Errorf("Expected 13.12, got %.2f\n", roundedVolume)
	}
}

func TestRollD6(t *testing.T) {
  rand.Seed(1) // this seed should produce the expected sequence of numbers
  expectedRolls := [10]int{6,4,6,6,2,1,2,3,5,1}
  for i := 0; i < 10; i++ {
    dieRoll := RollD6()
    if dieRoll != expectedRolls[i] {
      t.Errorf("Failed on roll %d, expected %d, got %d",
        i, expectedRolls[i], dieRoll)
    }
  }
}

func TestRollD20(t *testing.T) {
  rand.Seed(10) // this seed should produce the expected sequence of numbers
  expectedRolls := [10]int{15,9,8,20,9,5,20,16,16,19}
  for i := 0; i < 10; i++ {
    dieRoll := RollD20()
    if dieRoll != expectedRolls[i] {
      t.Errorf("Failed on roll %d, expected %d, got %d",
        i, expectedRolls[i], dieRoll)
    }
  }
}
