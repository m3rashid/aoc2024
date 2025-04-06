package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"unicode"
)

func Solution1() {
	arr1 := []int64{}
	arr2 := []int64{}

	file, err := os.Open("day1/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		var number1 int64 = 0
		var number2 int64 = 0

		useNum1 := true
		for _, ch := range str {
			if ch == ' ' {
				useNum1 = false
				continue
			}
			if unicode.IsDigit(ch) {
				if useNum1 {
					number1 = number1*10 + int64(ch-'0')
				} else {
					number2 = number2*10 + int64(ch-'0')
				}
			}
		}

		arr1 = append(arr1, number1)
		arr2 = append(arr2, number2)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	var diff int64 = 0
	for i := 0; i < len(arr1); i++ {
		if arr1[i] > arr2[i] {
			diff += arr1[i] - arr2[i]
		} else {
			diff += arr2[i] - arr1[i]
		}
	}

	fmt.Printf("%#v\n", diff) // 1320851
}
