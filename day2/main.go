package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func ReadInput(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var content []string
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content, nil
}

func getReports(s []string) [][]int {
	result := make([][]int, len(s))
	for i, report := range s {
		r := strings.Split(report, " ")
		result[i] = make([]int, len(r))
		for j, level := range r {
			result[i][j], _ = strconv.Atoi(level)
		}
	}
	return result
}

func isSafeReport(report []int) bool {
	isInc := false
	if report[1]-report[0] > 0 {
		isInc = true
	}
	for idx := 1; idx < len(report); idx++ {
		if isInc && (report[idx]-report[idx-1]) > 0 && (report[idx]-report[idx-1]) < 4 {
			continue
		}
		if !isInc && (report[idx]-report[idx-1]) < 0 && (report[idx]-report[idx-1]) > -4 {
			continue
		}
		return false
	}
	return true
}

func isSafeReportTolerance(report []int) bool {
	inc := 0
	for idx := 1; idx < len(report); idx++ {
		if report[idx]-report[idx-1] > 0 {
			inc++
		} else {
			inc--
		}
	}
	if inc < len(report)-3 && inc > -(len(report)-3) {
		return false // there are more than one bad number
	}
	isInc := false
	if inc > 0 {
		isInc = true
	}
	for idx := 1; idx < len(report); idx++ {
		if isInc && (report[idx]-report[idx-1]) > 0 && (report[idx]-report[idx-1]) < 4 {
			continue
		}
		if !isInc && (report[idx]-report[idx-1]) < 0 && (report[idx]-report[idx-1]) > -4 {
			continue
		}
		// If there is a wrong level check current one and adjacent
		newReport := slices.Clone(report)
		newReport = append(newReport[:idx], newReport[idx+1:]...)
		if isSafeReport(newReport) {
			return true
		} else {
			newReport = slices.Clone(report)
			newReport = append(newReport[:idx-1], newReport[idx:]...)
			if isSafeReport(newReport) {
				return true
			}
		}
		return false
	}
	return true
}

func getNumSafeReports(s []string, tolerateBadLevel bool) int {
	result := 0
	reports := getReports(s)
	for _, r := range reports {
		if !tolerateBadLevel && isSafeReport(r) {
			result += 1
		} else if tolerateBadLevel && isSafeReportTolerance(r) {
			result += 1
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("./input1.txt")
	output, _ := ReadInput(absPathName)

	fmt.Println(getNumSafeReports(output, false))
	fmt.Println(getNumSafeReports(output, true))
}
