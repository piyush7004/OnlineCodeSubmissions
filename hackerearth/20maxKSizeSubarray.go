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
	k, _ := strconv.Atoi(strarr[1])
	bytes, _, _ = reader.ReadLine()
	var m, i int
	arr := []int{}
	// subarr := []int{}
	strarr = strings.Fields(string(bytes))
	for _, s := range strarr {
		m, _ = strconv.Atoi(s)
		arr = append(arr, m)
	}
	// max = findMax(arr[:k])
	// fmt.Printf("%d ", max)
	for i = 1; i <= len(arr)-k; i++ {
		fmt.Printf("%d ", findMax(arr[i:i+k]))

		// fmt.Println("test", arr[i], max)
		// fmt.Println(i, "-", arr[i])
		// if arr[i] > max {
		// 	max = arr[i]
		// } else if arr[i-1] == max {
		// 	fmt.Println(arr[i : i+k])
		// 	max = findMax(arr[i : i+k])
		// } else {
		// 	// } else if arr[i-1] == max {
		// 	fmt.Printf("%d ", max)
		// 	subarr = arr[i : i+k]
		// 	fmt.Println(subarr)
		// 	max = findMax(subarr)
		// } else {
		// 	continue
		// }
		// fmt.Printf("%d ", max)
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
