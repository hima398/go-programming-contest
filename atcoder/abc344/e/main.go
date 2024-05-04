package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/list/bidlist"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	q := nextInt()
	var t, x []int
	y := make([]int, q)
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
		if t[i] == 1 {
			y[i] = nextInt()
		}
	}

	ans := solve(n, a, q, t, x, y)

	PrintHorizonaly(ans)
}

func solve(n int, a []int, q int, t, x, y []int) []int {
	b := bidlist.New[int]()
	m := make(map[int]*bidlist.Node[int])

	for _, ai := range a {
		b.PushBack(ai)
		m[ai] = b.BackNode()
	}

	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			target := m[x[i]]
			b.InsertAfter(y[i], target)
			m[y[i]] = target.Next()
		case 2:
			target := m[x[i]]
			b.Remove(target)
			delete(m, x[i])
		}
	}
	var ans []int
	for cur := b.FrontNode(); cur != nil; cur = cur.Next() {
		ans = append(ans, cur.Value)
	}
	return ans
}

func firstsolve(n int, a []int, q int, t, x, y []int) []int {
	const sentinel = int(1e9) + 1

	type node struct {
		v          int
		prev, next *node
	}
	m := make(map[int]*node)
	m[0] = &node{0, nil, nil}

	cur := m[0]
	for _, ai := range a {
		nd := new(node)
		nd.v = ai
		nd.prev = cur

		cur.next = nd
		m[ai] = nd
		cur = nd
	}
	nd := new(node)
	nd.v = sentinel
	nd.prev = cur
	cur.next = nd
	m[sentinel] = nd

	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			cur := m[x[i]]

			nd := new(node)
			nd.v = y[i]
			nd.prev = cur
			nd.next = cur.next

			cur.next = nd
			m[y[i]] = nd
		case 2:
			cur := m[x[i]]
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
			delete(m, x[i])
		}

	}

	var ans []int
	for cur := m[0].next; cur.v != sentinel; cur = cur.next {
		ans = append(ans, cur.v)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
