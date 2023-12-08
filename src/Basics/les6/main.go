package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		path := "f.txt"
		getFile(path)
		log()
	*/
	i := 1
	defer func() {
		prt(i)
	}()

	i++
	fmt.Printf("main func execution: %d\n", i)
}

func getFile(path string) {
	f, _ := os.Open(path)
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

func log() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func prt(i int) {
	fmt.Printf("Deffered func call: %d\n", i)
}
