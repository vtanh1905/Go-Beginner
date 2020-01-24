package main

import "fmt"

// 1. Khai variadic function
func Xuat(list ...int) {
	fmt.Println(list)
}

// 3. change slice
func change_slice(list ...int) {
	list[0] = 69
}

func main() {
	Xuat(1, 2, 3, 4, 5)

	// 2. pass 1 slice vao 1 variadic function
	slice := []int{1, 2, 3, 4, 5, 6}
	Xuat(slice...)

	change_slice(slice...)
	fmt.Println(slice)
	/*
		Giải thích
		Khi ta truyền như sau Xuat(1, 2, 3, 4, 5)
		=> Nó sẽ lấy 1 slice để hứng giá trị từ 1 -> 5

		Tương tự ở 2
			slice := []int{1, 2, 3, 4, 5, 6}
			Xuat(slice...)
		=> Khi nào 1 temp_slice = slice
		Đồng nghĩ với viec 2 slice có chung 1 refernce(đỉa chỉ)
	*/
}
