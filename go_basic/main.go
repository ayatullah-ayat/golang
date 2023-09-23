package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Sayan struct {
	*Person
	Power int
}

func main() {

	// goku := NewSayan("John", 3000)

	goku := &Sayan{
		Person: new(Person),
		Power:  3300,
	}
	goku.Name = "William"

	goku.introduce()
	fmt.Println(goku.Person)
	// Super(goku)
	// scores := make([]int, 0, 10)

	// cop := cap(scores)

	// for i := 0; i < 25; i++ {
	// 	scores = append(scores, i)

	// 	fmt.Println("cap", cap(scores), "cop", cop)
	// 	if cap(scores) != cop {
	// 		cop = cap(scores)
	// 		fmt.Println("cop", cop)
	// 	}
	// }
	// fmt.Println(scores, cop)
}

func (p *Person) introduce() {
	fmt.Printf("My name is: %s\n", p.Name)
}

// func NewSayan(name string, power int) *Sayan {
// 	return &Sayan{
// 		Name:  name,
// 		Power: power,
// 	}
// }

func (s *Sayan) Super() {
	s.Power += 2000
}
