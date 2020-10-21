package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	n := new(big.Int)
	n.MulRange(1, 100)
	s := n.String()
	sum := 0
	for _, c := range s {
		fmt.Printf("c = %#U, int c = %d, sum = %d\n", c, int(c), sum)
		tmp, _ := strconv.Atoi(string(c))
		sum += tmp
	}
	fmt.Println(sum)
}
