package main

import (
	"encoding/json"
	"fmt"
)

var dictEng = map[int]string{
	2: "two",
	1: "one",
	3: "three",
}

var dictThai = map[int]string{
	3: "สาม",
	1: "หนึ่ง",
	2: "สอง",
}

var dictEngGer = map[string]string{
	"one":   "eine",
	"two":   "zwei",
	"three": "drei",
}

var arr = []int{1, 2, 3}

type Person struct {
	Name string `json:"personName"`
	Age  int    `json:"personAge"`
}

func main() {
	for i := range arr {
		valueInt := arr[i]

		en := dictEng[valueInt]
		th := dictThai[valueInt]
		ger := dictEngGer[en]

		fmt.Println("en", en, "th", th, "ger", ger)
	}

	p := new(Person)
	s := `{"personNamesssssssss":"art","personAge":30}`
	err := json.Unmarshal([]byte(s), p)
	if err != nil {
		panic(err.Error())
	}

	var mm = make(map[string]interface{})
	err = json.Unmarshal([]byte(s), &mm)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", mm)
}

// ["dog", "cat", "dog"] -> {dog -> 2, cat -> 1}
func count(arr []string) map[string]int {
	m := make(map[string]int)

	for _, v := range arr {
		m[v]++
	}

	return m
}

func uniqueArr(arr []string) []string {
	result := []string{}

foo:
	for i := range arr {
		v := arr[i]

		for j := range result {
			r := result[j]

			if v == r {
				continue foo
			}
		}

		result = append(result, v)
	}

	return result
}
