package main

import (
	"bufio"
	"fmt"
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

	n, m, q := nextInt(), nextInt(), nextInt()
	var x, y []int
	for i := 0; i < m; i++ {
		x = append(x, nextInt()-1)
		y = append(y, nextInt()-1)
	}
	var a, b []int
	for i := 0; i < q; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	//ans := solveHonestly(n, m, q, x, y, a, b)
	ans := solve(n, m, q, x, y, a, b)
	for _, v := range ans {
		PrintString(v)
	}
}

const bitSize = 64

//type BitSet []uint
//
//func (bs BitSet) isTrue(x int) bool {
//	return bs[x/bitSize]&(1<<(x%bitSize)) > 0
//}

func solve(n, m, q int, x, y, a, b []int) []string {
	type edge struct {
		from, to int
	}
	var es []edge
	for i := 0; i < m; i++ {
		es = append(es, edge{x[i], y[i]})
	}
	//Xi < Yiという制約があるのでXi、Yiが小さい順にソートしておく
	sort.Slice(es, func(i, j int) bool {
		if es[i].from == es[j].from {
			return es[i].to < es[j].to
		}
		return es[i].from < es[j].from
	})
	var ans []string
	for i := 0; i < q; i += bitSize {
		canReach := make([]uint, n)
		for j := 0; j < bitSize && i+j < q; j++ {
			canReach[a[i+j]] |= 1 << j
		}

		for j := 0; j < m; j++ {
			canReach[es[j].to] |= canReach[es[j].from]
		}

		for j := 0; j < bitSize && i+j < q; j++ {
			if canReach[b[i+j]]>>j&1 > 0 {
				ans = append(ans, "Yes")
			} else {
				ans = append(ans, "No")
			}
		}
	}
	return ans
}

func solveHonestly(n, m, q int, x, y, a, b []int) []string {
	type edge struct {
		from, to int
	}
	var es []edge
	for i := 0; i < m; i++ {
		es = append(es, edge{x[i], y[i]})
	}
	sort.Slice(es, func(i, j int) bool {
		if es[i].from == es[j].from {
			return es[i].to < es[j].to
		}
		return es[i].from < es[j].from
	})
	var ans []string
	for i := 0; i < q; i++ {
		canReach := make([]bool, n)
		canReach[a[i]] = true
		for j := 0; j < m; j++ {
			canReach[es[j].to] = canReach[es[j].to] || canReach[es[j].from]
		}
		if canReach[b[i]] {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
