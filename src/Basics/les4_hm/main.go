package main

import "fmt"

func main() {
	bookmate := map[string]map[string][]string{
		"Petr Iv": {
			"Voina i mir": {"sad", "sad"},
			"Sladkii mir": {"Rus", "2019"},
		},
		"Evgenii Nek": {
			"80/20": {"China", "2017", "2018"},
		},
	}
	// Кол-во читателей:
	fmt.Println("Reader value:", len(bookmate))
	// Кол-во книг:
	sum := 0
	for _, value := range bookmate {
		sum = sum + len(value)
	}
	fmt.Println("Total books:", sum)

	for key, value := range bookmate {
		fmt.Printf("Reader: %s, own books: %d\n", key, len(value))
	}
}
