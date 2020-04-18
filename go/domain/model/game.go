package model

import (
	"time"

	"github.com/dddGameOnlineshop/go/domain/values"
)

// Game ゲームドメインモデル
type Game struct {
	ID          int
	Title       values.Title
	Detail      string
	Genre       string
	Price       values.Price
	ReleaseDate time.Time
	Maker       string
	Cero        values.Cero
}

// GemeInput Game生成用のインプット構造体
type GemeInput struct {
	Title       string
	Detail      string
	Genre       string
	Price       int
	ReleaseDate time.Time
	Maker       string
	Cero        string
}
