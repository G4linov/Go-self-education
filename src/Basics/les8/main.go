package main

import "fmt"

func main() {
	_, err := checknum(11)
	if err != nil {
		fmt.Println("invalid num:", err)
	}
}

func checknum(x int) (valid bool, err error) {
	if x > 0 && x > 10 {
		return false, fmt.Errorf("out of positive range: %v", x)
	}
	if x < 0 && x < -10 {
		return false, fmt.Errorf("out of negative range: %v", x)
	}

	return true, nil
}
