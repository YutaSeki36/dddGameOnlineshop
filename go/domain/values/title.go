package values

import (
	"errors"
	"unicode/utf8"
)

// Title ゲームタイトルの値オブジェクト
type Title string

// NewTitle ゲームタイトル生成メソッド
func NewTitle(title string) (*Title, error) {

	if utf8.RuneCountInString(title) < 0 || utf8.RuneCountInString(title) >= 50 {
		return nil, errors.New("文字数は0文字以上50文字以下で設定してください．")
	}
	resp := Title(title)
	return &resp, nil
}

// ToString 文字列へ変換
func (t *Title) ToString() string {
	return string(*t)
}
