package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {

	people := make([]person, 3)

	people[0] = person{name: "Grace", age: 12}
	people[1] = person{name: "Bob", age: 23}
	people[2] = person{name: "Alice", age: 46}

	newPeople := make([]person, 0)

	for _, person := range people {
		if(person.age > 18) {
			newPeople = append(newPeople, person)
		}
	}

	fmt.Println(newPeople)
}
