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

func swapHard(a **string, b **string) { // สลับ value โดยที่ ที่อยู่เดิม || วิธีคิด
	fmt.Println("before a:", **a)
	fmt.Println("before &a:", a)
	fmt.Println("before b:", **b)
	fmt.Println("before &b:", b)
	fmt.Println("====")

	vA := *b
	vVA := &vA
	vB := *a
	vVB := &vB

	**a = **vVB
	**b = **vVA

	// fmt.Println(**b)
	// fmt.Println(**vVA)

	fmt.Println("after a:", **a)
	fmt.Println("after &a:", a)
	fmt.Println("after b:", **b)
	fmt.Println("after &b:", b)

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
