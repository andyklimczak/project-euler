package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	f, _ := os.Open("/tmp/names.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := 0
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.ReplaceAll(input, "\"", "")
		names := strings.Split(input, ",")
		sort.Strings(names)
		for i, n := range names {
			sum := (i + 1) * name_sum(n)
			result += sum
		}
	}
	fmt.Println(result)
}

func name_sum(n string) int {
	sum := 0
	for _, c := range n {
		fmt.Printf("%s, %d, %d\n", string(c), int(c), int(c) - 64)
		sum += int(c) - 64
	}
	return sum
}
