package main

import (
	"fmt"
)

func funcName(param int) int {
	param++
	return param
}

func mySum(p1, p2, p3 int) int {
	return p1 + p2 + p3
}

func neSum(p1, p2 int) (sum int) {
	sum = p1 + p2
	return
}

func words(w ...string) []string {
	return w
}

type wordLogger func(w ...string) (cnt int)

func main() {
	logicValue := false

	if logicValue {
		fmt.Println("Value true")
	} else {
		fmt.Println("Value false")
	}

	customer := map[string]string{
		"name":     "John",
		"lastname": "Smith",
	}

	if lastname, isExist := customer["lastname"]; isExist {
		fmt.Printf("Hello Mr. %s\n", lastname)
	}

	if _, isExist := customer["dead"]; !isExist {
		fmt.Println("empty field!")
	}

	role := "sad"

	switch role {
	case "admin":
		fmt.Println("Admin access")
	case "user":
		fmt.Println("User access")
	default:
		fmt.Println("Access denied")
	}

	role = "guest"

	switch role {
	case "admin", "user":
		fmt.Println("Access granted")
	case "guest":
		fallthrough
	default:
		fmt.Println("Access denied")
	}

	for i := 0; i <= 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	isBreak := true

	for {
		if isBreak {
			break
		}
	}
Loop1:
	for {
		fmt.Println("Loop 1")
		for {
			fmt.Println("Loop 2")
			if isBreak {
				fmt.Println("break")
				break Loop1
			}
		}
	}
	fmt.Println("After Loop")

	iterate := true
	for iterate {
		fmt.Println("Iteration")
		iterate = false
	}

	ma := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}

	mb := []int{1, 2, 3, 4}

	for k, v := range ma {
		fmt.Println(k, v)
	}
	for k := range ma {
		fmt.Println(k)
	}
	for _, v := range ma {
		fmt.Println(v)
	}

	for k, v := range mb {
		fmt.Println(k, v)
	}

	fmt.Println(funcName(3))
	fmt.Println(mySum(1, 2, 3))
	fmt.Println(neSum(1, 2))
	fmt.Println(words("sad", "bois", "sad bois"))
	myMas := words("sad", "bois", "sad bois")
	maa := []string{"sad", "bois", "sad bois"}
	for _, v := range myMas {
		fmt.Println(v)
	}
	myMas = words(maa...)
	fmt.Println(myMas)

	printWord := func(w ...string) (cnt int) {
		for _, word := range w {
			fmt.Println(word)
		}
		return len(w)
	}

	printWord(myMas...)

	func(w ...string) {
		printWord(w...)
	}("str1", "str2")

	func(printer wordLogger, w ...string) {
		fmt.Printf("Count words: %d\n", printer(w...))
	}(printWord, "str1", "str2")

	sequence := sequenceGen(4)
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

}

func sequenceGen(n int) func() int {
	i := 0
	return func() int {
		i += n
		return i
	}
}
