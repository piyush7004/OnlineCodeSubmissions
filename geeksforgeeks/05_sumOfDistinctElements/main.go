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
	reader := bufio.NewReaderSize(file, 128*128)
	bytes, _, _ := reader.ReadLine()
	var m, count int
	arr := []int{}
	var err error
	var s string
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		if count%2 != 0 {
			for _, s = range strings.Fields(string(bytes)) {
				m, _ = strconv.Atoi(s)
				arr = append(arr, m)
			}
			fmt.Println(findSum(arr))
			arr = nil
		}
		count++
	}
}

func findSum(arr []int) int {
	nmap := make(map[int]int)
	var sum int
	for i := range arr {
		nmap[arr[i]]++
	}

	for i := range nmap {
		sum += i
	}
	return sum
}
