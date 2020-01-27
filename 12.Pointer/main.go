package main

import "fmt"

func main() {
	fmt.Println("Pointer")

	// Khai báo
	var pointer1 *int
	fmt.Println(pointer1)

	// Khai bao con trỏ = new(type) <=> var p *int
	pointer2 := new(int)
	fmt.Println(pointer2)

	/*
		Zero value of pointer
		Khac biet giua var *int va new(int)
		TH1 thi con trỏ chưa được cấp bổ nhớ, còn pointer2 thi đã được cấp bộ nhớ
	*/

	// Lay type cua pointer
	fmt.Printf("%T \n", pointer1)
	fmt.Printf("%T \n", pointer2)

	// Cách trỏ tới 1 biến khac
	a := 1000
	var pointer3 *int
	pointer3 = &a
	fmt.Println(*pointer3)
	*pointer3 = 696969
	fmt.Println("a", a)
	fmt.Println("*pointer3", *pointer3)

	// Cách lấy dia tri và giá tri con trỏ
	// & để lấy đia chi
	// * để lấy giá trị

	// Pointer trỏ tới 1 array
	arr := [3]int{1, 2, 3}
	var pointer4 *[3]int
	pointer4 = &arr
	fmt.Println(pointer4)
	fmt.Println(*pointer4)
	fmt.Println(&pointer4)
	fmt.Println()
	// truyen pointer vao 1 ham`
	pointer5 := new(int)
	fmt.Println(*pointer5)
	Truyen_Pointer(pointer5)
	fmt.Println(*pointer5)
	fmt.Println()

	b := 100
	fmt.Println(b)
	Truyen_Pointer(&b)
	fmt.Println(b)
}

func Truyen_Pointer(pointer *int) {
	*pointer = 696969
}
