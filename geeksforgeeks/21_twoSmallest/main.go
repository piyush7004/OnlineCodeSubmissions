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
	var s string
	var err error
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
			fmt.Println(getTwoSmallest(arr))
			arr = nil
		}
		count++
	}

}

func getTwoSmallest(arr []int) (int, int) {
	var min, min2 int
	max := ^uint(0)
	min = int(max >> 1)
	min2 = int(max >> 1)
	for _, v := range arr {
		if v < min {
			min2 = min
			min = v
		} else if v < min2 && v != min {
			min2 = v
		}
	}
	if min2 == int(max>>1) {
		return -1, -1
	}
	return min, min2
}
