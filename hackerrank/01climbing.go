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
	bytes, _, _ = reader.ReadLine()
	strarr := strings.Fields(string(bytes))
	scores := []int32{}
	alice := []int32{}
	var m int
	var s string
	for _, s = range strarr {
		m, _ = strconv.Atoi(s)
		scores = append(scores, int32(m))
	}
	bytes, _, _ = reader.ReadLine()
	bytes, _, _ = reader.ReadLine()
	strarr = strings.Fields(string(bytes))
	for _, s = range strarr {
		m, _ = strconv.Atoi(s)
		alice = append(alice, int32(m))
	}
	fmt.Println(climbingLeaderboard(scores, alice))
}

func climbingLeaderboard(scores, alice []int32) []int32 {
	out := []int32{}
	cScores := []int32{}
	var i, score, rank int32 = int32(len(alice) - 1), 0, 1

	cScores = append(cScores, scores[0])
	for _, score = range scores {
		if score == cScores[len(cScores)-1] {
			continue
		}
		cScores = append(cScores, score)
	}
	fmt.Println(cScores)
	fmt.Println(alice)

	for _, score = range cScores {
		if len(out) == len(alice) {
			break
		}
		fmt.Println("debug ", alice[i], score)

		if alice[i] < score {
			rank++
			continue
		} else {
			out = append([]int32{rank}, out...)
			fmt.Println("out", out)
			i--
			rank++
		}
	}

	// for _, aScore = range alice {
	// 	if aScore < cScores[len(cScores)-1] {
	// 		out = append(out, int32(len(cScores)+1))
	// 		continue
	// 	}
	// 	for _, score = range cScores {
	// 		if aScore < score {
	// 			rank++
	// 		} else {
	// 			out = append(out, rank)
	// 			rank = 1
	// 			break
	// 		}
	// 	}
	// }
	return out
}
