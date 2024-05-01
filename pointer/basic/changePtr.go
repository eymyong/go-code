package basic

import "fmt"

func basic() {
	fmt.Println("aa")
}

func basicMain() {

	a := "mond"
	b := "art"

	aa := &a
	bb := &b

	//swapHard(&a, &b)
	swapHard(&aa, &bb)

}

func swapHard(a **string, b **string) { // สลับ value โดยที่ ที่อยู่เดิม || วิธีคิดคือเราจะสลับvalue ก็ต้องเข้าถึงvalueให้ได้ก่อน
	fmt.Println("a ก่อน", a)        //&a
	fmt.Println("a value ก่อน", *a) // value a
	fmt.Println("b ก่อน", b)        //&a
	fmt.Println("b value ก่อน", *b)
	fmt.Println("====")
	//newA := a
	newValueA := *b   // สร้างตัวแปรใหม่ = value ของ b ซึ่ง newValueA เป็น type *int
	nnA := &newValueA // สร้างตัวแปรใหม่เป็น *newValueA ซึ่ง nnA เป็น type **int
	newValueB := *a
	nnB := &newValueB

	*a = *nnA // value ของ a = value ของ *nnA // deref`*nnA` = *b
	*b = *nnB

	fmt.Println("a หลัง", a)
	fmt.Println("a value หลัง", *a)
	fmt.Println("b หลัง", b)
	fmt.Println("b value หลัง", *b)

	//in main//
	// s := "mond"
	// t := "art"

	// ss := &s
	// tt := &t

	// swapHard(&ss, &tt)

	//=============================================================//

}

func addPtr(a, b *int) int { // รับ a , b type *int || return int โดยที่เปลี่ยนข้อมูลจาก `&` ไม่ Copy
	var new int
	vA := *a
	vB := *b
	new = *&vA + *&vB
	return new
}

func addPtrHard(a *int, b **int) int { // รับ a *int, b **int || return int โดยที่เปลี่ยนข้อมูลจาก `&` ไม่ Copy
	vA := *a   // สร้าง `vA` มารับ value ของ `*a` ซึ่งมี type เป็น int
	vB := *b   // สร้าง `vB` มารับ value ของ `*b` ซึ่งมี type เป็น int
	vB2 := &vB // สร้าง `vB2` มารับ value ของ *vB โดยที่ `vB` มี type เป็น *int >> `vB2` type **int
	return vA + **vB2
}

func changeValue(n *int) { // สร้าง'func' ที่รับ 'n *int' แล้วเปลี่ยนค่าโดยที่ไม่เปลี่ยนที่อยู่
	*n = 50 //
}

func addOnePtr(n *int) int {
	*n += 10
	return *n //return ถ้าfuncไม่ประกาศreturn จะปริ้น fmt.Println(addOnePtr(&i))  ใน main ไม่ได้
}

func changeInt(i ***int) *int {
	a := **i
	return a
}

// func changeValue(n *[]int) []int { // สร้าง'func' ที่รับ 'n []*int' แล้วเปลี่ยนค่าโดยที่ไม่เปลี่ยนที่อยู่
// 	box1 := &n
// 	for _, v := range *n {
// 		vv := v * 2
// 		box1 = append(box1, vv)
// 	}

// 	return *n
// }

//var array2 []int{10, 20, 30, 40, 50}
//array1 := []int{10, 20, 30, 40, 50}
//fmt.Println("before array1 =", array1)

//array2 := &array1

//fmt.Println(changeValue(&array1))
//fmt.Println("after array1 =",array1)
