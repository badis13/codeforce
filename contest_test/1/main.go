package main

import (
	"bufio"
	"fmt"
	"os"
)

type getContrast struct {
	index int
	val   int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b int
	min := getContrast{val: 10e9, index: -1}
	max := getContrast{val: -10e9, index: -1}
	_, _ = fmt.Fscan(reader, &a)
	for i := 1; i <= a; i++ {
		_, _ = fmt.Fscan(reader, &b)
		if min.val == b && min.index > i {
			min.index = i
			continue
		}
		if min.val > b {
			min.index = i
			min.val = b
		}

		if max.val <= b {
			max.index = i
			max.val = b
		}

	}
	fmt.Println(max.index, min.index)
}
