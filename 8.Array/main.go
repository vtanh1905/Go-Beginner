package main

import "fmt"

func main() {
	// Khai bao &  Gán giá trị
	// Cách 1
	var arr [2]int
	arr[0] = 69
	arr[1] = 96
	fmt.Println("arr=", arr)

	// Cách 2
	arr1 := [2]int{1, 2}
	fmt.Println("arr1=", arr1)

	// Khai báo không cần length
	arr2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("arr2", arr2)

	// Lay len
	fmt.Println("len(arr2)=", len(arr2))

	// * Array trong Golang la value type chứ không phải reference type
	arr3 := [...]int{1, 2, 3}
	temp := arr3
	temp[2] = 69
	fmt.Println("arr3 =", arr3)
	fmt.Println("temp =", temp)

	// Duyet array
	arr4 := [...]int{1, 2, 3}
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}

	// (forEach in Golang)
	for index, value := range arr4 {
		fmt.Println("Index = ", index, "Value = ", value)
	}

	// blank identifier

	for _, value := range arr4 {
		fmt.Println("Value = ", value)
	}

	// Matrix
	matrix := [3][2]int{
		{1, 2},
		{1, 2},
		{1, 2},
	}
	fmt.Println(matrix)
}
