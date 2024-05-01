package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	Sex       Sex // ซึ่งหมายความว่า feild `Sex` in struct is type Sex(แถว 16,19)
	BirthYear int
	Children  []*Person
}

// uint ค่าที่ไม่เป็นลบ
type Sex uint // จากเดิมfield Sex ควรจะมี type เป็น string// ที่สร้าง type Sex เพื่อที่จะจำกัดtypeของ Sex => เพราะจริงๆแล้ว Sex ควรจะมีแค่ Male,Female(ตามแถวที่ 18)

// ค่าคงที่
const (
	Male Sex = iota + 1 // iota จะมีค่าเป็น 0 และจะ +1 ไปเรื่อยๆตามจำนวนใน const // ที่ให้ iota +1 เพื่อที่จะให้ const[0] = 1
	Female
	gar
	tom
	// 2
) // ซึ่งหมายความว่า feild Sex in struct is type Sex(แถว 16,19) // ซึ่งถ้า Struct กำหนดให้ s.Sex เป็น 1,2 ก็จะกลายเป็น,Male Female ทันที :: แต่ถ้า s.Sex เป็นอะไรก็ตามที่นอกเหลือจาก iota in `const` => s.Sex ก็จะกลายเป็นค่านั้นแทน

// บอกว่าเป็นเพศอะไร
func (s Sex) String() string { //สร้าง `method` ของ `String type Sex` // return string // ส้รางมาเพื่อเมื่อมีการเรียกใช้ type Sex จะบอกเลยว่าคนนั้นเป็นเพศไหน
	if s == Male { // (`method` ของ String ก็คือ `func string` ที่รับตัวแปรที่มี type ของ method เท่านั้น) :: (คำถามคือ ทำไมยังไม่ได้เรียกใช้ method) แต่ struct ที่มี field Sex ถึงถูกทำงานและ return ค่าออกไปได้
		return "Male"
	}
	if s == Female {
		return "Female"
	}
	if s == 3 || s == 4 {
		return "LGTV+"
	}
	return "ไม่ทราบเพศ"
}

// `Age()` return อายุของ person
func (a Person) Age() int {
	return time.Now().Year() - a.BirthYear
}

// ถามอายุ
func (i Person) IsAdult() bool {
	if i.Age() >= 20 {
		return true
	}
	return false
}

//===================================================================================================================
//ฝึกทำเพิ่มคือ ให้ชื่อของลูกเป็น 2ตัวแรกของชื่อพ่อ + 29หลังของชื่อแม่ // และกำหนดเงื่อนไขเพิ่มคือให้สลับกันโดย ถ้า ช หรือ ญ ใครเป็นคนเย็ด //
//===================================================================================================================

// `Yed()` return person หากเป็นคนละเพศกัน เย็ดกันจะได้ลูกที่มีชื่อ ตามตัวแปรที่ส่งเข้าไป
// และมีเพศตามคนเย็ด โดยที่หากเบ็ดแล้วมีลูก c้hildren ของคนที่เย็ดกัน ต้องเพิ่มขึ้นมา

// method รับ *person, string และ return *person // คือจะให้ทำลูกออกมาถ้าเพศไม่เท่ากัน และลูกที่ออกมาได้ต้องเข้าไปอยู่ใน struck ของคนที่เย็ดกันแล้วมีลูก
func (konyed *Person) Yed(tokyed *Person, s string) *Person {
	//baby :=[]Person{}

	if konyed.Sex != tokyed.Sex { //เงื่อนไขแรกถ้าคนละเพศกันให้มีลูกได้
		baby := Person{ // สร้างตัวแป็นที่เป็น struct เพื่อมารับข้อมูลของลูก
			Name:      s,                 // ชื่อลูก = ตัวแปร string ที่รับเข้ามา
			Sex:       konyed.Sex,        // เพศของลูก = เพศของคนเย็ด
			BirthYear: time.Now().Year(), // ปีที่เกิด
		}
		fmt.Println("child of", konyed.Name, "yed", tokyed.Name, "::", baby)

		konyed.Children = append(konyed.Children, &baby) //ให้ลูกค้าไปอยู่ในคนที่เย็ดกัน
		tokyed.Children = append(tokyed.Children, &baby) //ให้ลูกค้าไปอยู่ในคนที่เย็ดกัน
		return &baby
		// b1 := s , konyed.Sex
		// baby = append(baby, b1)
		//return konyed.Children[]&Person{}
	}

	return nil // ที่ต้อง return nil เพราะว่า //
}

func printNotNil(p *Person) { // สร้างfuncที่เอาไว้เช็คว่า *person ที่รับเข้ามา เป็น nil หรือไม่
	if p != nil {
		fmt.Println("printNotNil", p)
	}
}

func checkSex(expect Sex, actual *Person) bool { // สร้างfuncที่เอาไว้เช็คSex โดยรับ  Sex และ *person โดยให้ return bool
	if actual != nil {

		if expect == actual.Sex {
			return true
		}
	}

	return false
}

func main() {

	yong := Person{
		Name:      "yong",
		Sex:       Male,
		BirthYear: 1997,
		Children:  nil,
	}
	_ = yong

	noon := Person{
		Name:      "noon",
		Sex:       Female,
		BirthYear: 1997,
		Children:  nil,
	}
	_ = noon

	pak := Person{
		Name:      "pak",
		Sex:       1,
		BirthYear: 1997,
		Children:  nil,
	}
	_ = pak

	namwhan := Person{
		Name:      "namwhan",
		Sex:       Female,
		BirthYear: 1997,
		Children:  nil,
	}
	_ = namwhan

	//fmt.Println(yong.Yed(noon, aaa))
	fmt.Println(yong)
	a := yong.Yed(&noon, "black")
	fmt.Println(yong)
	fmt.Println(a)
	// fmt.Println("===")
	// fmt.Println(yong, *a)

	//b := yong.Yed(&pak, "k1")

	//c := noon.Yed(&yong, "summer")

	//fmt.Println(b)

	// fmt.Println("c.Sex Female=", checkSex(Female, c))
	// fmt.Println("a.Sex Male=", checkSex(Male, a))

	// fmt.Println("yong =\t", yong)
	// fmt.Println("Yong Children 1 =\t", *yong.Children[0])
	// printNotNil(a)
	// printNotNil(b)

}
