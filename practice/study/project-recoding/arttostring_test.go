package main

import "testing"

func TestArtToString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "single zero",
			input: "" +
				" ___ \n" +
				"|   |\n" +
				"|   |\n" +
				"|   |\n" +
				"|___|",
			expected: "0",
		},
		{
			name: "single one",
			input: "" +
				"     \n" +
				"    |\n" +
				"    |\n" +
				"    |\n" +
				"    |",
			expected: "1",
		},
		{
			name: "multiple digits",
			input: "" +
				"      ___ \n" +
				"    ||   |\n" +
				"    ||   |\n" +
				"    ||   |\n" +
				"    ||___|",
			expected: "10",
		},
		{
			name: "invalid height",
			input: "" +
				" ___ \n" +
				"|   |\n" +
				"|___|",
			expected: "",
		},
		{
			name: "invalid width",
			input: "" +
				" ___ \n" +
				"|  |\n" +
				"|  |\n" +
				"|  |\n" +
				"|__|",
			expected: "",
		},
		{
			name: "unknown pattern",
			input: "" +
				"xxxxx\n" +
				"xxxxx\n" +
				"xxxxx\n" +
				"xxxxx\n" +
				"xxxxx",
			expected: "",
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ArtToString(tt.input)

			if result != tt.expected {
				t.Errorf("ArtToString() = %q, want %q", result, tt.expected)
			}
		})
	}
}
