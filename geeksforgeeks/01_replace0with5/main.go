package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
		// i, _ := strconv.Atoi(string(bytes))
		fmt.Println(rep0with5(bytes))
	}
}

func rep0with5(bs []byte) int {
	out := make([]byte, len(bs))
	for i, b := range bs {
		if b != '0' {
			out[i] = b
		} else {
			out[i] = '5'
		}
	}
	// fmt.Println(out)
	m, _ := strconv.Atoi(string(out))
	return m
}
