package main

import "fmt"

func someFunc() {
	defer fmt.Println("someFunc: defer")

	panic("something get wrong")
}

func newSomeFunc() (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			switch panicErr {
			case "regular error":
				err = fmt.Errorf("application error")
			default:
				panic("critical")
			}
		}
	}()

	panic("regular error")
}

func main() {
	//defer fmt.Println("main: defer")
	//someFunc()
	//fmt.Println("main after call someFunc")

	err := newSomeFunc()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("after func")
}
