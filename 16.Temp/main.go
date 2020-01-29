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
	"fmt"
	"time"
)

type Book struct {
	id       int
	name     string
	author   string
	kind     string
	price    float32
	createAt string
}

type ListBook struct {
	items []Book
}

func (l ListBook) getLength() int {
	return len(l.items)
}

func (l *ListBook) addBook(name string, author string, price float32) {
	l.items = append(l.items, Book{
		id:       l.getLength(),
		name:     "Toán 1",
		author:   "Nguyễn Văn A",
		price:    15.500,
		createAt: fmt.Sprintf("%d", (time.Now().UnixNano())/1000000),
	})
}

func (l ListBook) showBooks() {
	for i := 0; i < l.getLength(); i++ {
		fmt.Println("======================== Sách Thứ", i+1, "========================")
		fmt.Println(l.items[i])
		fmt.Println("============================================================")
	}
}

func (l *ListBook) removeBook(id int) {
	indexRemove := -1

	for i := 0; i < l.getLength(); i++ {
		if l.items[i].id == id {
			indexRemove = i
			break
		}
	}
	if indexRemove != -1 {
		l.items = append(l.items[:indexRemove], l.items[indexRemove+1:]...)
	}
}

func (l *ListBook) editBook(id int, name string, author string, price float32) {
	indexEdit := -1

	for i := 0; i < l.getLength(); i++ {
		if l.items[i].id == id {
			indexEdit = i
			break
		}
	}
	if indexEdit != -1 {
		// l.items[indexEdit] = data
		if name != "" {
			l.items[indexEdit].name = name
		}
		if author != "" {
			l.items[indexEdit].author = author
		}
		if price != 0 {
			l.items[indexEdit].price = price
		}
	}
}

func (l ListBook) findBook(id int) Book {
	index := -1

	for i := 0; i < l.getLength(); i++ {
		if l.items[i].id == id {
			index = i
			break
		}
	}
	if index != -1 {
		return l.items[index]
	}
	return Book{}
}

func main() {
	listBook := ListBook{}

	fmt.Println(listBook)

	listBook.addBook("Toán 1", "Nguyễn Văn A", 15.500)
	listBook.addBook("Toán 1", "Nguyễn Văn A", 15.500)
	listBook.addBook("Toán 1", "Nguyễn Văn A", 15.500)
	listBook.addBook("Toán 1", "Nguyễn Văn A", 15.500)
	listBook.addBook("Toán 1", "Nguyễn Văn A", 15.500)
	listBook.addBook("Toán 1", "Nguyễn Văn A", 15.500)
	listBook.addBook("Toán 1", "Nguyễn Văn A", 15.500)

	// listBook.removeBook(12)
	// listBook.removeBook(5)
	// listBook.showBooks()

	listBook.editBook(6, "Toán 2", "", 0)
	listBook.showBooks()
	fmt.Println(listBook.findBook(78))
}
