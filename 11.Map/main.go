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

	/*
		Note: null khác nil thế nào?
		Câu trả lời dài: Cả 2 đều xuất phát từ tiếng Latinh. null là nullus, nil là nihil.
		Tiếng Việt dịch là "không có gì cả". Tiếng Anh dịch là "no" và "nothing".

		Câu trả lời ngắn: nil dùng để chỉ việc biến chưa có giá trị, (chưa được khởi tạo).

		Câu trả lời thêm: Xem mã nguồn của các implementation của các ngôn ngữ động,
		để ý sẽ thấy các biến có cấu trúc hơi phức tạp trở lên đều được gói bằng con trỏ,
		chỉ khác là con trỏ này được tự động quản lí.
		Do đó khái niệm con trỏ thực ra chỉ bị dấu đi để lập trình viên đỡ điên đầu, tránh mắc lỗi liên quan đến bộ nhớ.
	*/

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
