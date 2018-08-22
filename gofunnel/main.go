package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var lines map[string]struct{}

// https://www.reddit.com/r/dailyprogrammer/comments/98ufvz/20180820_challenge_366_easy_word_funnel_1/

func main() {
	// CHALLENGE
	fmt.Printf("funnel(\"leave\", \"eave\") => %v\n", funnel("leave", "eave"))
	fmt.Printf("funnel(\"reset\", \"rest\") => %v\n", funnel("reset", "rest"))
	fmt.Printf("funnel(\"dragoon\", \"dragon\") => %v\n", funnel("dragoon", "dragon"))
	fmt.Printf("funnel(\"sleet\", \"lets\") => %v\n", funnel("sleet", "lets"))
	fmt.Printf("funnel(\"skiff\", \"ski\") => %v\n", funnel("skiff", "ski"))

	// BONUS:
	// Load file
	lines = make(map[string]struct{})
	err := readLines("enable1.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Printf("bonus(\"dragoon\") => %s\n", bonus("dragoon"))
	fmt.Printf("bonus(\"boats\") => %s\n", bonus("boats"))
	fmt.Printf("bonus(\"affidavit\") => %s\n", bonus("affidavit"))

	// BONUS 2:
	fmt.Printf("bonus2() => %s\n", bonus2())
}

func funnel(source, target string) bool {
	// defer TimeTrack(time.Now(), fmt.Sprintf("funnel(%s, %s)", source, target))
	if source == target {
		return false
	}
	runes := []rune(source)
	for i := 0; i < len(source); i++ {
		str := fmt.Sprintf("%s%s", string(runes[:i]), string(runes[i+1:]))
		if str == target {
			return true
		}
	}
	return false
}

func bonus(source string) []string {
	// defer TimeTrack(time.Now(), fmt.Sprintf("bonus(%s)", source))
	var res map[string]struct{}
	res = make(map[string]struct{})
	runes := []rune(source)
	for i := 0; i < len(source); i++ {
		str := fmt.Sprintf("%s%s", string(runes[:i]), string(runes[i+1:]))
		_, ok := lines[str]
		if ok {
			res[str] = struct{}{}
		}
	}
	keys := make([]string, 0)
	for k := range res {
		keys = append(keys, k)
	}
	return keys
}

func bonus2() []string {
	defer TimeTrack(time.Now(), "bonus2")
	var res []string
	for line := range lines {
		if len(bonus(line)) >= 5 {
			res = append(res, line)
		}
	}
	return res
}

func readLines(path string) error {
	// defer TimeTrack(time.Now(), "readLines")
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines[scanner.Text()] = struct{}{}
	}
	return scanner.Err()
}

// TimeTrack functions to measure execution time.
// usage: defer util.TimeTrack(time.Now(), "function")
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
