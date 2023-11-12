package main

import "fmt"

func main3() {
	fmt.Println("Hello, World!")

	// List
	arr1 := []int{1, 2, 3, 4}
	fmt.Println("arr1 = ", arr1)

	arr2 := make([]int, 4)
	fmt.Println("arr2 = ", arr2)
	arr2[0] = 30
	fmt.Println("arr2 = ", arr2)

	txt := "Today is Sunday"
	fmt.Println(txt[0:5])
	fmt.Println(arr2[0:1])

	fmt.Println(len(txt))
	fmt.Println(len(arr2))
}