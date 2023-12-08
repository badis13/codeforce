package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k, num int
	nums := make([]int, 0, 32)
	_, _ = fmt.Fscan(reader, &n)
	_, _ = fmt.Fscan(reader, &k)
	if n == 0 {
		fmt.Println(n)
		return
	}

	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &num)
		nums = append(nums, num)
	}

	if n == 1 {
		if nums[0] < k {
			fmt.Println(n)
			return
		}
	}

	fmt.Println(getSubslice(k, nums))
}

func getSubslice(k int, nums []int) int {
	sumZero := 0
	sum := 0
	result := 0
	l := 0
	r := l + 1

	for r <= len(nums) {
		cur := nums[l:r]
		if cur[r-l-1] == 0 {
			sumZero++
			if sumZero > 1 {
				l++
				r = l + 1
				sum = 0
				sumZero = 0
				continue
			}

		}
		sum += cur[r-l-1]
		if sum <= k {
			result++
			r++
			if r > len(nums) && len(cur) > 1 {
				r--
				l++
				r = l + 1
				sum = 0
				sumZero = 0
				continue
			}

			if sumZero < 2 {
				continue
			}

		}
		l++
		r = l + 1
		sum = 0
		sumZero = 0

	}

	return result
}
