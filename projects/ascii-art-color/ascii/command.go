package ascii

type Command struct {
	Color     string
	Text      string
	Substring string
}

var cmd = ParseArguments()
