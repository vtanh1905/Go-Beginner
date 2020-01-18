package main

import "fmt"

func main() {
	// Các loại khai bào câu điều khiên
	number := 10

	// If else
	if number == 10 {
		fmt.Println("number = 10")
	} else {
		fmt.Println("number != 10")
	}
	// =====================

	// If else else if
	if number > 10 {
		fmt.Println("number > 10")
	} else if number < 10 {
		fmt.Println("number < 10")
	} else {
		fmt.Println("number = 10")
	}
	// =====================

	// Switch case
	// + break
	// + fallthrough
	// + fallthrough & break

	switch number {
	case 10:
		fmt.Println("number = 10")
		// break trong Golang nó sẽ tự break mà ta không cần phải khai báo
	case 20:
		fmt.Println("number = 20")
	case 30:
		fmt.Println("number = 20")
	default:
		fmt.Println("unknow")
	}

	switch number {
	case 10:
		fmt.Println("number = 10")
		fallthrough // khi có này thi nó sẽ bỏ break và tự xuống case dưới (chỉ 1 case thoi nhá)
	case 20:
		// if number == 10 {
		// 	break TH ta cũng có thể sai break để câu điều kiện
		// }
		fmt.Println("number = 20")
	case 30:
		fmt.Println("number = 20")
	default:
		fmt.Println("unknow")
	}
}
