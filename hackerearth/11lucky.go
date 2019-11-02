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
	reader := bufio.NewReaderSize(file, 10*1024*1024)
	val, _, _ := reader.ReadLine()
	var count, m, v, min, total int
	strarr := []string{}
	s := ""
	arr := []int{}
	var err error
	for {
		val, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		count++
		if count == 1 {
			continue
		}
		if count == 2 {
			strarr = strings.Fields(string(val))
			for _, s = range strarr {
				m, _ = strconv.Atoi(s)
				arr = append(arr, m)
			}
			min = findMin(arr)
			for _, v = range arr {
				if v == min {
					total++
				}
			}
			if total%2 == 0 {
				fmt.Println("Unlucky")
			} else {
				fmt.Println("Lucky")
			}
			count = 0
			arr = nil
			total = 0
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
