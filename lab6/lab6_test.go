package main

import (
	"math"
	"testing"
)

func TestHelloWorld(t *testing.T) {
  greeting := HelloWorld()
  expected := "Hello world!"
  if greeting != expected {
    t.Errorf("Expected message to be '%s', got '%s'\n", expected, greeting)
  }
}

func TestGreetUserNoName(t *testing.T) {
  greeting := GreetUser("")
  expected := "Hello !"
  if greeting != expected {
    t.Errorf("Expected greeting to be '%s', got '%s'\n", expected, greeting)
  }
}

func TestGreetUserWithName(t *testing.T) {
  greeting := GreetUser("Gophers")
  expected := "Hello Gophers!"
  if greeting != expected {
    t.Errorf("Expected greeting to be '%s', got '%s'\n", expected, greeting)
  }
}

func TestEstimateAgeZero(t *testing.T) {
  age := EstimateAge(2021)
  expected := 0
  if age != expected {
    t.Errorf("Expected age to be %d, got %d\n", expected, age)
  }
}

func TestComputeAverageZeroPercent(t *testing.T) {
	average := ComputeAverage(0.0, 1.0)
	averageRounded := math.Floor(average*100) / 100
  expected := 0.0
	if averageRounded != expected {
		t.Errorf("Expected average to be %.1f, got %.1f\n", expected, averageRounded)
	}
}

func TestComputeAverageOneHundredPercent(t *testing.T) {
  average := ComputeAverage(1.0, 1.0)
	averageRounded := math.Floor(average*100) / 100
  expected := 100.0
	if averageRounded != expected {
		t.Errorf("Expected average to be %.1f, got %.1f\n", expected, averageRounded)
	}
}

func TestComputeAverageFiftyPercent(t *testing.T) {
  average := ComputeAverage(1.0, 2.0)
	averageRounded := math.Floor(average*100) / 100
  expected := 50.0
	if averageRounded != expected {
		t.Errorf("Expected average to be %.1f, got %.1f\n", expected, averageRounded)
	}
}
