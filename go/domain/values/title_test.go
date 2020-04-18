package values

import "testing"

func TestNewTitleSuccess(t *testing.T) {
	testCase := []string{
		"ゼルダの伝説",
		"スーパーマリオ",
		"デスストランディング",
		"アウトラスト",
	}
	for _, c := range testCase {
		_, err := NewTitle(c)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
	}
}

func TestNewTitleFailed(t *testing.T) {
	testCase := []string{
		"",
		"スーパーマリオfffffffffffffffffffffffffffffffffffffffffffffffffff",
	}
	for _, c := range testCase {
		_, err := NewTitle(c)
		if err == nil {
			t.Fatalf("failed test %#v", "エラーが発生しませんでした．")
		}
		if err.Error() != "文字数は0文字以上50文字以下で設定してください．" {
			t.Fatalf("failed test %#v", "想定するエラーではありません．")
		}
	}
}
