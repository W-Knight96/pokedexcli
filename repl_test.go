package main

import "testing"

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
			input:    "foo bar",
			expected: []string{"foo", "bar"},
		},
		{
			input:    "   leading and trailing spaces   ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
		{
			input:    "multiple   spaces   between   words",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "singleword",
			expected: []string{"singleword"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d words, expected %d", c.input, len(actual), len(c.expected))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) returned word %q at index %d, expected %q", c.input, actual[i], i, c.expected[i])
			}
		}
	}
}
