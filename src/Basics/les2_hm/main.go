package main

import "fmt"

func main() {
	A := new(int)
	B := 9999
	A = &B
	fmt.Println(A)
}
