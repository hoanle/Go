package question11

import (
	"testing"
)

type TestCase struct {
	input    string
	expected bool
}

func TestIsUnique(t *testing.T) {
	tests := []TestCase{
		{"abcde", true},
		{"123abc", true},
		{"123455abca", false},
		{"", true},
		{" ", true},
		{"  ", false},
	}

	for _, c := range tests {
		actual := isUnique(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected %b. Actual %b\n", c.input, c.expected, actual)
		}
	}
}
