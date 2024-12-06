package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var leftList []int
	var rightList []int

	file, err := os.Open("./day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		leftNumber, _ := strconv.Atoi(numbers[0])
		rightNumber, _ := strconv.Atoi(numbers[1])
		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	firstResult := firstStar(leftList, rightList)
	secondResult := secondStart(leftList, rightList)

	fmt.Printf("First star result : %d\n", firstResult)
	fmt.Printf("Second star result : %d\n", secondResult)
}

func firstStar(leftList, rightList []int) (count int) {
	for i := range leftList {
		res := abs(leftList[i], rightList[i])

		count += res
	}

	return count
}

func secondStart(leftList, rightList []int) (count int) {
	for i := range leftList {
		occu := 0
		for j := range rightList {
			if rightList[j] == leftList[i] {
				occu++
			}
		}
		count += leftList[i] * occu
	}

	// another solution
	// countMap := make(map[int]int)
	// for _, i := range rightList {
	// 	countMap[i]++
	// }

	// for _, j := range leftList {
	// 	count += j * countMap[j]
	// }

	return count
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
