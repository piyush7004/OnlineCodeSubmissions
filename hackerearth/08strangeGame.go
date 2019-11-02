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
	var count, cost, m, v, max, total int
	sarr := []string{}
	arrA := []int{}
	arrB := []int{}
	var err error
	for {
		val, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		count++
		sarr = strings.Fields(string(val))
		//get cost
		if count == 1 {
			cost, _ = strconv.Atoi(sarr[1])
		}
		//get arrA
		if count == 2 {
			for _, s := range sarr {
				m, _ = strconv.Atoi(s)
				arrA = append(arrA, m)
			}
		}
		//get arrB
		if count == 3 {
			for _, s := range sarr {
				m, _ = strconv.Atoi(s)
				arrB = append(arrB, m)
			}
			max = findMax(arrB)
			for _, v = range arrA {
				if v >= max {
					continue
				}
				total = total + ((max + 1 - v) * cost)
			}
			fmt.Println(total)
			total = 0
			arrA = nil
			arrB = nil
			count = 0
		}
	}

}

func findMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
