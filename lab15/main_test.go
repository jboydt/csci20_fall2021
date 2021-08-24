package main

import (
	"math"
	"testing"
)

const tolerance = 0.001

func makeMovie() *Movie {
	return NewMovie("Godzilla", "action")
}

func TestNewMovie(t *testing.T) {
	m := makeMovie()
	if m.Title != "Godzilla" {
		t.Errorf("Expected 'Godzilla', got %s", m.Title)
	}
	if m.Genre != "action" {
		t.Errorf("Expected 'action', got %s", m.Genre)
	}
	if m.Likes != 0 {
		t.Errorf("Expected 0 likes, got %d", m.Likes)
	}
	if m.Dislikes != 0 {
		t.Errorf("Expected 0 dislikes, got %d", m.Likes)
	}
	if m.Views != 0 {
		t.Errorf("Expected 0 views, got %d", m.Views)
	}
	if math.Abs(m.GetPercentageLikes() - 0.0) > tolerance {
		t.Errorf("Expected 0.0%% likes, got %.1f%%", m.GetPercentageLikes())
	}
	if math.Abs(m.GetPercentageDislikes() - 0.0) > tolerance {
		t.Errorf("Expected 0.0%% dislikes, got %.1f%%", m.GetPercentageDislikes())
	}
}

func TestLikeViews(t *testing.T) {
	m := makeMovie()
	m.AddLikeView()
	if m.Views != 1 {
		t.Errorf("Expected 1 view, got %d", m.Views)
	}
	if m.Likes != 1 {
		t.Errorf("Expected 1 like, got %d", m.Likes)
	}
	if math.Abs(m.GetPercentageLikes() - 100.0) > tolerance {
		t.Errorf("Expected 100.0%% likes, got %.1f%%", m.GetPercentageLikes())
	}
}

func TestDislikeViews(t *testing.T) {
	m := makeMovie()
	m.AddDislikeView()
	if m.Views != 1 {
		t.Errorf("Expected 1 view, got %d", m.Views)
	}
	if m.Dislikes != 1 {
		t.Errorf("Expected 1 dislike, got %d", m.Dislikes)
	}
	if math.Abs(m.GetPercentageDislikes() - 100.0) > tolerance {
		t.Errorf("Expected 100.0%% dislikes, got %.1f%%", m.GetPercentageDislikes())
	}
}

func TestLikesAndDislikes(t *testing.T) {
	m := makeMovie()
	m.AddLikeView()
	m.AddDislikeView()
	m.AddLikeView()
	m.AddDislikeView()
	m.AddLikeView()

	if m.Views != 5 {
		t.Errorf("Expected 5 views, got %d", m.Views)
	}
	if m.Likes != 3 {
		t.Errorf("Expected 3 likes, got %d", m.Likes)
	}
	if m.Dislikes != 2 {
		t.Errorf("Expected 2 dislikes, got %d", m.Dislikes)
	}
	if math.Abs(m.GetPercentageLikes() - 60.0) > tolerance {
		t.Errorf("Expected 60.0%% likes, got %.1f%%", m.GetPercentageLikes())
	}
	if math.Abs(m.GetPercentageDislikes() - 40.0) > tolerance {
		t.Errorf("Expected 40.0%% dislikes, got %.1f%%", m.GetPercentageDislikes())
	}
}

func TestToString(t *testing.T) {
	m := makeMovie()
	m.AddLikeView()
	m.AddLikeView()
	m.AddLikeView()
	m.AddDislikeView()

	mStr := m.ToString()
	expStr := "Godzilla, action, 4 views, 75.0% likes"
	if mStr != expStr {
		t.Errorf("Expected '%s', got '%s'", expStr, mStr)
	}
}
