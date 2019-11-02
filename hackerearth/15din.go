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
	strarr := []string{}
	arr := []int{}
	nMap := make(map[int]int)
	s := ""
	var err error
	var count, m, v int

	for {
		val, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		count++
		strarr = strings.Fields(string(val))
		if count == 1 {
			continue
		}
		if count == 2 {
			for _, s = range strarr {
				m, _ = strconv.Atoi(s)
				arr = append(arr, m)
			}
			for _, v = range arr {
				nMap[v]++
			}
			for m, v = range nMap {
				if v == 1 {
					fmt.Println(m)
				}
				delete(nMap, m)
			}
			arr = nil
			count = 0
		}
	}

}
