package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	left  []int
	right []int
)

func main() {

	if len(os.Args) < 2 {
		panic("No input file given")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	left, right = prepareInput(file)
	sort.Ints(left)
	sort.Ints(right)
	total_distance := findTotalDistance(left, right)
	fmt.Println(total_distance)
	fmt.Println(findSimilarityScore(left, right))
}

func prepareInput(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		trimmedInput := strings.Fields(strings.TrimSpace(scanner.Text()))
		l, lerr := strconv.Atoi(trimmedInput[0])
		r, rerr := strconv.Atoi(trimmedInput[1])
		if lerr != nil || rerr != nil {
			panic(fmt.Errorf("%v,%v", lerr, rerr))
		}
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func findTotalDistance(left []int, right []int) int {
	total_distance := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff *= -1
		}
		total_distance += diff
	}
	return total_distance
}

func findSimilarityScore(left []int, right []int) int {
	similarity_score := 0
	m := make(map[int]int)
	for _, vals := range right {
		m[vals]++
	}
	for _, vals := range left {
		if m[vals] > 0 {
			similarity_score += vals * m[vals]
		}
	}
	return similarity_score
}
