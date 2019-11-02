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
	var sum, m, i int
	strarr := []string{}
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		strarr = processAdd(strarr, string(bytes))
	}
	for i = 0; i < len(strarr); i++ {
		m, _ = strconv.Atoi(strarr[i])
		sum = sum + m
	}
	fmt.Println(sum)
}

func processAdd(strarr []string, s string) []string {
	if s == "0" {
		return strarr[:len(strarr)-1]
	}
	return append(strarr, s)

}
