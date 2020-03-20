package model

import "time"

// 後でフィールドを値オブジェクトにする
type Game struct {
	id          int
	title       string
	genre       string
	releaseDate time.Time
	maker       string
}
