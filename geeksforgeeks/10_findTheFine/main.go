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
	reader := bufio.NewReaderSize(file, 256*256)
	bytes, _, _ := reader.ReadLine()
	var m, date, count int
	var s string
	fines := []int{}
	numbers := []int{}
	var err error
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		if count%3 == 1 {
			for _, s = range strings.Fields(string(bytes)) {
				m, _ = strconv.Atoi(s)
				numbers = append(numbers, m)
			}
		} else if count%3 == 0 {
			strarr := strings.Fields(string(bytes))
			s = strarr[1]
			date, _ = strconv.Atoi(s)

		} else {
			for _, s = range strings.Fields(string(bytes)) {
				m, _ = strconv.Atoi(s)
				fines = append(fines, m)
			}
			fmt.Println(findFine(date, numbers, fines))
			fines = nil
			numbers = nil
			date = 0
		}
		count++
	}
}

func findFine(date int, numbers, fines []int) int {
	isEven := false
	if date%2 != 0 {
		isEven = true
	}
	sum := 0
	for i, v := range numbers {
		if (v%2 == 0) == isEven {
			sum += fines[i]
		}
	}
	return sum
}
