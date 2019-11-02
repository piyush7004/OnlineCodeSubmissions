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
	val, _, _ := reader.ReadLine()
	val, _, _ = reader.ReadLine()
	sarr := strings.Fields(string(val))
	arr := []int64{}
	m := 0
	for _, s := range sarr {
		m, _ = strconv.Atoi(s)
		arr = append(arr, int64(m))
	}
	count := 1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			count++
		}
	}
	fmt.Println(count)
}
