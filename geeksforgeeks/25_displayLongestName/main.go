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
	reader := bufio.NewReaderSize(file, 128*128)
	bytes, _, _ := reader.ReadLine()
	bytes, _, _ = reader.ReadLine()
	strarr := []string{}
	var s, maxs string
	var max int
	var err error
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		strarr = append(strarr, string(bytes))
	}
	for _, s = range strarr {
		if len(s) > max {
			max = len(s)
			maxs = s
		}
	}
	fmt.Println(maxs)
}
