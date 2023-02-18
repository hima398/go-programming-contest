package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	q := nextInt()
	var s, t []int
	for i := 0; i < q; i++ {
		s = append(s, nextInt()-1)
		t = append(t, nextInt()-1)
	}
	//ans := solveHonestly(n, m, a, b, q, s, t)
	ans := solve(n, m, a, b, q, s, t)
	PrintVertically(ans)
}

type BitSet []uint

func (bs BitSet) isTrue(x int) bool {
	return bs[x/64]&(1<<(x%64)) > 0
}

func solve(n, m int, a, b []int, q int, s, t []int) []int {
	const bitSize = 64
	canReach := make([]BitSet, n)
	for i := range canReach {
		canReach[i] = make(BitSet, n/bitSize+1)
	}
	for i := 0; i < m; i++ {
		canReach[a[i]][b[i]/bitSize] |= 1 << (b[i] % bitSize)
	}
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		ans[i] = -1
	}

	for k := range canReach {
		for i := range canReach {
			if canReach[i].isTrue(k) {
				for x, pat := range canReach[k] {
					canReach[i][x] |= pat
				}
			}
		}

		for i := 0; i < q; i++ {
			if ans[i] == -1 && canReach[s[i]].isTrue(t[i]) {
				ans[i] = Max(k, Max(s[i], t[i])) + 1
			}
		}
	}

	return ans

}

func solveHonestly(n, m int, a, b []int, q int, s, t []int) []int {
	canReach := make([][]bool, n)
	cost := make([][]int, n)
	for i := range cost {
		canReach[i] = make([]bool, n)
		cost[i] = make([]int, n)
		for j := range cost[i] {
			cost[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		canReach[i][i] = true
		cost[i][i] = i + 1
	}
	for i := 0; i < m; i++ {
		canReach[a[i]][b[i]] = true
		cost[a[i]][b[i]] = Max(a[i]+1, b[i]+1)
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j || j == k || k == i {
					continue
				}
				if !canReach[i][j] && canReach[i][k] && canReach[k][j] {
					canReach[i][j] = true
					cost[i][j] = Max(i+1, Max(j+1, k+1))
				}
			}
		}
	}
	var ans []int
	for i := 0; i < q; i++ {
		ans = append(ans, cost[s[i]][t[i]])
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
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
