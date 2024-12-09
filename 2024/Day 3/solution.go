package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	numbers1, numbers2 := prepareInput(file)

	fmt.Println(sumOfProducts(numbers1))
	fmt.Println(sumOfProducts(numbers2))
}

func findMatchAndConvert(str string) (numbers []int) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(str, -1)

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		numbers = append(numbers, num1, num2)
	}
	return numbers
}

func findMatchAndConvertConditional(str string) (numbers []int) {
	re := regexp.MustCompile(`((mul\((\d{1,3}),(\d{1,3})\)))|(do\(\))|(don\'t\(\))`)

	allMatches := re.FindAllStringSubmatch(str, -1)
	mulEnabled := true
	for _, match := range allMatches {
		if match[0] != "" {
			if match[0] == "do()" {
				mulEnabled = true
			} else if match[0] == "don't()" {
				mulEnabled = false
			}
		}

		// Check for multiplication commands (mul(x, y))
		if match[3] != "" && match[4] != "" {
			if mulEnabled {
				// Convert match[2] and match[3] to integers
				num1, _ := strconv.Atoi(match[3])
				num2, _ := strconv.Atoi(match[4])
				numbers = append(numbers, num1, num2)
			}
		}
	}
	return numbers
}

func prepareInput(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)
	var trimmedInput string
	for scanner.Scan() {
		trimmedInput = trimmedInput + strings.TrimSpace(scanner.Text())
	}
	return findMatchAndConvert(trimmedInput), findMatchAndConvertConditional(trimmedInput)
}

func sumOfProducts(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i = i + 2 {
		sum = sum + mul(numbers[i], numbers[i+1])
	}
	return sum
}

func mul(num1, num2 int) int {
	return num1 * num2
}
