package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"

	"github.com/golang-collections/collections/set"
)

func Solution2() {
	set1 := set.New()
	arr2 := map[int64]int64{}

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

		set1.Insert(number1)
		arr2[number2]++
	}

	var simScore int64 = 0
	set1.Do(func(item interface{}) {
		if _, ok := arr2[item.(int64)]; ok {
			simScore += (item.(int64) * arr2[item.(int64)])
		}
	})

	fmt.Println(simScore)
}
