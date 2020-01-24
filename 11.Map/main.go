package main

import "fmt"

func main() {
	// Khai báo map
	map1 := make(map[string]int)
	// fmt.Println(map1)
	if map1 == nil {
		fmt.Println("map1 == nil")
	} else {
		fmt.Println(map1)
	}

	var map2 map[string]int
	// fmt.Println(map2)
	if map2 == nil {
		fmt.Println("map1 == nil")
	} else {
		fmt.Println(map1)
	}

	// Khai bao voi gia tri khoi tao map
	map3 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(map3)

	// Them 1 phan tu vao map
	map3["key3"] = 69
	fmt.Println(map3)

	// delete 1 phan tu trong map
	delete(map3, "key3")
	fmt.Println(map3)

	// length của map : len(map)
	fmt.Println(len(map3))

	// map là 1 reference type
	map4 := map3
	map4["key3"] = 69
	fmt.Println(map3)
	fmt.Println(map4)

	// Truy cap 1 phan tu trong map
	fmt.Println(map4["key3"])

	fmt.Println(map4["key4"]) // neu key khong tồn tài thi sẽ tra ra inital value là 0

	// Kiem tra key đó co tồn tại trong map
	value, checkExist := map4["key4"]
	if checkExist {
		fmt.Println(value)
	} else {
		fmt.Println("Khong ton tai key4 trong map4")
	}

	// Note map kh có toàn tư ==
	// Err : invalid operation: map3 == map4 (map can only be compared to nil) go
	// if map3 == map4 {

	// }
}
