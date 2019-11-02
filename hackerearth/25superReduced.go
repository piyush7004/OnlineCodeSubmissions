package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	fmt.Println(processString(bytes))
}

func processString(bytes []byte) string {
	out := []byte{}
	for _, b := range bytes {
		if len(out) == 0 {
			out = append(out, b)
			continue
		}
		if b == out[len(out)-1] {
			out = out[:len(out)-1]
			continue
		}
		out = append(out, b)
	}
	if len(out) == 0 {
		return "Empty String"
	}
	return string(out)
}
