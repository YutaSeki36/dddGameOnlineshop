package model

import (
	"time"

	"github.com/dddGameOnlineshop/go/domain/values"
)

// 後でフィールドを値オブジェクトにする
type Game struct {
	id          int
	title       values.Title
	detail      string
	genre       string
	price       int
	releaseDate time.Time
	maker       string
}
