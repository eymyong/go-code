package main

import "fmt"

func main() {

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

	vA := *b // vA จากหม่อนเป็น อาท
	vVA := &vA
	vB := *a
	vVB := &vB

	**a = **vVA
	**b = **vVB

	fmt.Println(**a)
	fmt.Println(**b)
	fmt.Println(**vVA)

	fmt.Println(**vVB)

	fmt.Println("after a:", **a)
	fmt.Println("after &a:", a)
	fmt.Println("after b:", **b)
	fmt.Println("after &b:", b)

}
