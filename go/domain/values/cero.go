package values

import (
	"fmt"
)

// Cero 年齢別レーティング
type Cero struct {
	Rate             string
	Target           string
	TargetMinimumAge int64
}

// NewCero 年齢別レーティング生成メソッド
func NewCero(cero string) (*Cero, error) {

	switch cero {
	case "A":
		resp := Cero{Rate: cero, Target: "全年齢対象", TargetMinimumAge: 0}
		return &resp, nil
	case "B":
		resp := Cero{Rate: cero, Target: "１２才以上対象", TargetMinimumAge: 12}
		return &resp, nil
	case "C":
		resp := Cero{Rate: cero, Target: "１５才以上対象", TargetMinimumAge: 15}
		return &resp, nil
	case "D":
		resp := Cero{Rate: cero, Target: "１７才以上対象", TargetMinimumAge: 17}
		return &resp, nil
	case "Z":
		resp := Cero{Rate: cero, Target: "１８才以上対象", TargetMinimumAge: 18}
		return &resp, nil
	default:
		return nil, fmt.Errorf("不正な値 cero: %s", cero)
	}
}

// NewACero Aレーティングオブジェクトを生成
func NewACero() (*Cero, error) {
	cero, err := NewCero("A")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return cero, nil
}

// isARate レーティングAかチェック
func (c *Cero) isARate() bool {
	return c.Rate == "A"
}

// isBRate レーティングBかチェック
func (c *Cero) isBRate() bool {
	return c.Rate == "B"
}

// isCRate レーティングCかチェック
func (c *Cero) isCRate() bool {
	return c.Rate == "C"
}

// isDRate レーティングDかチェック
func (c *Cero) isDRate() bool {
	return c.Rate == "D"
}

// isZRate レーティングZかチェック
func (c *Cero) isZRate() bool {
	return c.Rate == "Z"
}

// validate 引数で与えられた年齢がレーティング対象年齢を満たしているかチェック
func (c *Cero) validate(age int64) bool {
	return c.TargetMinimumAge <= age
}
