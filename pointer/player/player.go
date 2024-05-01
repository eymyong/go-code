package pointer

import "fmt"

type Kuy struct{}

type player struct {
	name  string
	level int
}

func (p *player) upLevel1() player { // รับ `*player` || p.level += 1 || return person ที่ value ถูกเปลี่ยนจาก address ไม่ใช่การ Copy
	p.level += 1
	return *p
}

func (p *player) upLevel2() *player { // รับ `*player` || p.level += 1 || return `*person`` ซึ่ง value จะออกไปเป็น &(value)
	p.level += 1 // change p in main
	pCopy := *p  // copy from p in main to p2
	return &pCopy

}
func (p player) upLevel3() player { // เปลี่ยน value in func || แต่ไม่ได้เปลี่ยนที่ address ทำให้พอออกจาก func player.level ยังไม่ถูกเปลี่ยน

	p.level += 1 // copy
	return p     // copy
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
	// fmt.Println(addOnePtr(&i)) //return ถ้าfuncไม่ประกาศreturn จะPrintln ไม่ได้
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

//==========================================================================================//

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

//================================================================================//

func playerInMain() {
	p1 := player{
		name:  "yong",
		level: 1,
	}

	fmt.Println("before level =", p1.level)
	fmt.Println("before address p1 =", &p1.level)
	fmt.Println("===")

	//p1.upLevel1()
	p2 := p1.upLevel2()
	fmt.Println(*p2)

	fmt.Println("after level =", p1.level)
	fmt.Println("after address p1 =", &p1.level)

}

// array := []int{10, 20, 30, 40, 50}
// b := 2

// changeValue(array, b)
// fmt.Println(array)

// func changeValue(n []int, b int) {
// // c := n

// for i := 0; i < len(n); i++ {
// 	n[i] = n[i] * b
// }

// }

// func eym(a *[]int, b int) {
// 	*a = append(*a, b)
// }
