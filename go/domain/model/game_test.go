package model

import (
	"testing"
	"time"
)

func TestNewGame(t *testing.T) {
	testCase := []*GameInput{
		{
			Title:       "スーパーメリオ",
			Detail:      "冒険RPG",
			Genre:       "アクション",
			Price:       5600,
			ReleaseDate: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
			Maker:       []int{1},
			Cero:        "A",
		},
	}
	for _, c := range testCase {
		_, err := NewGame(c)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
	}
}

func TestIsAlreadyReleased(t *testing.T) {
	testCase := []*GameInput{
		{
			Title:       "スーパーメリオ",
			Detail:      "冒険RPG",
			Genre:       "アクション",
			Price:       5600,
			ReleaseDate: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
			Maker:       []int{1},
			Cero:        "A",
		},
	}
	for _, c := range testCase {
		game, _ := NewGame(c)
		result := game.IsAlreadyReleased()
		if !result {
			t.Fatalf("failed test %#v", result)
		}
	}
}
