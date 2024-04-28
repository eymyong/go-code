package main

import "fmt"

// type person struct {
// 	name string
// 	sex  string
// 	age  int
// }

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return 3.14 * c.r * c.r
}

// c1.isBigger(c2)
func (c1 circle) isBigger(c2 circle) bool {
	return c1.area() > c2.area()
}

func greet(p1, p2 person) {
	fmt.Println("hello", p2.name, "my name is", p1.name)
}

func (p person) greet(other person) {
	fmt.Println("hello", other.name, "my name is", p.name)
}

func haveSex(p1, p2 person) {
	fmt.Println(p1.name, "is fucking", p2.name)
}

func (p person) haveSex(other person) {
	fmt.Println(p.name, "is fucking", other.name)
}

func isOlder(p1, p2 person) bool {
	return p1.age > p2.age
}

func (p person) isOlder(other person) bool {
	return p.age > other.age
}

// func main() {
// 	pong := person{name: "pong", age: 20}
// 	ped := person{name: "ped", age: 28}

// 	c1 := circle{r: 4}
// 	c2 := circle{r: 6}

// 	haveSex(ped, pong)
// 	ped.haveSex(pong)

// 	fmt.Println(pong.isOlder(ped)) //false
// 	fmt.Println(ped.isOlder(pong)) // true

// 	fmt.Println(isOlder(ped, pong)) // false
// 	fmt.Println(isOlder(pong, ped)) // true
// 	fmt.Println(pong.isOlder(ped))  // false

// 	fmt.Println(c1.isBigger(c2))

// }

// ===================================

// type person struct {
// 	name     string
// 	sex      string
// 	age      int
// 	children []person
// }

func birthday(p person) {
	fmt.Printf("birthday: before: %+v\n", p)
	p.age++
	fmt.Printf("birthday: after: %+v\n", p)
}

func addOne(i *int) {
	fmt.Println("addOne: before", i)
	*i++
	fmt.Println("addOne: after", i)
}

// func main() {
// 	art := person{name: "art", age: 27}

// 	fmt.Printf("main: before: %+v\n", art)
// 	birthday(art)
// 	fmt.Printf("main: after: %+v\n", art)

// 	i := 10 // int
// 	addOne(&i)
// 	fmt.Println(i)
// }
