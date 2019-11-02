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
	nMap := make(map[int]int)
	var m, v, count, total int
	strarr := []string{}
	var err error
	s := ""
	var found bool
	for {
		val, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		strarr = strings.Fields(string(val))
		count++
		if count == 1 {
			total, _ = strconv.Atoi(strarr[0])
			for total > 0 {
				nMap[total] = 0
				total--
			}
		}

		if count == 2 {
			for v, s = range strarr {
				if v == 0 {
					continue
				}
				m, _ = strconv.Atoi(s)
				nMap[m]++
			}
		}
		if count == 3 {
			for v, s = range strarr {
				if v == 0 {
					continue
				}
				m, _ = strconv.Atoi(s)
				nMap[m]++
			}

			for _, v = range nMap {
				if v == 0 {
					found = true
					fmt.Println("NO")
					break
				}

			}
			if !found {
				fmt.Println("YES")
			}
			count = 0
			for v, _ = range nMap {
				delete(nMap, v)
			}
			found = false
		}

	}
}
