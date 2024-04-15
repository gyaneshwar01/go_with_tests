package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://demo.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://github.com",
		"waat://demo.com",
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	want := map[string]bool{
		"https://google.com": true,
		"https://github.com": true,
		"waat://demo.com":    false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}
}
