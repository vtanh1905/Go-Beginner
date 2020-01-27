package main

import "fmt"

type Student struct {
	id   int
	name string
}

// Khai bao method
// func (t Type) methodName(params) returns {// code}
func (s Student) getName() string {
	return s.name
}

/*
	method != func ở điểm nào ?
	method thuộc về class => nó là phương thức trong OOP
*/

// (t Type) => Receiver
// 1 value receiver (không thay đổi giá trị)
func (s Student) setName1(str string) {
	s.name = str
}

// 2 pointer receiver (thay đổi giá trị)
func (s *Student) setName2(str string) {
	s.name = str
}

// non-struct
// Chúng ta define theo kiểu dữ liệu của golang
type String string

func (s String) getFirstCharacter() string {
	return string(s[0])
}

func main() {
	student1 := Student{
		id:   1,
		name: "Nguyễn Văn A",
	}
	fmt.Println(student1.getName())

	student1.setName1("ABCD")
	fmt.Println(student1.getName())

	student1.setName2("ABCD")
	fmt.Println(student1.getName())

	myName := String("Tuấn Anh")
	fmt.Println(myName.getFirstCharacter())

}
