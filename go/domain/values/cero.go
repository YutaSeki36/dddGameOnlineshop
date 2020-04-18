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

// TODO: 購入者の年齢がCeroを満たしているかのチェック
