package values

import (
	"fmt"
	"regexp"
)

// URL URLモデル
type URL string

// NewURL URL生成メソッド
func NewURL(url string) (*URL, error) {
	reg := regexp.MustCompile(`https?://[\w/:%#\$&\?\(\)~\.=\+\-]+`)
	if !reg.MatchString(url) {
		return nil, fmt.Errorf("不正な値 cero: %s", url)
	}

	resp := URL(url)
	return &resp, nil
}

// ToString 文字列へ変換
func (u *URL) ToString() string {
	return string(*u)
}
