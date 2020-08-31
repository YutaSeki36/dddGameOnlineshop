package values

import "testing"

func TestNewURL(t *testing.T) {
	testCases := []string{
		"https://www.nintendo.co.jp/",
		"https://www.jp.square-enix.com/",
		"http://www.capcom.co.jp/",
		"https://www.bandainamcoent.co.jp/",
	}

	for _, u := range testCases {
		_, err := NewURL(u)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
	}
}

func TestNewURLError(t *testing.T) {
	testCases := []string{
		"httpss:/www.nintendo",
		"httfps://www.jp.square-enix.com/",
		"fadfasdfasdfasd",
		"12340",
	}

	for _, u := range testCases {
		_, err := NewURL(u)
		if err == nil {
			t.Fatalf("failed test %#v", "エラーが発生しませんでした．")
		}
	}
}
