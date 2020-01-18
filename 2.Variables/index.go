package main

import "fmt"

func main() {
	// var number int = 10
	var number int
	number = 10
	fmt.Println(number)

	//Type inference
	var number2 = 10 // Golang se tự hiểu kiểu dữ liệu
	fmt.Println(number2)

	// -  Khai báo nhieu biến trong golang

	// + Cùng kiểu dư liệu
	var a, b int = 9, 8
	fmt.Printf("%d, %d ", a, b)

	// Khac kiểu dư lieu
	var (
		c int    = 69
		d string = "Kem Map"
	)
	fmt.Printf("%d, %s ", c, d)

	// + Cách khac dùng type inference
	var e, f = 69, "Kem Ngu Nhieu Qua"
	fmt.Printf("%d, %s", e, f)

	// - Mốt cách khai báo phô biến trong golang
	name := "Tuan Anh"
	fmt.Printf(name)
}
