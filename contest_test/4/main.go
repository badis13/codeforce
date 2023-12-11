package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// не решено, не работает

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, heightWeed, resultLine, resultColumn int
	gard := make(Heap, 0)
	_, _ = fmt.Fscan(reader, &n)
	_, _ = fmt.Fscan(reader, &m)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			_, _ = fmt.Fscan(reader, &heightWeed)
			heap.Push(&gard, Garden{
				line:     i,
				column:   j,
				priority: heightWeed,
			})

		}
	}

	item1 := heap.Pop(&gard).(*Garden)
	resultLine = item1.line
	item2 := heap.Pop(&gard).(*Garden)
	resultColumn = item2.column

	fmt.Println(resultLine, resultColumn)
}

type Heap []*Garden

type Garden struct {
	line     int
	column   int
	priority int
}

func (h Heap) Less(i, j int) bool {
	return h[i].priority > h[j].priority
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	item := x.(Garden)
	*h = append(*h, &item)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return item
}
