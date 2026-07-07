```go
package helper

func Punct(input string) string {
	runes := []rune(input)
	result := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' && i+1 < len(runes) {
			switch runes[i+1] {
			case ':', '?', '.', ',':
				continue
			}
		}
		result = append(result, runes[i])
	}

	return string(result)
}
```