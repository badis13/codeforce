package main

import (
	"bufio"
	"fmt"
	"os"
)

type Nail struct {
	x int
	y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c int
	nails := make(map[Nail]struct{})
	_, _ = fmt.Fscan(reader, &a)
	for i := 0; i <= a; i++ {
		_, _ = fmt.Fscan(reader, &b)
		_, _ = fmt.Fscan(reader, &c)
		nails[Nail{x: b, y: c}] = struct{}{}
	}

	var result int

	for nail1, _ := range nails {
		for nail2, _ := range nails {
			_, ok := nails[Nail{nail2.x, nail1.y}]
			if ok {
				_, ok := nails[Nail{nail1.x, nail2.y}]
				if ok {
					cur := ((nail2.y - nail1.y) * (nail2.x - nail1.x))
					if result < cur {
						result = cur
					}
				}
			}
		}

	}

	fmt.Println(abc(result))
}

func abc(res int) int {
	if res < 0 {
		return -res
	}
	return res
}
