package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var r, c, x []int
	for i := 0; i < n; i++ {
		r = append(r, nextInt())
		c = append(c, nextInt())
		x = append(x, nextInt())
	}
	ans := solve(n, r, c, x)
	PrintInt(ans)
}

type line struct {
	r, c, x, i, j int
}

func solve(n int, r, c, x []int) int {
	mr, mc := make(map[int]int), make(map[int]int)
	field := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		mr[r[i]] += x[i]
		mc[c[i]] += x[i]
		if field[r[i]] == nil {
			field[r[i]] = make(map[int]int)
		}
		field[r[i]][c[i]] = x[i]
	}

	var rows, cols []line
	for k, v := range mr {
		rows = append(rows, line{k, 0, v, -1, -1})
	}
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].x > rows[j].x
	})
	for k, v := range mc {
		cols = append(cols, line{0, k, v, -1, -1})
	}
	sort.Slice(cols, func(i, j int) bool {
		return cols[i].x > cols[j].x
	})

	q := &PriorityQueue{}
	heap.Init(q)
	for i := range rows {
		heap.Push(q, line{rows[i].r, cols[0].c, rows[i].x + cols[0].x, i, 0})
	}

	var ans int
	for q.Len() > 0 {
		node := heap.Pop(q).(line)
		if node.x <= ans {
			break
		}
		ans = Max(ans, node.x-field[node.r][node.c])
		if node.j+1 < len(cols) {
			ni, nj := node.i, node.j+1
			heap.Push(q, line{node.r, cols[nj].c, rows[ni].x + cols[nj].x, ni, nj})
		}
	}

	return ans
}

type PriorityQueue []line

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].x > pq[j].x
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(line))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

func solveHonestly(n int, r, c, x []int) int {
	rows, cols := make(map[int]int), make(map[int]int)
	field := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		rows[r[i]] += x[i]
		cols[c[i]] += x[i]
		if field[r[i]] == nil {
			field[r[i]] = make(map[int]int)
		}
		field[r[i]][c[i]] = x[i]
	}
	type line struct {
		r, c, x int
	}
	var q []line

	//ここから
	for k1, v1 := range rows {
		for k2, v2 := range cols {
			q = append(q, line{k1, k2, v1 + v2})
		}
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i].x > q[j].x
	})
	//ここまでO(n**2)かかっているので
	//n+1程度に収める。n=1のとき注意

	//fmt.Println(q)
	var ans int
	for i := 0; i <= n; i++ {
		if v, found := field[q[i].r][q[i].c]; found {
			ans = Max(ans, q[i].x-v)
		} else {
			ans = Max(ans, q[i].x)
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
}

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	/*
		if x < y {
			x, y = y, x
		}
	*/
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	// x*yのオーバーフロー対策のため先にGcdで割る
	// Gcd(x, y)はxの約数のため割り切れる
	ret := x / Gcd(x, y)
	ret *= y
	return ret
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
