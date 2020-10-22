package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10000; i++ {
		i_sum := divisors_sum(i)
		j_sum := divisors_sum(i_sum)
		if (i == j_sum && i != i_sum) {
			fmt.Printf("i = %d, isum = %d, jsum = %d\n", i, i_sum, j_sum)
			sum += i
		}
	}
	fmt.Printf("sum = %d", sum)
}

func divisors_sum(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n % i == 0 {
			sum += i
		}
	}
	return sum
}