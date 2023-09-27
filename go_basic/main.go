package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Person struct {
	Name string
}

type Sayan struct {
	*Person
	Power int
}

func main() {

	scores := make([]int, 100)

	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, len(scores))

	copy(worst, scores[5:])
	fmt.Println(scores)
	fmt.Println(worst)
	// scores := []int{1, 2, 3, 4, 5}

	// scores = removeAtIndex(scores, 2)

	// slice := scores[2:4]
	// slice[0] = 9999
	// fmt.Println(slice)
	// haystack := "the spiec must flow"
	// slice := strings.Index(haystack[5:], "m")
	// fmt.Println(scores)
	// goku := NewSayan("John", 3000)

	// goku := &Sayan{
	// 	Person: new(Person),
	// 	Power:  3300,
	// }
	// goku.Name = "William"

	// goku.introduce()
	// fmt.Println(goku.Person)
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

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1

	fmt.Println("lastIndex", lastIndex)
	source[index], source[lastIndex] = source[lastIndex], source[index]
	fmt.Println("swap_src", source)
	return source[:lastIndex]
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
