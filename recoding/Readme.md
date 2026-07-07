Question 4: Dynamic ASCII Art Builder


Create an ArtBuilder type that constructs ASCII art dynamically using a builder pattern. It should support adding text with different "styles" (bold, italic, outline) through method chaining, all generated programmatically.

Requirements:

Create ArtBuilder struct with methods:

NewArtBuilder() constructor

AddText(text string) returns *ArtBuilder

SetStyle(style string) returns *ArtBuilder (supports "bold", "italic", "outline", "normal")

Build() returns string

Bold: double the
 width of characters

Italic: add forward slant to characters

Outline: add border around each character

Must support method chaining

File Structure:

ascii-art/
├── main.go
├── builder/
│   ├── builder.go
│   └── builder_test.go


