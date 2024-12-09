package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	modifiedRules, updates := prepareInput(file)

	fmt.Println(countMiddlePageNumber(modifiedRules, updates))
}

func countMiddlePageNumber(modifiedRules map[int]map[int]int, updates [][]int) (int, int) {
	count, count2 := 0, 0
	for _, update := range updates {
		if checkValidUpdate(update, modifiedRules) {
			count = count + update[len(update)/2]
		} else {
			count2 = count2 + middleAfterCorrection(update, modifiedRules)
		}
	}
	return count, count2
}

func middleAfterCorrection(update []int, modifiedRules map[int]map[int]int) int {
	i := 1
	fmt.Println(update)
	for i < len(update) {
		fmt.Println("i = ", i)
		j := 0
		for j < i {
			rules := modifiedRules[update[i]]
			if _, ok := rules[update[j]]; ok {
				update[i], update[j] = update[j], update[i]
				i = j + 1
				break
			}
			j++
		}
		i++
	}
	return update[len(update)/2]
}

func checkValidUpdate(update []int, modifiedRules map[int]map[int]int) bool {
	for i := 1; i < len(update); i++ {
		for j := 0; j < i; j++ {
			rules := modifiedRules[update[i]]
			if _, ok := rules[update[j]]; ok {
				return false
			}
		}
	}
	return true
}
func prepareInput(file *os.File) (map[int]map[int]int, [][]int) {
	scanner := bufio.NewScanner(file)

	rules := make(map[int][]int)

	modifiedRules := make(map[int]map[int]int)
	updates := make([][]int, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		pairs := strings.Split(scanner.Text(), "|")
		key, _ := strconv.Atoi(pairs[0])
		value, _ := strconv.Atoi(pairs[1])
		rules[key] = append(rules[key], value)
	}

	for key, value := range rules {
		modifiedRules[key] = make(map[int]int)
		for _, val := range value {
			modifiedRules[key][val] = key
		}
	}

	for scanner.Scan() {
		update := strings.Split(scanner.Text(), ",")
		var nums []int
		for _, value := range update {
			num, _ := strconv.Atoi(value)
			nums = append(nums, num)
		}
		updates = append(updates, nums)
	}
	return modifiedRules, updates
}
