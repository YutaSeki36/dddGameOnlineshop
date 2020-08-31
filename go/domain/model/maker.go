package model

import (
	"github.com/dddGameOnlineshop/go/domain/values"
)

// Maker ゲームメーカモデル
type Maker struct {
	id           int
	name         string
	publishGames []int
	url          values.URL // ゲームメーカーの公式ホームページ
}

// MakerInput ゲームメーカモデルインプット
type MakerInput struct {
	Name         string
	PublishGames []int
	URL          string
}
