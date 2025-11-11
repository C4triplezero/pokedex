package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   John   Doe   ",
			expected: []string{"john", "doe"},
		},
		{
			input:    "1 2 3 4 5 6",
			expected: []string{"1", "2", "3", "4", "5", "6"},
		},
		{
			input:    "    a    b    c    ",
			expected: []string{"a", "b", "c"},
		},
		{
			input:    "  HELLO  WORLD  ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		aC := len(actual)
		eC := len(c.expected)
		if aC != eC {
			t.Errorf("length of response: %v, %v does not match length of expected: %v, %v", aC, actual, eC, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("response word '%v' does not match expected word '%v'", word, expectedWord)
			}
		}
	}
}
