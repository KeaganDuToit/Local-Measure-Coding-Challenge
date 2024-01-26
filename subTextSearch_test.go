package main

import (
	"testing"
)

// TestSimpleTextSearch tests the simpleTextSearch function
// by passing in a searchAble string and a searchText string
// and comparing the actual result with the expected result
func TestSimpleTextSearch(t *testing.T) {
	searchAble := "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
	tests := []struct {
		searchText string
		expected   string
	}{
		{"peter", "1, 20, 75"},
		{"Peter", "1, 20, 75"},
		{"pick", "30, 58"},
		{"pi", "30, 37, 43, 51, 58"},
		{"z", ""},
		{"peterz", ""},
	}

	for _, test := range tests {
		actual := simpleTextSearch(searchAble, test.searchText)
		if actual != test.expected {
			t.Errorf("simpleTextSearch(%s, %s) = %s; expected %s", searchAble, test.searchText, actual, test.expected)
		}

	}
}
