package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	var err error
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(processBrackets(bytes))
	}
}

func processBrackets(bytes []byte) string {
	out := []byte{}
	for _, b := range bytes {
		if b == '{' || b == '[' || b == '(' {
			out = append(out, b)
			continue
		}
		if len(out) == 0 {
			return "NO"
		}
		if b == '}' && out[len(out)-1] == '{' {
			out = out[:len(out)-1]
			continue
		}
		if b == ')' && out[len(out)-1] == '(' {
			out = out[:len(out)-1]
			continue
		}
		if b == ']' && out[len(out)-1] == '[' {
			out = out[:len(out)-1]
			continue
		}
	}
	if len(out) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}
