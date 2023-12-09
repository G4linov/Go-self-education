package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inFile, err := os.Open("basLes8in.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := inFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	reader := bufio.NewReader(inFile)

	i := 0
	for {
		i++
		_, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("End of file. io.EOF")
				break
			} else {
				fmt.Println(err)
				break
			}
		}
	}
	fmt.Printf("Strings value: %d\n", i)
}
