package model

import (
	"time"

	"github.com/dddGameOnlineshop/go/domain/values"
)

// Game ゲームドメインモデル
type Game struct {
	id          int
	title       values.Title
	detail      string
	genre       string
	price       values.Price
	releaseDate time.Time
	maker       string
}

// GemeInput Game生成用のインプット構造体
type GemeInput struct {
	title       string
	detail      string
	genre       string
	price       int
	releaseDate time.Time
	maker       string
}
