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
	val, _, _ = reader.ReadLine()
	strarr := strings.Fields(string(val))
	res := getIndices(strarr)
	fmt.Println(len(res))
	strarr = nil
	arr := []int{}
	var err error
	var l, r, v int
	val, _, _ = reader.ReadLine()
	for {
		val, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		strarr = strings.Fields(string(val))
		l, _ = strconv.Atoi(strarr[0])
		r, _ = strconv.Atoi(strarr[1])
		for _, v = range res {
			if v < l || v > r {
				continue
			}
			if v >= l && v <= r {
				arr = append(arr, v)
			}
		}
		fmt.Println(len(arr))
		arr = nil

	}
}

func getIndices(strarr []string) []int {
	// var ul, ur, m, total int
	// lstr := []string{}
	// rstr := []string{}
	// threestr := []string{}
	s := ""
	out := []int{}
	arr := []int{}
	var m, sum int
	for _, s = range strarr {
		m, _ = strconv.Atoi(s)
		arr = append(arr, m)
		sum = sum + m
	}
	exr := arr[len(arr)-1]
	if (arr[0] == 0) && (exr == 0) {
		out = append(out, 1)
	}
	if (arr[len(arr)-2] == 0) && ((sum-arr[len(arr)-1])%3 == 0) {
		out = append(out, len(arr))
	}
	for i := 1; i < len(arr); i++ {
		// m, _ = strconv.Atoi(strarr[i-1])
		m = arr[i-1]

		if m != (10 - exr) {
			continue
		}
		//if here that means it is divisible by 2 and 5
		if ((sum - arr[i]) % 3) != 0 {
			continue
		}
		//if here means it is divisible by 3 too
		out = append(out, i+1)
	}

	// for i := 0; i < len(strarr); i++ {
	// 	ur = 0
	// 	ul = 0
	// 	lstr = nil
	// 	total = 0
	// 	rstr = nil
	// 	threestr = nil
	// 	for j := 0; j < i; j++ {
	// 		lstr = append(lstr, strarr[j])
	// 	}
	// 	for k := i + 1; k < len(strarr); k++ {
	// 		rstr = append(rstr, strarr[k])
	// 		// n, _ = strconv.Atoi(strarr[k])
	// 		// r = (r * 10) + n
	// 	}
	// 	threestr = append(lstr, rstr...)
	// 	//check divisibility by 3
	// 	for _, s = range threestr {
	// 		m, _ = strconv.Atoi(s)
	// 		total = total + m
	// 	}
	// 	if (total % 3) != 0 {
	// 		continue
	// 	}
	// 	if len(lstr) == 0 {
	// 		ul = 0
	// 	} else {
	// 		ul, _ = strconv.Atoi(lstr[len(lstr)-1])
	// 	}
	// 	if len(rstr) == 0 {
	// 		ur = 0
	// 	} else {
	// 		ur, _ = strconv.Atoi(rstr[len(rstr)-1])
	// 	}
	// 	//check divisibility by 2
	// 	if ((ur + ul) % 2) != 0 {
	// 		continue
	// 	}

	// 	//check divisibility by 5
	// 	if ((ur + ul) % 5) != 0 {
	// 		continue
	// 	}

	// 	//if you are here then add this index
	// 	out = append(out, i+1)
	// }
	return out
}
