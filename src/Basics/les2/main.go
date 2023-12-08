package main

import "fmt"

func main() {
	var x int
	pntr := &x
	fmt.Println(pntr)

	x = 64
	fmt.Println(*pntr)

	*pntr = 31
	fmt.Println(x, *pntr)

	pntr2 := new(int)
	fmt.Println(pntr2)
}
