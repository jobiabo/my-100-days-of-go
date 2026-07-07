package ascii

import (
	"fmt"
	"os"
	"strings"
)

func ParseArguments() Command {
	args := os.Args
	var cmd = Command{}
	switch len(args) {
	case 2:
		input := args[1]

		cmd.Text = input

	case 3:
		flag := args[1]
		text := args[2]

		if !(strings.HasPrefix(flag, "--color=")) {
			fmt.Println("Usage == go run . --color=<Red> Hello, World")
			os.Exit(1)
		}

		parts := strings.Split(flag, "=")
		if parts[1] == "" {
			fmt.Println("please specify a color color must not be empty")
			fmt.Println("use --color=red")
		}
		cmd.Color = parts[1]

		cmd.Text = text

	case 4:
		flag := args[1]
		substring := args[2]
		text := args[3]

		if !(strings.HasPrefix(flag, "--color=")) {
			fmt.Println("Usage == go run . --color=<Red> Hel Hello, World")
			os.Exit(1)
		}

		parts := strings.Split(flag, "=")
		if parts[1] == "" {
			fmt.Println("please specify a color color must not be empty")
			fmt.Println("use --color=red")
		}
		cmd.Color = parts[1]
		cmd.Substring = substring
		cmd.Text = text

	default:
		fmt.Println("Usage == go run . Hello, World")
		fmt.Println("Usage == go run . --color=<Red> Hello, World")
		fmt.Println("Usage == go run . --color=<Red> Hel Hello, World")

	}
	return cmd
}
