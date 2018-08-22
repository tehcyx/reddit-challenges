package main

import (
	"encoding/json"
	"fmt"
	"math"
)

var steps map[string]bool

// https://www.reddit.com/r/dailyprogrammer/comments/8sjcl0/20180620_challenge_364_intermediate_the_ducci/

func main() {
	steps = make(map[string]bool)

	//fmt.Printf("%d Steps\n", ducci([]int{0, 653, 1854, 4063}))

	// Challenge input:
	// (1, 5, 7, 9, 9)
	fmt.Printf("%d Steps\n", ducci([]int{1, 5, 7, 9, 9}))
	// (1, 2, 1, 2, 1, 0)
	fmt.Printf("%d Steps\n", ducci([]int{1, 2, 1, 2, 1, 0}))
	// (10, 12, 41, 62, 31, 50)
	fmt.Printf("%d Steps\n", ducci([]int{10, 12, 41, 62, 31, 50}))
	// (10, 12, 41, 62, 31)
	fmt.Printf("%d Steps\n", ducci([]int{10, 12, 41, 62, 31}))
}

func ducci(input []int) int {
	s, _ := json.Marshal(input)
	fmt.Println(string(s))
	if steps[string(s)] {
		return 1
	}
	steps[string(s)] = true
	if input[0] == 0 && allEqual(input) {
		return 1
	}
	newDucci := []int{}
	for i := 0; i < len(input); i++ {
		cmp := i + 1
		if cmp == len(input) {
			cmp = 0
		}
		diff := int(math.Abs(float64(input[i] - input[cmp])))
		newDucci = append(newDucci, diff)
	}
	return ducci(newDucci) + 1
}

func allEqual(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}
