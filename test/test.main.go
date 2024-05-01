package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	Sex       Sex
	BirthYear int
	Children  []*Person
}

type Sex int

// ค่าคงที่
const (
	// ไอโอต้า
	Male Sex = iota + 1
	Female
	LGTV
)

const (
	_Male   Sex = 1
	_Female Sex = 2
	_LGTV   Sex = 3
)

func (s Sex) String() string {
	if s == Male {
		return "Male"
	}
	if s == Female {
		return "Female"
	}

	if s == LGTV {
		return "LGTV+"
	}

	return "ไม่รู้เพศ"

}

func (a Person) Age() int {

	return time.Now().Year() - a.BirthYear
}

func (i Person) IsAdult() bool {
	if i.Age() >= 20 {
		return true
	}
	return false
}

func (konyed *Person) Yed(tokyed *Person, s string) *Person {
	//baby :=[]Person{}

	if konyed.Sex != tokyed.Sex {
		baby := Person{
			Name:      s,
			Sex:       konyed.Sex,
			BirthYear: time.Now().Year(),
		}
		fmt.Println("child of", konyed.Name, "yed", tokyed.Name, "::", baby)

		konyed.Children = append(konyed.Children, &baby)
		tokyed.Children = append(tokyed.Children, &baby)
		return &baby
	}

	return nil
}

func printNotNil(p *Person) {
	if p != nil {
		fmt.Println("printNotNil", p)
	}
}

func checkSex(expect Sex, actual *Person) bool {
	if actual != nil {

		if expect == actual.Sex {
			return true
		}
	}

	return false
}

func (p *Person) Adopt(kid *Person) {
	p.Children = append(p.Children, kid)
	fmt.Println(p.Name, "รับ", kid.Name, "มาเลี้ยง")
}

func mapOddEven(arr []int) map[int]string {
	m := make(map[int]string)
	for i := range arr {
		v := arr[i]
		if v%2 != 0 {
			m[v] = "odd"
		} else {
			m[v] = "even"
		}

		// if v%2 != 0 {
		// 	m[v] = "odd"
		// 	continue
		// }

		// m[v] = "even"
	}
	return m
}

func main() {

	numbers := []int{1, 3, 4, 3, 3, 2, 1}

	uniqueNumbers := unique(numbers)

	fmt.Println(uniqueNumbers) // [1,2,3,4] (ห้ามมีตัวซ้ำ)

}

func uniqueArt(arr []int) []int {
	result := []int{}
	seen := make(map[int]struct{})

	for _, v := range arr {
		_, ok := seen[v] // ท่า _,ok => `ok` คือ bool ซึ่งเอาไว้เช็คว่าใน `map seen` มี key v หรือไม่
		if ok {
			continue
		}

		seen[v] = struct{}{}
		result = append(result, v)
	}

	return result
}

func unique(arr []int) []int {
	box1 := []int{}
	a := map[int]bool{}
	for _, v := range arr {
		_, ok := a[v]
		if ok {
			continue
		}

		a[v] = true
		box1 = append(box1, v)
	}

	return box1
}
