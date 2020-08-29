package values

import (
	"fmt"
	"testing"
)

var testCases = []*Cero{
	{Rate: "A", Target: "全年齢対象", TargetMinimumAge: 0},
	{Rate: "B", Target: "１２才以上対象", TargetMinimumAge: 12},
	{Rate: "C", Target: "１５才以上対象", TargetMinimumAge: 15},
	{Rate: "D", Target: "１７才以上対象", TargetMinimumAge: 17},
	{Rate: "Z", Target: "１８才以上対象", TargetMinimumAge: 18},
}

func TestNewCeroSuccess(t *testing.T) {
	testCase := []string{
		"A",
		"B",
		"C",
		"D",
		"Z",
	}
	for _, c := range testCase {
		_, err := NewCero(c)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
	}
}

func TestNewCeroFailed(t *testing.T) {
	testCase := []string{
		"",
		"E",
		"F",
		"fasdfasdfa",
	}
	for _, c := range testCase {
		_, err := NewCero(c)
		if err == nil {
			t.Fatalf("failed test %#v", "エラーが発生しませんでした．")
		}
		if err.Error() != fmt.Sprintf("不正な値 cero: %s", c) {
			t.Fatalf("failed test %#v", "想定するエラーではありません．")
		}
	}
}

func TestIsMethods(t *testing.T) {
	for _, c := range testCases {
		switch c.Rate {
		case "A":
			result := c.isARate()
			if !result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		case "B":
			result := c.isBRate()
			if !result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		case "C":
			result := c.isCRate()
			if !result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		case "D":
			result := c.isDRate()
			if !result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}

		case "Z":
			result := c.isZRate()
			if !result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		}
	}
}

func TestValidate(t *testing.T) {
	for _, c := range testCases {
		switch c.Rate {
		case "A":
			result := c.validate(4)
			if !result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		case "B":
			result := c.validate(12)
			if !result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		case "C":
			result := c.validate(16)
			if !result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		case "D":
			result := c.validate(17)
			if !result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}

		case "Z":
			result := c.validate(18)
			if !result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		}
	}
}

func TestValidateError(t *testing.T) {
	for _, c := range testCases {
		switch c.Rate {
		case "A":
			result := c.validate(-1)
			if result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		case "B":
			result := c.validate(10)
			if result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		case "C":
			result := c.validate(13)
			if result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		case "D":
			result := c.validate(10)
			if result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}

		case "Z":
			result := c.validate(2)
			if result {
				t.Fatalf("failed test %#v", "想定する値ではありません")
			}
		}
	}
}
