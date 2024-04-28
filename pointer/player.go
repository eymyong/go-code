package pointer

import "fmt"

// var a *string  //
// var b string
// s := "kuy" // สร้างตัวแปร 's' มีtypeเป็น string
// a = &s     // 'a' เป็น '*string' ชี้ไปที่ 's'
// b = *a     // 'b' deref '*a'

// c := &a  // **string
// d := **c // d เป็น string deref a

// e := &c
// //f := ***e

//fmt.Println("f ", f)

type Kuy struct{}

type player struct {
	level int
}

func (p player) upLevel1() player {
	p.level += 1 // copy
	return p     // copy
}

func (p player) upLevel2() *player {
	p.level += 1
	return &p
}

func (p *player) upLevel3() player {
	p.level += 1
	return *p
}

func (p *player) upLevel4() *player {
	p.level += 1 // change p in main
	pCopy := *p  // copy from p in main to p2
	return &pCopy

}

func changeValue(n *int) { // สร้าง'func' ที่รับ 'n *int' แล้วเปลี่ยนค่าโดยที่ไม่เปลี่ยนที่อยู่
	*n = 70
}

func addOnePtr(i *int) int {
	*i += 1
	return *i

	//==== in main ====
	//i := 69
	// addOnePtr(&i)  //ส่ง i เข้ามา
	// fmt.Println(i) // 70
	// fmt.Println(addOnePtr(&i)) //return ถ้าfuncไม่return จะPrintln ไม่ได้
}

func changeInt(i ***int) *int {
	a := **i
	return a
}

func swapValue(a *int, b *int) { // init 'func' ที่รับ '(a,b *int)' แล้วให้ value 'a' swap 'b' โดยที่ '&a,&b' ไม่เปลี่ยน
	fmt.Println("pointer a:\t", a)
	fmt.Println("value a:\t", *a)
	fmt.Println("pointer b:\t", b)
	fmt.Println("value b:\t", *b)
	fmt.Println("=====")

	va := *a // init 'va' to be value of 'a' >>สร้างตัวแปร 'va' มารับ value ของ 'a' (deref 'a' = '*a' = เอาvalueออกจาก 'a')
	vb := *b // init 'vb' to be value of 'b'
	*a = vb  // ให้ '*a'(value ของ a ที่อยู่เดิม)43 = 'vb'(value ใหม่)69
	*b = va  // ให้ '*b'(69) = 'va'(43)

	fmt.Println("swap pointer a:\t", a)
	fmt.Println("swap value a:\t", *a)
	fmt.Println("swap pointer b:\t", b)
	fmt.Println("swap value b:\t", *b)

	//=== in main ===
	// yong := 69
	// pak := 43
	// swapValue(&yong, &pak)

}

func eym(hee **[]int) {
	w := []int{1, 2, 3, 4}
	h := 69
	//t := &h
	w = append(w, h)
	//fmt.Println(w)
	//ww := &w

	//=== in main ===
	// var v **[]int

	// eym(v)
	// fmt.Println() // [1,2,3,4,69]

}

func kuy() {

	//=== in main ===
	// var h []string
	// //kuy(h, 2)
	// fmt.Println(h) //ใส่int 2 ต้องมี [hee,hee] , ใส่int 3 ต้องมี [hee,hee,hee] ,ถ้าintติด- [kuy]

}

func addPtrMain() {
	s := 8
	t := 9

	pS := &s             // ประกาศ `pS` มารับ address ของ s ซึ่ง `pS` เป็น type *int
	pT := &t             // ประกาศ `pT` มารับ address ของ  t ซึ่ง `pT` เป็น type *int
	a1 := addPtr(pS, pT) // return 17
	fmt.Println(a1)

	pPT := &pT                // ประกาศ `pPT` มารับ address ของ pT ซึ่ง `pPT` เป็น type **int
	a2 := addPtrHard(pS, pPT) // return 17
	fmt.Println(a2)

}

func addPtr(a, b *int) int {
	vA := *a // vA เอา value ของ a ออกมา ซึ่ง `vA` เป็น type int
	vB := *b //  vB deref `b` => `vB` type int ---- type of `vB` is int
	return vA + vB
}

func addPtrHard(a *int, b **int) int {
	vA := *a
	// `vB` เอา value ของ b ออกมา ซึ่ง `vB` เป็น type *int
	// ที่เป็น *int ไม่ใช่ int เพราะ b ที่รับเข้ามาเป็น type **int > deref ออกมาจึงเป็น *int
	vB := *b
	// `vVB` เอา value ของ `vB` ออกมา ซึ่ง `vVB` เป็น type **int
	vVB := *vB
	return vA + vVB
}

func swapHard(a **string, b **string) { // สลับ value โดยที่ ที่อยู่เดิม
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

}

// j := 10      // 1
// p1 := &j     // 2           // p1 = 0x7
// p2 := foo(j) // 1           // p2 = 0x8

// func foo(j int){

// }

// fmt.Printf("p1 %p p2 %p\n", p1, p2)
// *p2 = 7
// fmt.Println("j", j)

// var p3 *int // nil deref = panic
// var p4 *int = new(int) // จะไปหาที่ว่าง แล้วสร้าง default int (0) แล้วเอาที่อยู่อันนั้นมาให้ ซึ่งถ้า deref `p4` = 0
// fmt.Println(p3, p4, *p3)
