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
	var err error
	var s string
	var m, count int
	arr := []int{}
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
			fmt.Println(findMaxStruct(arr))
			arr = nil
		}
		count++
	}
}

func findMaxStruct(arr []int) int {
	var max, tmp int
	for i, v := range arr {
		if i%2 == 0 {
			tmp += (v * 12)
		} else {
			tmp += v
			if tmp > max {
				max = tmp
			}
			tmp = 0
		}
	}
	return max
}
