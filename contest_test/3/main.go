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
	var l, r, sum, zeroCount, result int

	for r < len(nums) {

		if nums[r] == 0 {
			zeroCount++
		}

		sum += nums[r]

		for sum > k || zeroCount > 1 {
			if nums[l] == 0 {
				zeroCount--
			}
			sum -= nums[l]
			l++
		}

		result += r - l + 1
		r++
	}

	return result
}
