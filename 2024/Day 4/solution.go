package main

import (
	"bufio"
	"fmt"
	"os"
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

	inputString := prepareInput(file)
	fmt.Println(ceresSearchXMAS(inputString))
	fmt.Println(ceresSearchX_MAS(inputString))
}
func prepareInput(file *os.File) (inputString []string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString = append(inputString, strings.TrimSpace(scanner.Text()))
	}
	return inputString
}

func ceresSearchXMAS(inputString []string) int {
	count := 0
	di := []int{-1, 0, 1}
	dj := []int{-1, 0, 1}
	for i := 0; i < len(inputString); i++ {
		for j := 0; j < len(inputString[i]); j++ {
			if inputString[i][j] == 'X' {
				for _, k := range di {
					for _, l := range dj {
						if k == 0 && l == 0 {
							continue
						}
						if !(i+3*k < len(inputString) && i+3*k >= 0 && j+3*l >= 0 && j+3*l < len(inputString[i])) {
							continue
						}
						if inputString[i+k][j+l] == 'M' && inputString[i+2*k][j+2*l] == 'A' && inputString[i+3*k][j+3*l] == 'S' {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

func ceresSearchX_MAS(inputString []string) int {
	count := 0
	di := []int{-1, 1}
	dj := []int{-1, 1}
	for i := 1; i < len(inputString)-1; i++ {
		for j := 1; j < len(inputString[i])-1; j++ {
			if inputString[i][j] == 'A' {
				//fmt.Printf("\nIndex of A : (%d,%d)", i, j)
				str := ""
				for _, k := range di {
					for _, l := range dj {
						str = str + string(inputString[i+k][j+l])
					}
				}
				//fmt.Println("str :", str)
				if str == "MMSS" || str == "SMSM" || str == "SSMM" || str == "MSMS" {
					count++
				}
			}
		}
	}
	return count
}
