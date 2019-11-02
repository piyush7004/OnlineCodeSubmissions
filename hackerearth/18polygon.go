package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	var m, count int
	var s string
	arr := []int{}
	var err error
	strarr := []string{}
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		count++
		if count == 1 {
			continue
		}
		if count == 2 {
			strarr = strings.Fields(string(bytes))
			for _, s = range strarr {
				m, _ = strconv.Atoi(s)
				arr = append(arr, m)
			}
			checkPolygon(arr)
			arr = nil
			count = 0
		}
	}
}

func checkPolygon(arr []int) {
	sort.Ints(arr)
	var sum, max int
	for _, v := range arr {
		sum = sum + v
		max = v
	}
	sum = sum - max
	if max >= sum {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
