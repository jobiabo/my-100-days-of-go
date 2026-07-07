package main

import "fmt"

type person struct {
	name string
}

func Newname() *person {
	return &person{
		name: "",
	}
}

func (p *person) talk(name string) *person {

	p.name = name
	fmt.Println("my name is ", p.name)
	return p
}
func main() {
	Newname().talk("hassan")
}
