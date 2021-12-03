package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	horizontal, depth, aim int
}

type instruction struct {
	command string
	value   int
}

type instructions []instruction

func main() {
	input := getInput("./2021/02/input.txt")
	fmt.Println("part 1:", part1(input)) // 1882980
	fmt.Println("part 2:", part2(input)) // 1971232560
}

func part1(ins instructions) int {
	var p pos
	for _, in := range ins {
		switch in.command {
		case "forward":
			p.horizontal += in.value
		case "down":
			p.depth += in.value
		case "up":
			p.depth -= in.value
		}
	}
	return p.horizontal * p.depth
}

func part2(ins instructions) int {
	var p pos
	for _, in := range ins {
		switch in.command {
		case "forward":
			p.horizontal += in.value
			p.depth += p.aim * in.value
		case "down":
			p.aim += in.value
		case "up":
			p.aim -= in.value
		}
	}
	return p.horizontal * p.depth
}

// utils

func getInput(fileName string) instructions {
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var ins instructions
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(parts[1])
		ins = append(ins, instruction{parts[0], v})
	}
	return ins
}
