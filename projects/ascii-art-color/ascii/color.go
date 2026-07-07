package ascii

func Color(selct string) string {
	var Colors = map[string]string{
		// Reset
		"reset": "\033[0m",

		// Regular colors
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}
	return Colors[selct]
}
