package main

import "fmt"
import "math"

func main() {
	fmt.Println("hello world")
	count := 1
	num := 1
	n := 0
	for n < 500 {
		n = num_divisors(num)
		fmt.Println(num, n)
		count++
		num += count
	}
	fmt.Println("done")
}

func num_divisors(n int) (int) {
	count := 2
	i := 2.0

	for int(math.Pow(i, 2)) < n {
		if n % int(i) == 0 {
			count += 2
		}
		i += 1
	}
	if int(math.Pow(i, 2)) == n {
		count += 1
	}
	return count
}
