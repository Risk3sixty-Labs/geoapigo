package geoapigo

import (
	"testing"
)

func TestGet(t *testing.T) {
	res1, err := Get("8.8.8.8")
	if err != nil {
		t.Error("Error returned from Get.", err)
	}

	if res1.IP != "8.8.8.8" {
		t.Error("Response did not include IP address.", res1)
	}

	if res1.Country != "US" {
		t.Error("Response did not include correct country.", res1)
	}
}
