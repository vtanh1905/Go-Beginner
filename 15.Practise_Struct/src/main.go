/*
	Quản lý sách
	Sách gồm các thông tin : id, name, author, kind, price, createAt
	1. Thêm sách
	2. Xem danh sách
	3. Xóa sách
	4. Sữa sách
	5. Tìm kiếm
*/

package main

import (
	book "book"
	"fmt"
)

func main() {
	book1 := book.Book{
		id:   1,
		name: "Toán 1",
	}

	fmt.Println(book1)
}
