package main

import "fmt"

// Khai báo func trong Go
// + func func_name (params) return_type { // code }

// Hàm không trả về
func sayHello() {
	fmt.Println("Hello World")
}

func sayMyName(yourName string) {
	fmt.Println("Hello", yourName)
}

func sum(a int, b int) int {
	return a + b
}

// + Multiple return values
func countArea(width int, height int) (int, int, int) {
	return width, height, width * height
}

// + Named return values (Chú thich cách bien trả về bởi vi VD tren ta kh bik 3 biến nó trả về là gi)
func isSquare(w int, h int) (width int, height int, checkSquare bool) {
	// width = w
	// height = h
	return w, h, w == h
}

func main() {
	sayHello()
	sayMyName("Tuấn Anh")
	fmt.Println("Tổng =", sum(1, 2))
	width, height, area := countArea(100, 200)
	fmt.Printf("%d %d %d\n", width, height, area)
	w, h, checkSquare := isSquare(100, 100)
	fmt.Printf("%d %d %t\n", w, h, checkSquare)
}
