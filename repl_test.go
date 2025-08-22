package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{"hello", []string{"hello"}},
		{"hello world", []string{"hello", "world"}},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) = %v; want %v", c.input, actual, c.expected)
			continue
		}
		// Iterate over the actual slice and compare each word with the expected slice
		// If they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q)[%d] = %q; want %q", c.input, i, actual[i], c.expected[i])
				continue
			}
			// If they match, you can do something with the word if needed
		}
	}

}
