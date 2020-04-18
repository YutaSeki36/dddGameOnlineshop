package values

import "testing"

func TestExampleSuccess(t *testing.T) {
	_, err := NewTitle("hoge")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
