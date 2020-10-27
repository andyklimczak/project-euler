package main

import "fmt"

func main() {
	n := 28123
	abundant_nums := calculate_adundant_nums(n)
	fmt.Println(abundant_nums)

	res := sum_non_two_abundant_numbers(abundant_nums, n)
	fmt.Println(res)
}

func calculate_adundant_nums(n int) (tmp map[int]bool) {
	tmp = make(map[int]bool)
	for i := 0; i < n; i++ {
		is_abundant_num := is_abundant_num(i)
		if (is_abundant_num) {
			tmp[i] = true
		}
	}
	return tmp;
}

func is_abundant_num(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n % i == 0 {
			sum += i
		}
	}
	return sum > n
}

func sum_non_two_abundant_numbers(abundant_nums map[int]bool, n int) (result int) {
	result = 0

Outer:
	for i := 1; i <= n; i++ {
		for key, _ := range abundant_nums {
			tmp := i - key
			if abundant_nums[tmp] {
				continue Outer
			}
		}
		result += i
	}

	return result
}