package main

import "fmt"

func contains(x string, a []string) (issub bool) {
	issub = false
	for _, v := range a {
		if v == x {
			issub = true
		}
	}
	return
}

func getMax(p ...int) (max int) {
	max = 0
	for _, v := range p {
		if v > max {
			max = v
		}
	}
	return
}

func main() {
	m := []string{"sad", "bois"}
	check := "sa"
	fmt.Println(contains(check, m))
	fmt.Println(getMax(1, 2, 3, 4, 5))
}
