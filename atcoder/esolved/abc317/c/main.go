package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

	n, m := nextInt(), nextInt()
	var a, b, c []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt())
	}
	//ans := solve(n, m, a, b, c)
	ans := solveByBitDP(n, m, a, b, c)
	Print(ans)
}

func solve(n, m int, a, b, c []int) int {
	e := make([][]int, n)
	for i := range e {
		e[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		e[a[i]][b[i]] = c[i]
		e[b[i]][a[i]] = c[i]
	}
	var idx []int
	for i := 0; i < n; i++ {
		idx = append(idx, i)
	}
	var ans int
	for {
		var s int
		for i := 0; i < n-1; i++ {
			from, to := idx[i], idx[i+1]
			if e[from][to] == 0 {
				break
			}
			s += e[from][to]
		}
		ans = Max(ans, s)
		if !NextPermutation(sort.IntSlice(idx)) {
			break
		}
		//fmt.Println(idx)
	}
	return ans
}

func solveByBitDP(n, m int, a, b, c []int) int {
	e := make([][]int, n)
	for i := range e {
		e[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		e[a[i]][b[i]] = c[i]
		e[b[i]][a[i]] = c[i]
	}
	//patの街を訪問してiに入るまでに通った最大の長さ
	dp := make([][]int, 1<<n)
	for pat := 0; pat < (1 << n); pat++ {
		dp[pat] = make([]int, n)
		for i := range dp[pat] {
			dp[pat][i] = -1
		}
	}
	for i := 0; i < n; i++ {
		dp[1<<i][i] = 0
	}
	var ps []int
	for pat := 1; pat < (1 << n); pat++ {
		ps = append(ps, pat)
	}
	sort.Slice(ps, func(i, j int) bool {
		return bits.OnesCount(uint(ps[i])) < bits.OnesCount(uint(ps[j]))
	})
	for _, pat := range ps {
		//今いる場所
		for i := 0; i < n; i++ {
			//到達不能な状態
			if dp[pat][i] < 0 {
				continue
			}
			//次の場所
			for j := 0; j < n; j++ {
				//すでに訪問済み
				if (pat>>j)&1 == 1 {
					continue
				}
				next := pat | (1 << j)
				//iからjへ結ぶ道路がなし
				if e[i][j] == 0 {
					continue
				}
				dp[next][j] = Max(dp[next][j], dp[pat][i]+e[i][j])
			}
		}
	}
	var ans int
	for _, pat := range ps {
		for i := 0; i < n; i++ {
			ans = Max(ans, dp[pat][i])
		}
	}
	return ans
}

type Edge struct {
	t, w int
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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
