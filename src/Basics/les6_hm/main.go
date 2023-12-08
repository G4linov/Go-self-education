package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func logTime() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Execution time %v.", time.Since(start))
	}
}

func main() {
	defer logTime()()
	outFile, _ := os.Create("les6out.txt")
	inFile, _ := os.Open("les6in.txt")
	byteval := 0
	strval := 0

	defer func() {
		fmt.Printf("Strings: %d.\nBytes: %d.\n", strval, byteval)
		outFile.Close()
	}()
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		writer.WriteString(string(line))
		strval++
		byteval = byteval + len([]byte(line))

	}
	writer.Flush()
}
