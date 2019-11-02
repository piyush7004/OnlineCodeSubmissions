package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var n, i, max, out, v, k int
	var m, max, k, v, out int
	max = -9999
	nmap := make(map[int]int)

	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	val, _, _ := reader.ReadLine()
	val, _, _ = reader.ReadLine()
	strarr := strings.Fields(string(val))
	for _, s := range strarr {
		m, _ = strconv.Atoi(s)
		nmap[m]++
	}
	fmt.Println(nmap)
	for k, v = range nmap {
		if v > max {
			max = v
			out = k
		}
	}
	fmt.Println(out, max)

	// fmt.Scanln("%d", &n)
	// for n > 0 {
	// 	fmt.Scanf("%d", &i)
	// 	nmap[i]++
	// 	n--
	// }
	// for k, v = range nmap {
	// 	if v > max {
	// 		max = v
	// 		out = k
	// 	}
	// }
	// fmt.Println(out)
}
