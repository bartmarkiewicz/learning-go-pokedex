package main

import (
	"testing"
)

func TestCleanInput(testing *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ff  dog  name randomWord, haha multiple",
			expected: []string{"ff", "dog", "name", "randomword,", "haha", "multiple"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			testing.Errorf("Expected %d words, got %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				testing.Errorf("Expected %s, got %s", expectedWord, word)
			}
		}
	}
}
