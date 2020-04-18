package values

import (
	"errors"
	"fmt"

	"github.com/dustin/go-humanize"
)

// Price ゲームの価格
type Price int64

// NewPrice 価格値オブジェクト生成メソッド
func NewPrice(price int) (*Price, error) {

	// キャンペーンなどで0円になる可能性があるので0円を許可
	if price < 0 || price > 9999999 {
		return nil, errors.New("不正な価格設定です")
	}
	resp := Price(price)
	return &resp, nil
}

// ToInt intに変換
func (p *Price) ToInt() int64 {
	return int64(*p)
}

// ToString 価格を文字列で取得
func (p *Price) ToString() string {
	return fmt.Sprintf("%s円", humanize.Comma(p.ToInt()))
}
