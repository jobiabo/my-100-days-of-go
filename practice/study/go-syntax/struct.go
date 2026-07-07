package main

import "fmt"

type Employee struct {
	name     string
	age      int
	sex      string
	isRemote bool
}

func main() {

	employee1 := Employee{
		name:     "julian",
		age:      25,
		sex:      "female",
		isRemote: true,
	}

	emmployee2 := Employee{
		name:     "joel",
		age:      22,
		sex:      "male",
		isRemote: true,
	}
	emmployee2.updateName("mr joel")
	employee1.updateName("joyce")

	fmt.Println(employee1)
	fmt.Println(emmployee2)

}

func (e *Employee) updateName(newName string) {
	e.name = newName
}
