package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	strarr := strings.Fields(string(bytes))
	max, _ := strconv.Atoi(strarr[1])
	bytes, _, _ = reader.ReadLine()
	strarr = strings.Fields(string(bytes))
	var m, v, score int
	arr := []int{}
	allowed := 2
	for _, s := range strarr {
		m, _ = strconv.Atoi(s)
		arr = append(arr, m)
	}
	for _, v = range arr {
		if v <= max {
			score++
		} else {
			allowed--
		}
		if allowed == 0 {
			break
		}
	}
	fmt.Println(score)
}
