package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	row int
	value int
	maxValue int
}

func main() {
	fmt.Println("asdf")
//	input := `3
//7 4
//2 4 6
//8 5 9 3`
	input := ` 75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23 `

	nodes := parseInput(input)
	max := calculateMax(nodes)
	fmt.Printf("Max is: %d", max)
}

func parseInput(input string) ([]node) {
	result := []node{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	rowCount := 1
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), " ")
		for _, s := range splitString {
			val, err := strconv.Atoi(s)
			if err != nil {
				continue
			}
			temp := node{
				row: rowCount,
				value: val,
			}
			result = append(result, temp)
		}
		rowCount++
	}
	return result
}

func calculateMax(nodes []node) (int) {
	max := maxFromNode(0, nodes)
	return max
}

func maxFromNode(index int, nodes []node) (int) {
	node := nodes[index]
	leftIndex := index + node.row
	rightIndex := index + node.row + 1

	fmt.Printf("Looking at value: %d, index %d, left index: %d, right index: %d\n", node.value, index, leftIndex, rightIndex)

	if leftIndex > len(nodes) - 1 && rightIndex > len(nodes) - 1 {
		return node.value
	}

	//leftMax := make(chan int, 1)
	//defer close(leftMax)
	//rightMax := make(chan int, 1)
	//defer close(rightMax)
	//
	//go func(index int) {
	//	leftMax <- maxFromNode(index, nodes)
	//}(leftIndex)
	//go func(index int) {
	//	rightMax <- maxFromNode(index, nodes)
	//}(rightIndex)

	leftMax := maxFromNode(leftIndex, nodes)
	rightMax := maxFromNode(rightIndex, nodes)

	var max int
	if leftMax > rightMax {
		max = leftMax
	} else {
		max = rightMax
	}

	return node.value + max
}
