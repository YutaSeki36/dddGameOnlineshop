package values

import (
	"fmt"
	"testing"
)

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
