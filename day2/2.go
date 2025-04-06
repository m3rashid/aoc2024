package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkSafe(numbers []int64) bool {
	isDecreasing := true
	if numbers[1] > numbers[0] {
		isDecreasing = false
	}

	isSafe := true

	indexToExclude := -1

	for i := 0; i < len(numbers); i++ {

	}

	return isSafe
}

func Solution2() {
	file, err := os.Open("day2/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int64{}

	safeCounts := 0

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		str := scanner.Text()

		var number = 0
		for _, ch := range str {
			if ch == ' ' {
				numbers = append(numbers, int64(number))
				number = 0
				continue
			}

			if ch >= '0' && ch <= '9' {
				number = number*10 + int(ch-'0')
			}
		}

		if number != 0 {
			numbers = append(numbers, int64(number))
		}

		// maxUnsafe := 0
		// for i := 1; i < len(numbers); i++ {
		// 	if (numbers[i] == numbers[i-1]) ||
		// 		(isDecreasing && numbers[i] > numbers[i-1]) ||
		// 		(!isDecreasing && numbers[i] < numbers[i-1]) ||
		// 		math.Abs(float64(numbers[i]-numbers[i-1])) > 3 {
		// 		maxUnsafe++
		// 	}
		// }

		// if maxUnsafe == 0 || maxUnsafe == 1 {
		// 	safeCounts++
		// 	// fmt.Println(lineNumber, "safe")
		// } else {
		// 	// fmt.Println(lineNumber, "unsafe")
		// }

		if checkSafe(numbers) {
			safeCounts++
		}

		numbers = []int64{}
	}

	fmt.Println(safeCounts)
}
