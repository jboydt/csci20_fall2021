// Lab 15 -- pointers and methods
//
// Programmer name: __________
// Date completed:  __________

package main

import "fmt"

// Movie represents a movie for a database of movies.
type Movie struct {
	Title    string
	Genre    string
	Likes    int
	Dislikes int
	Views    int
}

// NewMovie returns a Movie with the given title and genre.
// Views, Likes, and Dislikes all left in their zero-value states.
func NewMovie(title, genre string) *Movie {
	return nil
}

// AddLikeView increments Views and Likes.
func (m *Movie) AddLikeView() {

}

// AddDislikeView increments Views and Dislikes.
func (m *Movie) AddDislikeView() {

}

// GetPercentageLikes returns the percentage of likes against
// the Views for the Movie. The value is between 0.0 and 100.0.
func (m *Movie) GetPercentageLikes() float64 {
	return 0.0
}

// GetPercentageDislikes returns the percentage of dislikes against
// the Views for the Movie. The value is between 0.0 and 100.0.
func (m *Movie) GetPercentageDislikes() float64 {
	return 0.0
}

// ToString returns a string representation of the Movie.
//
// TEMPLATE: TITLE, GENRE, X views, X.X% likes
// EXAMPLE:  Godzilla, action, 100 views, 80.0% likes
func (m *Movie) ToString() string {
	return ""
}

// You can use main to experiment as you implement your methods. Then,
// instead of "go test *.go" simply run the program.
// Do not forget that you still must pass the unit tests.
func main() {
}
