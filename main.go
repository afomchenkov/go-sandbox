package main

import (
	"fmt"
)

const PACKAGE_NAME string = "MAIN"

func plus(a int, b int) int {
	return a + b
}

func plus1(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 8
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 33
	return &p
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func new_card() string {
	return "Five of Dimonds"
}

func updateMySlice(s []string) {
	s[0] = "Bye"
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// go run main.go
func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateMySlice(mySlice)
	fmt.Println((mySlice))
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	fmt.Println(card)

	// cards := deck{"one", "two", "three"}
	cards := newDeck()
	cards.shuffle()
	// cards = append(cards, "four")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()

	s := make([]string, 3)
	fmt.Println(s)

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(s)

	var another [3]int
	fmt.Println(another)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	m := make(map[string]int)
	var key1 = "key1"
	m[key1] = 34

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// fmt.Println(person{"Bob", 20})
	// fmt.Println(person{name: "Alice", age: 30})
	// fmt.Println(person{name: "Fred"})
	// fmt.Println(&person{name: "Ann", age: 40})
	// fmt.Println(newPerson("Jon"))

	// type conversion
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))

	alex := human{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	// by default int - 0, string = "", bool = false
	// var john human
	// fmt.Println(john)

	jim := human{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 10099,
		},
	}

	jimPointer := &jim
	// fmt.Printf("%+v", jim)
	jim.print()
	jimPointer.updateName("jimmy")
	jim.print()

	fmt.Println("-------------")

	// create maps

	// var colors map[string]string
	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors["white"] = "#ffffff"
	colors["green"] = "#4bf745"
	colors["red"] = "#ff0000"
	// delete(colors, "white")
	// fmt.Println(colors)
	printMap(colors)

	fmt.Println("-------------")

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	LoadData()

	tr1 := triangle{base: 10, height: 10}
	sq1 := square{sideLength: 10}

	printArea(tr1)
	printArea(sq1)
}

func (pointerToHuman *human) updateName(newFirstName string) {
	(*pointerToHuman).firstName = newFirstName
}

// pass-by-value -M makes a copy and updates name, the original object is intact
// func (p human) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (p human) print() {
	fmt.Printf("%+v", p)
}
