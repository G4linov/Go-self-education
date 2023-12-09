package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func writeString(writer bufio.Writer, line []string, i int) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			switch panicErr {
			case "empty":
				err = fmt.Errorf("Empty field on string %d.", i)
			default:
				panic("critical")
			}
		}
	}()
	for _, v := range line {
		if len(strings.TrimSpace(v)) == 0 {
			panic("empty")
		}
	}
	outline := "Row: " + strconv.Itoa(i) + "\nName: " + line[0] + "\nAdress: " + line[1] + "\nCity: " + line[2] + "\n\n\n"
	writer.WriteString(outline)
	writer.Flush()
	return nil
}

func main() {
	outFile, _ := os.Create("les7out.txt")
	inFile, _ := os.Open("les7in.txt")
	defer outFile.Close()
	defer inFile.Close()

	i := 0

	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)

	for {
		i = i + 1
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		newLine := strings.Split(line, "|")
		err = writeString(*writer, newLine, i)
		if err != nil {
			fmt.Println(err)
		}
	}
}
