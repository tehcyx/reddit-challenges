package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

// https://www.reddit.com/r/dailyprogrammer/comments/7ttiq5/20180129_challenge_349_easy_change_calculator/

// run like this:
// $ go run main.go 10 5 5 2 2 1
//
// Output:
// Solution found: n = 2
// Numbers used: [5 5]

func main() {
	need, _ := strconv.Atoi(os.Args[1])
	have := make([]int, len(os.Args[:0]))
	for i := 2; i < len(os.Args); i++ {
		value, _ := strconv.Atoi(os.Args[i])
		have = append(have, value)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(have)))
	var tmp, pointer, currentSum int
	var change []int
	pointer = 0
	currentSum = 0
	for tmp != need && pointer != len(have) {
		if currentSum+have[pointer] <= need {
			currentSum = currentSum + have[pointer]
			change = append(change, have[pointer])
			if currentSum == need {
				break
			} else if currentSum > need {
				change = change[:len(change)-1]
			}
		}
		pointer++
	}
	if currentSum < need {
		fmt.Printf("Not possible: n >= %d\n", len(change))
	} else {
		fmt.Printf("Solution found: n = %d\nNumbers used: %d\n", len(change), change)
	}
}
