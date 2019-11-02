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
			fmt.Println(getSecondLargest(arr))
			arr = nil
		}
		count++
	}

}

func getSecondLargest(arr []int) int {
	var max, max2 int
	for _, v := range arr {
		if v > max {
			max2 = max
			max = v
		} else if v > max2 {
			max2 = v
		}
	}
	return max2
}
