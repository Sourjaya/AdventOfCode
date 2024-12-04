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
	records := prepareInput(file)
	fmt.Println(checkForSafeRecords(records))
	fmt.Println(checkForSafeRecordsWithProblemDampener(records))
	// safeReportsCount :=
	// fmt.Println(safeReportsCount)
}

func checkForSafeRecords(records [][]int) int {
	safeReportCount := 0
	for _, record := range records {
		safe := 0
		length := len(record)
		for i := 0; i < length-1; i++ {
			diff := record[i] - record[i+1]
			if diff >= 1 && diff <= 3 {
				safe++
			}
			if diff <= -1 && diff >= -3 {
				safe--
			}
		}
		if safe == length-1 || safe == -(length-1) {
			safeReportCount++
		}
	}
	return safeReportCount
}

func checkForSafeRecordsWithProblemDampener(records [][]int) int {
	safeReportCount := 0
	for index, record := range records {
		if len(record) > 2 {
			withoutFirst := record[1:]
			safe := 0
			length := len(withoutFirst)
			for i := 0; i < length-1; i++ {
				diff := withoutFirst[i] - withoutFirst[i+1]
				if diff >= 1 && diff <= 3 {
					safe++
				}
				if diff <= -1 && diff >= -3 {
					safe--
				}
			}
			if safe == length-1 || safe == -(length-1) {
				safeReportCount++
				continue
			}
		}
		safe, skip, i, length, countOfSkips := 0, 0, 1, len(record), 0
		msg := ""
		for i < length-1-skip {
			skipped := false
			diff1 := record[i-1] - record[i]
			diff2 := record[i] - record[i+1]

			if diff1 >= 1 && diff1 <= 3 && diff2 >= 1 && diff2 <= 3 {
				safe++
			} else if diff1 <= -1 && diff1 >= -3 && diff2 <= -1 && diff2 >= -3 {
				safe--
			} else {
				if countOfSkips > 0 {
					break
				}
				if skip < 1 {
					diff := record[i-1] - record[i+1]
					if diff >= 1 && diff <= 3 {
						//skipped = true
						safe++
						record[i] = record[i-1]
					} else if diff <= -1 && diff >= -3 {
						//skipped = true
						safe--
						record[i] = record[i-1]
					} else {
						record = append(record[:i+1], record[i+2:]...)
						skip = skip + 1
						skipped = true
					}
					msg = msg + fmt.Sprintf("here length: %v, diff: %v, record: %v ", length, diff, record)
				}
				countOfSkips++
			}
			msg = msg + fmt.Sprintf("diff1: %v, diff2: %v, safe: %v, record: %v\n", diff1, diff2, safe, record)
			if !skipped {
				i++
			}
		}
		if skip > 0 {
			if safe == length-3 || safe == -(length-3) {
				fmt.Println(index, " : ", record)
				safeReportCount++
			}
		} else {
			if safe == length-2 || safe == -(length-2) {
				fmt.Println(index, " : ", record)
				safeReportCount++
			}
		}
		fmt.Println(msg)
	}
	return safeReportCount
}

func convertStringSliceToIntSlice(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("error converting %q to int: %w", s, err)
		}
		ints[i] = num
	}
	return ints, nil
}

func prepareInput(file *os.File) (records [][]int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trimmedInput := strings.Fields(strings.TrimSpace(scanner.Text()))
		record, err := convertStringSliceToIntSlice(trimmedInput)
		if err != nil {
			panic(err)
		}
		records = append(records, record)
	}
	return records
}
