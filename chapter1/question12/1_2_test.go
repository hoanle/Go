package main

import (
	"testing"
)

type TestCase struct {
	source   string
	permu    string
	expected bool
}

func TestIsPermutation(t *testing.T) {
	tests := []TestCase{
		{"abc", "abc", true},
		{"abc", "abd", false},
		{"1 2 3 abc", "1 abc ", false},
		{"123 321", "123321 ", true},
	}

	for _, c := range tests {
		actual := isPermutation(c.source, c.permu)
		if actual != c.expected {
			t.Fatalf("Source %s. Permutation %s. %Expected %b. Actual %b\n", c.source, c.permu, c.expected, actual)
		}
	}
}
