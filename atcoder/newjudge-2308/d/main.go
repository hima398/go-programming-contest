package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/queue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n1, n2, m := nextInt(), nextInt(), nextInt()
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	ans := solve(n1, n2, m, a, b)
	Print(ans)
}

func solve(n1, n2, m int, a, b []int) int {
	e := make([][]int, n1+n2)

	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	const INF = math.MaxInt
	d := make([]int, n1+n2)
	for i := range d {
		d[i] = INF
	}
	bfs := func(cur int) int {
		q := queue.New[int]()

		q.Push(cur)
		d[cur] = 0
		//mx := 0
		//idx := cur
		var res int
		for !q.Empty() {
			cur := q.Pop()
			for _, next := range e[cur] {
				if d[next] != INF {
					continue
				}
				q.Push(next)
				d[next] = d[cur] + 1
				res = Max(res, d[next])
			}
		}
		return res
	}
	d1 := bfs(0)
	d2 := bfs(n1 + n2 - 1)
	ans := d1 + d2 + 1
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
