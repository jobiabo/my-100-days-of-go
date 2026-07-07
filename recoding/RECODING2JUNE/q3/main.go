package main
import "fmt"

func main() {
	for i := 32; i <= 126; i++ {
		fmt.Println(i, string(i))
	}
}