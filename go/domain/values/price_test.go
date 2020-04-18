package values

import (
	"fmt"
	"testing"

	"github.com/dustin/go-humanize"
)

func TestNewPriceSuccess(t *testing.T) {
	testCase := []int{
		0,
		100,
		200,
		5000,
		8332,
	}
	for _, c := range testCase {
		_, err := NewPrice(c)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
	}
}

func TestNewPriceFailed(t *testing.T) {
	testCase := []int{
		-233333,
		-500,
		99999999999,
	}
	for _, c := range testCase {
		_, err := NewPrice(c)
		if err == nil {
			t.Fatalf("failed test %#v", "エラーが発生しませんでした．")
		}
		if err.Error() != "不正な価格設定です" {
			t.Fatalf("failed test %#v", "想定するエラーではありません．")
		}
	}
}

func TestPriceToString(t *testing.T) {
	testCase := []Price{
		Price(100),
		Price(200),
		Price(5000),
		Price(8332),
	}
	for _, c := range testCase {
		str := c.ToString()
		if str != fmt.Sprintf("%s円", humanize.Comma(c.ToInt())) {
			t.Fatalf("failed test: %s != %s", str, humanize.Comma(c.ToInt()))
		}
	}
}
