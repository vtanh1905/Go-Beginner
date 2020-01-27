package main

import "fmt"

// 1.Khai bao struct
type Student struct {
	id   string
	name string
}

func main() {
	// 2. Khoi tao(1) tường mình (khuyên dùng) (có thể khai báo thiếu field)
	student1 := Student{
		id:   "1",
		name: "A",
	}
	fmt.Println(student1)

	// 3. Truy xuat vào biên trong struct
	fmt.Println("id", student1.id)
	fmt.Println("name", student1.name)

	// 4. khoi tao(2) ngắn gọn (phải khai báo hêt tat ca các field khong la no báo lỗi)
	student2 := Student{"2", "B"}
	fmt.Println(student2)

	// 5. Khoi tao(3) dài dòng
	var student3 Student = struct {
		id   string
		name string
	}{
		id:   "3",
		name: "C",
	}
	fmt.Println(student3)

	// 6. anonymous struct (struct vo danh)
	var student4 = struct {
		class string
		grade int
	}{
		class: "Math",
		grade: 9,
	}
	fmt.Println(student4)

	// 7.pointer tro toi struct
	var pointer1 *Student
	pointer1 = &student3
	fmt.Println("student3", student3)
	fmt.Println("pointer1", *pointer1)
	// (*pointer1).id = "69"
	pointer1.id = "69" // không cần * cũng không sao golang tự hiểu
	fmt.Println("student3", student3)
	fmt.Println("pointer1", *pointer1)

	pointer2 := &Student{
		id:   "999",
		name: "ABC",
	}
	fmt.Println(pointer2)

	// 8. anonymous fields
	type NoName struct {
		string
		int
	}
	temp := NoName{"123", 69}
	fmt.Println(temp)

	// 9. Struct lồng struct - nested stuct
	// Khai bao vs khoi tao struct long
	type Class struct {
		id       string
		students []Student
	}
	class1 := Class{
		id: "1",
	}
	fmt.Println(class1)
	class1.students = append(class1.students, Student{id: "69", name: "123"})
	class1.students = append(class1.students, Student{id: "96", name: "123"})
	fmt.Println(class1)
	fmt.Println(class1.students[1].id)

	// 10 compare 2 struct
	// Note có the compare được nếu các field bên trong có kieu dữ lieu co the so sanh duoc (VD: map slice array)
	student5 := Student{
		id:   "9",
		name: "A",
	}

	student6 := Student{
		id:   "9",
		name: "A",
	}
	if student5 == student6 {
		fmt.Println("student5 == student6")
	} else {
		fmt.Println("student5 != student6")
	}

	// 11. zero value
	// Ta chi khai báo thì nó sẽ lấy zero value tương ứng với kieu dữ liệu của flied đó
	student7 := Student{
		id: "69",
	}
	fmt.Println(student7)
}
