package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/queue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type person struct {
	x, y int
	err  error
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var a, b, x, y []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := solve(n, m, a, b, x, y)
	for _, v := range ans {
		if v.err != nil {
			Print("undecidable")
		} else {
			Print2(v.x, v.y)
		}
	}
}

func solve(n, m int, a, b, x, y []int) []person {
	type edge struct {
		from, to int
		x, y     int
	}
	e := make([][]edge, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], edge{a[i], b[i], x[i], y[i]})
		e[b[i]] = append(e[b[i]], edge{b[i], a[i], -x[i], -y[i]})
	}
	ans := make([]person, n)
	for i := range ans {
		ans[i].err = errors.New("undecidable")
	}
	q := queue.New[edge]()
	q.Push(edge{-1, 0, 0, 0})
	ans[0].err = nil
	for !q.Empty() {
		cur := q.Pop()
		for _, next := range e[cur.to] {
			//すでに訪問済み
			if ans[next.to].err == nil {
				continue
			}
			ans[next.to].x = ans[cur.to].x + next.x
			ans[next.to].y = ans[cur.to].y + next.y
			ans[next.to].err = nil
			q.Push(edge{cur.to, next.to, next.x, next.y})
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

func Print2(x, y any) {
	defer out.Flush()
	fmt.Fprintln(out, x, y)
}
