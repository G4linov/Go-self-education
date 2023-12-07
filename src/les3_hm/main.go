package main

import "fmt"

func main() {
	/*
		Создать слайс дней недели
		Отделить рабочие в отдельный слайс, в исходном удалить рабочие
	*/
	week := []string{"pn", "vt", "sr", "cht", "pt", "sb", "vs"}
	work := make([]string, 5, 5)
	copy(work, week)
	week = append(week[:0], week[5:]...)
	fmt.Println(work, week)
	// Объеденить 2 слайса в один
	allweek := append(work, week...)
	fmt.Println(allweek)
}
