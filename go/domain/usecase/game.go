package game

import (
	"time"

	"github.com/dddGameOnlineshop/go/domain/model"
)

// Usecase ユースケース構造体
type Usecase struct {
}

// NewGameUsecase ユースケース生成メソッド
func NewGameUsecase() *Usecase {
	return &Usecase{}
}

// CreateGemaInput CreateGemaのインプット
type CreateGemaInput struct {
	Title       string
	Detail      string
	Genre       string
	Price       int
	ReleaseDate time.Time
	Maker       string
	Cero        string
}

// CreateGeme ゲーム情報作成メソッド
func (u *Usecase) CreateGeme(input CreateGemaInput) (*model.Game, error) {
	return nil, nil
}

// FindAllGame ゲーム情報取得メソッド
func (u *Usecase) FindAllGame() (*model.Game, error) {
	return nil, nil
}
