package pointer

// pointer  จะมี '*'(ตัวชี้),'&'(ที่อยู่)
// '*' เรียกว่า star , '&' เรียกว่าแอ่น
//ตัวแปร `*` ไม่ได้เอาไว้เก็บค่า string,int, float64 =>แต่เอาไว้เก็บตำแหน่ง
// ถ้ากำหนด '*' ขึ้นมา คือกำหนดว่าเป็น pointer ซึ่งเป็นการ(ชี้ไปหา) >> ตัวที่ถูกชี้ต้องต้องใส่ '&' (เพื่อบอกว่าชี้ไปหาที่อยู่ตรงไหน)
// อยู่หน้าตัวแปร '*input'  จะเป็นการ deref(คือการเรียกหาค่าที่อยู่ในตัวแปร หรือเป็นการคืนค่าที่อยู่ในตัวแปร)
// อยู่หน้าtype '*string' จะเป็นการ ref(คือการเรียกหาที่อยู่ หรือเป็นการคืนค่าที่อยู่ของtypeที่ถูกชี้) (รับแค่& ไม่สามารถรับvalueได้)

// func ex(a *int) คือfuncรับ 'a type *int' ==> เรียกfunc จะต้องเป็น ex(&a)
// s := 10, a := &s, b = *a, b = 10,  ==>  ใส่'*ข้างหน้าตัวแปรเพื่อจะ deref &ที่อยู่หน้าตัวแปร
// 101      102     103
//เช็คย้อนกลับว่าตัวแปรที่เราใช้อยู่เป็น typeอะไร ==> นับย้อน&, e

// var a *string //
// var b string

// s := "kuy" // สร้างตัวแปร 's' มีtypeเป็น string
// a = &s     // 'a' เป็น '*string' ชี้ไปที่ 's'
// b = *a     // 'b' deref '*a'

// c := &a  // **string
// d := **c // d เป็น string deref a

// e := &c
//f := ***e
