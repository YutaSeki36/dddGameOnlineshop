package model

import (
	"fmt"
	"time"

	"github.com/dddGameOnlineshop/go/domain/values"
)

// Game ゲームドメインモデル
// フィールドの頭文字を小文字にして，不用意にGetされないようにする
// 必要な項目はGetterを用意する
type Game struct {
	id          int
	title       values.Title
	detail      string
	genre       string
	price       values.Price
	releaseDate time.Time
	maker       []int // マリオ&ソニックのように共同開発される場合があるため
	cero        values.Cero
}

// Games ゲームリスト
type Games []Game

// GameInput Game生成用のインプット構造体
type GameInput struct {
	Title       string
	Detail      string
	Genre       string
	Price       int
	ReleaseDate time.Time
	Maker       []int
	Cero        string
}

// NewGame Gameオブジェクトの生成メソッド
func NewGame(input *GameInput) (*Game, error) {
	title, err := values.NewTitle(input.Title)
	if err != nil {
		return nil, fmt.Errorf("タイトル生成に失敗しました．エラー内容: %s", err.Error())
	}

	price, err := values.NewPrice(input.Price)
	if err != nil {
		return nil, fmt.Errorf("価格生成に失敗しました．エラー内容: %s", err.Error())
	}

	cero, err := values.NewCero(input.Cero)
	if err != nil {
		return nil, fmt.Errorf("Ceroオブジェクト生成に失敗しました．エラー内容: %s", err.Error())
	}

	return &Game{
		title:       *title,
		detail:      input.Detail,
		genre:       input.Genre,
		price:       *price,
		releaseDate: input.ReleaseDate,
		maker:       input.Maker,
		cero:        *cero,
	}, nil
}

// IsAlreadyReleased 対象ゲームが発売済みかチェック
func (g *Game) IsAlreadyReleased() bool {
	now := time.Now()
	return now.After(g.releaseDate)
}
