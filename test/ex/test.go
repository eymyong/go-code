package ex

// import "fmt"

// type person struct {
// 	name string
// 	sex  string
// 	age  int
// }

// 	// personList := [][]person{
// 	// 	[]person{
// 	// 		person{
// 	// 			name: "art", sex: "male", age: 20,
// 	// 		},
// 	// 		person{
// 	// 			name: "pak", sex: "female", age: 25,
// 	// 		},
// 	// 		person{
// 	// 			name: "yong", sex: "female", age: 24,
// 	// 		},
// 	// 	},
// 	// 	[]person{
// 	// 		person{
// 	// 			name: "sa", sex: "female", age: 19,
// 	// 		},
// 	// 		person{
// 	// 			name: "wap", sex: "male", age: 15,
// 	// 		},
// 	// 		person{
// 	// 			name: "jj", sex: "male", age: 18,
// 	// 		},
// 	// 	},
// 	// 	[]person{
// 	// 		person{
// 	// 			name: "kalim", sex: "female", age: 30,
// 	// 		},
// 	// 		person{
// 	// 			name: "huu", sex: "female", age: 33,
// 	// 		},
// 	// 		person{
// 	// 			name: "sp", sex: "female", age: 36,
// 	// 		},
// 	// 	},
// 	// }

// 	input := []int{50, 3, 1, 6, 8, 20, 10}

// 	// fmt.Println(min(input))
// 	fmt.Println(sort(input))
// 	aa := sort(input)
// 	fmt.Println(valueMid(aa))

// }

// func valueMid(input []int) int {
// 	var output int
// 	mid := len(input) / 2
// 	if len(input)%2 == 0 {
// 		mid = mid - 1
// 	}
// 	for i := range input {
// 		v := input[i]
// 		if i == mid {
// 			output = v
// 		}
// 	}
// 	return output

// }

// func sort(input []int) []int {
// 	minValue := min(input)
// 	box1 := []int{}
// 	for _, v := range input {
// 		if v >= minValue {
// 			minValue = v
// 			if minValue <= v {
// 				box1 = append(box1, v)
// 			}
// 		}
// 	}
// 	return box1
// }

// func min(input []int) int {
// 	if len(input) == 0 {
// 		panic("len0 Yong Yong")
// 	}

// 	min := input[0]
// 	for _, v := range input {
// 		if v <= min {
// 			min = v
// 		}
// 	}
// 	return min
// }

// func age18(personList []person) []person {
// 	personNew := []person{}
// 	for i := range personList {
// 		v := personList[i]
// 		vv := personList[i].age
// 		if vv > 18 {
// 			personNew = append(personNew, v)
// 		}
// 	}
// 	return personNew
// }

// // mid of age in personList[][]person
// // [[{name: "yong", sex: "female", age: 24}],[{name: "wap", sex: "male", age: 15}],[{name: "wap", sex: "male", age: 15}]]
// // func arrayStruck2(personList [][]person) [][]person {
// // 	newArray := [][]person{}
// // 	mid := len(personList) / 2
// // 	for i := range personList {
// // 		v := personList[i].age
// // 		box := []person{}
// // 		for ii := range i {

// // 		}

// // 	}

// // 	return newArray
// // }

// func arrayStruck(personList []person) [][]person {
// 	box1 := []person{}
// 	box2 := []person{}

// 	for i := range personList {
// 		v := personList[i]
// 		va := personList[i].age
// 		vs := personList[i].sex
// 		if va <= 30 {
// 			box1 = append(box1, v)
// 		}

// 		if vs == "female" {
// 			box2 = append(box2, v)
// 		}
// 	}
// 	return [][]person{box1, box2}
// }

//input := []int{10, 20, 30, 40, 50, 60, 70, 80}
// หาตัวตรงกลาง 40
// หาตัวน้อยที่สุด 10
// หาตัวมากที่สุด 80

// loop++ i++
// loop++ range
// loop กลับหลัง
// loop next add break ถ้าถึงตัวตรงกลางให้ break
// loop back and break ถ้าถึงตัวตรงกลางให้ break
// รับ[]int , int;  (index in array[]) * (int); return array[]int
// รับ[],int ;return[]; add if else
// หาสูตรคูณ return
