package main

import "fmt"

func main() {
	var s []int
	fmt.Println(len(s))
	fmt.Println(cap(s))

	news := []string{"e", "l", "e", "m", "e", "n", "t"}
	news[0] = "E"
	fmt.Println(news)

	makes := make([]int, 4, 6)
	fmt.Println(makes, len(makes), cap(makes))

	news1 := news[:4]
	news2 := news[1:6]
	fmt.Println(news1, news2)
	news1[1] = "L"
	fmt.Println(news1, news2)

	si := []int{1, 2, 3}
	sp := []int{4, 5, 6}

	si = append(si, sp...)
	fmt.Println(si)

	si = append(si, 33, 44, 55)
	fmt.Println(si)

	fmt.Println(cap(news1))
	news1 = append(news1, "one", "two", "three", "four")
	fmt.Println(news, news1, news2)
	fmt.Println(cap(news1))

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3, 10)

	qnt := copy(dst, src)
	fmt.Println(qnt, dst)

	dst = append(dst, 0, 0)

	qnt = copy(dst, src)
	fmt.Print(qnt, dst)

}
