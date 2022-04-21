package main

import "fmt"

func main() {

	// Declaration Of  an array

	var arr [5]int = [5]int{10, 20, 30}

	arr2 := [5]int{1, 2, 3, 4, 5}

	arr2[0] = 10

	arr2[2] = 32

	fmt.Printf("%v\n", arr2[0])

	fmt.Println(arr)

	fmt.Println(arr2)

}
