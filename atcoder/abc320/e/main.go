package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/priorityqueue"
	"github.com/liyue201/gostl/utils/comparator"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var t, w, s []int
	for i := 0; i < m; i++ {
		t = append(t, nextInt())
		w = append(w, nextInt())
		s = append(s, nextInt())
	}
	ans := solve(n, m, t, w, s)
	for _, v := range ans {
		Print(v)
	}
}

func solve(n, m int, t, w, s []int) []int {
	persons := priorityqueue.New[int](comparator.IntComparator)
	for i := 0; i < n; i++ {
		persons.Push(i)
	}
	type event struct {
		time, mode int
		w, s       int
		i          int
	}
	eventComparator := func(a, b *event) int {

		if a.time == b.time {
			if a.mode == b.mode {
				return 0
			}
			if a.mode > b.mode {
				return -1
			}
			return 1
		}
		if a.time < b.time {
			return -1
		}
		return 1
	}
	events := priorityqueue.New[*event](eventComparator)
	for i := 0; i < m; i++ {
		events.Push(&event{t[i], 1, w[i], s[i], -1})
	}
	ans := make([]int, n)
	for !events.Empty() {
		e := events.Pop()
		switch e.mode {
		case 1:
			if persons.Empty() {
				continue
			}
			i := persons.Pop()
			ans[i] += e.w
			events.Push(&event{e.time + e.s, 2, 0, 0, i})
		case 2:
			persons.Push(e.i)
		}
	}

	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
