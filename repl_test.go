package cli_test

import (
	"testing"

	"github.com/kn1ghtm0nster/go-pokedex/internal/cli"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		{
			input: "Hello, World",
			expected: []string{"hello", "world"},
		},
		{
			input: "  Leading and trailing spaces  ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
		{
			input: "Multiple   spaces between words",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "Punctuation! Should; be: removed.",
			expected: []string{"punctuation", "should", "be", "removed"},
		},
	}

	for _, c := range cases {
		actual := cli.CleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected length %d but got %d", c.input, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("For input '%s', expected word '%s' but got '%s'", c.input, expectedWord, word)
			}
		}
	}
}