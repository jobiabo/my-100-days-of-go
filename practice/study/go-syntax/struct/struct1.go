package main

import "fmt"

type student struct {
	name    string
	age     int
	sex     string
	boarder bool
}

func ModiPerson() *student {
	return &student{
		name:    "Stephanie",
		age:     16,
		sex:     "female",
		boarder: true,
	}
}

func main() {
	fmt.Println(ModiPerson())

	student2 := student{
		name:    "sandra",
		age:     15,
		sex:     "male",
		boarder: false,
	}

	fmt.Println(student2)
}
