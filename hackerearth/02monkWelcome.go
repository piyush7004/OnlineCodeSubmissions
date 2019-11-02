package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	strarr := []string{}
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		strarr = append(strarr, string(line))
	}
	arrA := strings.Split(strarr[1], " ")
	arrB := strings.Split(strarr[2], " ")
	arrC := []string{}

	for i := 0; i < len(arrA); i++ {
		n1, _ := strconv.Atoi(arrA[i])
		n2, _ := strconv.Atoi(arrB[i])
		arrC = append(arrC, strconv.Itoa(n1+n2))
	}
	for _, v := range arrC {
		fmt.Printf("%s ", v)
	}

}
