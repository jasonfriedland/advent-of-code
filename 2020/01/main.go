package main

// two 567171
// three 212428694

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Input file
	file, _ := os.Open("./2020/01/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// Scan input
	var nums []int
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, n)
	}

	// Computed values
	var pairs = make(map[int]int)
	var triplets = make(map[int]int)

	// Populate maps
	for _, i := range nums {
		for _, j := range nums {
			// pairs
			key := i + j
			if _, exists := pairs[key]; !exists {
				pairs[key] = i * j
			}
			for _, k := range nums {
				// triplets
				key := i + j + k
				if _, exists := triplets[key]; !exists {
					triplets[key] = i * j * k
				}
			}
		}
	}

	// Results
	for k, v := range pairs {
		if k == 2020 {
			fmt.Println("two", v)
		}
	}
	for k, v := range triplets {
		if k == 2020 {
			fmt.Println("three", v)
		}
	}
}
