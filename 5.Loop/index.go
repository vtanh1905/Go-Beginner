package main

import "fmt"

func main() {
	// Vòng lặp

	// goto
	// Loop1:
	// 	number := 10
	// 	fmt.Println(number)

	// 	goto Loop1

	// for

	for index := 0; index < 10; index++ {
		fmt.Println(index)
		// break giống mọi ngôn ngữ khác
		// continue giống mọi ngôn ngữ khác
	}

	for n, m := 0, 2; n < 10 && m < 10; n, m = n+1, m+1 {
		fmt.Println(n + m)
		// break giống mọi ngôn ngữ khác
		// continue giống mọi ngôn ngữ khác
	}

	// white
	// Trong Go không có white mà ta có thể mô phỏng nó ở vòng lặp for
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Infinite loop
	for {
		fmt.Println("Hello World")
	}
}
