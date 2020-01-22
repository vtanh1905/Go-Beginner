package main

import "fmt"

func main() {
	// Khai bao slice
	var slice []int
	fmt.Println(slice)

	// Khai bao vs khoi tao
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	// Tao slice từ 1 array
	var arr = [...]int{1, 2, 3}
	slice2 := arr[1:2]
	//slice2 := arr[:2]
	//slice2 := arr[1:]
	fmt.Println(slice2)

	// Slice => reference Type , Array => Value Type
	slice3 := []int{1, 2, 3}
	slice4 := slice3
	slice4[1] = 69
	fmt.Println(slice3)
	fmt.Println(slice4)

	// length va capacity của slice
	// + length : số phân tử của slice
	// + capacity : so phần tử của underlying array bat dau vi tri start khi ma slice duoc tao
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7}
	slice5 := arr2[2:3]
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	// cac ham làm viec voi slice make copy append

	// make(type, len, (cap)) : tạo ra 1 slice với len va cap mình thik
	slice6 := make([]int, 3, 5)
	fmt.Println(slice6)
	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))

	// append(param1, param2) : để thêm 1 phần từ vào 1 phần tử (thêm 1 phan tu vao slice hoặc 1 slice vào slice)
	slice7 := []int{1, 2, 3}
	fmt.Println(slice7)
	slice7 = append(slice7, 69)
	fmt.Println(slice7)

	// copy(param1, param2): trả ve số phần tử copy dc, copy phân từ param2 wa param1 (neu len(param1) < len(param2) thì chỉ copy đúng len(param1)
	src := []int{1, 2, 3, 4, 5, 6, 7, 8}
	dest := make([]int, 2) // 1 cách khac tao slice
	copy(dest, src)
	fmt.Println(dest)
}
