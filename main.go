package main

import (
	"fmt"
)

type person struct {
	name     string
	sex      string
	age      int
	children []*person
}

func main() {
	k1 := person{
		name: "k1",
		sex:  "m",
	}

	k2 := person{
		name: "k2",
		sex:  "f",
	}

	dad := person{
		name:     "dad",
		sex:      "m",
		children: []*person{&k1, &k2}, // จากเดิม children: []person{k1, k2} || เปลี่ยนเป็น children: []*person{&k1, &k2} เพื่อให้ชี้ไปหา person เพราะเราต้องการเปลี่ยน k1 ให้มันเปลี่ยนทุกที่
	}

	mom := person{
		name:     "mom",
		sex:      "f",
		children: []*person{&k1, &k2},
	}

	k1.newName("kuy")                        // เปลี่ยน k1 เป็น kuy โดยผ่าน method `newName`
	fmt.Println("k1 kid", k1)                // เปลี่ยนแล้ว
	fmt.Println("dad kids", dad.children[0]) //
	fmt.Println("mom kids", mom.children[0]) //
}

func (p *person) newName(s string) {
	p.name = s
}

// func newName(p *person, s string) {

// }
