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
	val, _, _ := reader.ReadLine()
	count := 0
	var err error
	max := 0
	for {
		arr := []int{}
		m := 0
		val, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		count++
		sarr := strings.Fields(string(val))
		//get the max
		if count == 1 {
			max, _ = strconv.Atoi(sarr[1])
		} else {
			//get the array
			for _, s := range sarr {
				m, _ = strconv.Atoi(s)
				arr = append(arr, m)
			}
		}
		if count == 2 {
			count = 0
			min := findMin(arr)
			fmt.Println(min)
			if max > min {
				fmt.Println(max - min)
			} else {
				fmt.Println("0")
			}
			arr = nil
		}

	}
}

func findMin(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
