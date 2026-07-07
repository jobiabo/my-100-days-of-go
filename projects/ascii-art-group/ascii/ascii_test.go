package ascii

import (
	"os"
	"strings"
	"testing"
)

func createTestBanner(t *testing.T) string {
	t.Helper()

	tmp, err := os.CreateTemp("", "banner*.txt")
	if err != nil {
		t.Fatal(err)
	}

	var builder strings.Builder

	// 95 printable ASCII characters × 9 lines = 855 lines
	// Final newline creates the 856th empty line after Split().
	for ch := 32; ch <= 126; ch++ {
		builder.WriteString("\n") // separator line
		for i := 0; i < 8; i++ {
			builder.WriteString(string(rune(ch)))
			builder.WriteString("\n")
		}
	}

	if _, err := tmp.WriteString(builder.String()); err != nil {
		t.Fatal(err)
	}

	if err := tmp.Close(); err != nil {
		t.Fatal(err)
	}

	return tmp.Name()
}

func TestLoadBannerSuccess(t *testing.T) {
	file := createTestBanner(t)
	defer os.Remove(file)

	banner, err := LoadBanner(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(banner) != 95 {
		t.Fatalf("expected 95 characters, got %d", len(banner))
	}

	if len(banner['A']) != 8 {
		t.Fatalf("expected 8 rows for A, got %d", len(banner['A']))
	}

	for _, row := range banner['A'] {
		if row != "A" {
			t.Fatalf("expected row to contain A, got %q", row)
		}
	}
}

func TestLoadBannerFileNotFound(t *testing.T) {
	_, err := LoadBanner("does-not-exist.txt")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestLoadBannerInvalidLineCount(t *testing.T) {
	tmp, err := os.CreateTemp("", "badbanner*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	tmp.WriteString("only\nthree\nlines\n")
	tmp.Close()

	_, err = LoadBanner(tmp.Name())
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestRenderText(t *testing.T) {
	banner := map[rune][]string{
		'A': {
			"A", "A", "A", "A",
			"A", "A", "A", "A",
		},
	}

	got := RenderText([]string{"A"}, banner)

	expected := ""
	for i := 0; i < 8; i++ {
		expected += "A\n"
	}

	if got != expected {
		t.Errorf("expected %q got %q", expected, got)
	}
}

func TestRenderTextMultipleCharacters(t *testing.T) {
	banner := map[rune][]string{
		'A': {"A", "A", "A", "A", "A", "A", "A", "A"},
		'B': {"B", "B", "B", "B", "B", "B", "B", "B"},
	}

	got := RenderText([]string{"AB"}, banner)

	expected := ""
	for i := 0; i < 8; i++ {
		expected += "AB\n"
	}

	if got != expected {
		t.Errorf("expected %q got %q", expected, got)
	}
}

func TestRenderTextEmptyLine(t *testing.T) {
	got := RenderText([]string{""}, map[rune][]string{})

	if got != "\n" {
		t.Errorf("expected newline, got %q", got)
	}
}

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		args []string
		want bool
	}{
		{[]string{}, false},
		{[]string{"prog"}, false},
		{[]string{"prog", "hi"}, true},
		{[]string{"prog", "hi", "standard"}, true},
		{[]string{"a", "b", "c", "d"}, false},
	}

	for _, tt := range tests {
		if got := ValidateArgs(tt.args); got != tt.want {
			t.Errorf("ValidateArgs(%v)=%v want %v", tt.args, got, tt.want)
		}
	}
}

func TestPrepareInput(t *testing.T) {
	input := "Hello\\nWorld\\t!"
	expected := "Hello\nWorld    !"

	got := PrepareInput(input)

	if got != expected {
		t.Errorf("expected %q got %q", expected, got)
	}
}

func TestIsOnlyNewlines(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"\n", true},
		{"\n\n\n", true},
		{"abc", false},
		{"\na", false},
	}

	for _, tt := range tests {
		if got := IsOnlyNewlines(tt.input); got != tt.want {
			t.Errorf("IsOnlyNewlines(%q)=%v want %v", tt.input, got, tt.want)
		}
	}
}

func TestSplitLines(t *testing.T) {
	input := "Hello\nWorld\nGo"

	expected := []string{"Hello", "World", "Go"}

	got := SplitLines(input)

	if len(got) != len(expected) {
		t.Fatalf("expected %d lines got %d", len(expected), len(got))
	}

	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("line %d expected %q got %q", i, expected[i], got[i])
		}
	}
}
