package main

import "fmt"

func main() {
	var sad string = "sad"
	var a, b, c = 0, 1, 2

	const status int = 202
	var tr = true

	fmt.Println(status, a, b, c, sad, tr)

	var hello = "Hello"

	fmt.Println(hello[1])

	var convHello = []rune(hello)
	convHello[1] = 'E'

	fmt.Println(string(convHello))

	dynamic := [...]int{3: 3}

	arr := dynamic
	arr[2] = 2

	arr2 := &dynamic
	arr2[2] = 1

	fmt.Println(dynamic)
	fmt.Println(arr)
	fmt.Println(arr2)
}
