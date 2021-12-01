package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := getInput("./2021/01/input.txt")
	fmt.Println("part 1:", getIncreases(input, 1)) // 1676
	fmt.Println("part 2:", getIncreases(input, 3)) // 1706
}

func getIncreases(nums []int, sumOffset int) int {
	var sums []int
	for i := range nums {
		sums = append(sums, sum(nums[i:i+sumOffset]))
	}
	var result int
	for i := 1; i < len(sums); i++ {
		if sums[i] > sums[i-1] {
			result++
		}
	}
	return result
}

func sum(nums []int) int {
	var result int
	for _, n := range nums {
		result += n
	}
	return result
}

// utils

func getInput(fileName string) []int {
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var nums []int
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, n)
	}
	return nums
}
