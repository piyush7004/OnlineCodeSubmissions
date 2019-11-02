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
	arr := []int32{}
	for _, s := range strarr {
		m, _ := strconv.Atoi(s)
		arr = append(arr, int32(m))
	}
	res := fairRations(arr)
	if res == -1 {
		fmt.Println("NO")
	} else {
		fmt.Println(res)
	}
}

func fairRations(arr []int32) int32 {
	var sum, count, v int32
	for _, v = range arr {
		sum = sum + v
		if sum%2 == 1 {
			count = count + 2
		}
	}
	if (sum % 2) == 1 {
		return -1
	} else {
		return count
	}
}
