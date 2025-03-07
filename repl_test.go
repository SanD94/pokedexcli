package main

import (
	"testing"
)


func TestCleanInput(t *testing.T) {
	cases := []struct { input    string
	expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  SMALL  LetteR  ",
			expected: []string{"small", "letter"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("size do not match actual(%d) != expected(%d)", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word %s does not math with expected word %s", word, expectedWord)
			}
		}
	}
}
